/*

	Copyright 2021, Pure Storage Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

*/

package auth

import (
	"bufio"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"strings"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golang.org/x/crypto/ssh"
)

const rootPgroupName = "pgroup-auto"

func sshSetup(ctx context.Context, host string, pureuserPrivateKey []byte) (*ssh.Client, error) {
	authMethod, err := pureuserPublicKeyAuth(pureuserPrivateKey)
	if err != nil {
		return nil, err
	}

	sshConfig := &ssh.ClientConfig{
		User:            "pureuser",
		Auth:            []ssh.AuthMethod{authMethod},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	tflog.Trace(ctx, "SSHSetup dial")
	conn, err := net.Dial("tcp", net.JoinHostPort(host, "22"))
	if err != nil {
		return nil, fmt.Errorf("dial error: %+v", err)
	}

	sshConn, chans, reqs, err := newClientConnWithRetries(ctx, conn, host, sshConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create new ssh client session: %+v", err)
	}

	client := ssh.NewClient(sshConn, chans, reqs)
	tflog.Trace(ctx, "SSHSetup connected")
	return client, nil
}

// newClientConnWithRetries wraps ssh.NewClientConn and retries
// to establish the connection until it's successful with a
// 60 second delay between each attempt.
//
// When starting new Purity instance it's not guaranteed that
// PUREUSER_PUBLIC_KEY will be set before Azure is signaled
// about creation completion.
//
// This adds a simple workaround which could be removed once
// this issue is resolved on the Purity side and backported
// to all versions.
func newClientConnWithRetries(
	ctx context.Context,
	c net.Conn,
	addr string,
	config *ssh.ClientConfig,
) (sshConn ssh.Conn, chans <-chan ssh.NewChannel, reqs <-chan *ssh.Request, err error) {
	err = retry.Do(
		func() error {
			sshConn, chans, reqs, err = ssh.NewClientConn(c, addr, config)

			return err
		},
		retry.Attempts(math.MaxUint16),
		retry.Delay(time.Duration(60)*time.Second),
		retry.DelayType(retry.FixedDelay),
		retry.Context(ctx),
		retry.OnRetry(func(n uint, err error) {
			tflog.Trace(ctx, fmt.Sprintf("NewClientConn retry %d: %s", n, err))
		}),
		retry.RetryIf(func(err error) bool {
			return strings.Contains(err.Error(), "ssh: unable to authenticate")
		}),
	)

	return
}

func generateSecretPayloadReal(ctx context.Context, host string, pureuserPrivateKey []byte) ([]byte, error) {
	tflog.Trace(ctx, "GenerateSecretPayload")
	privatePem, publicPem, err := generateKeyPair()
	if err != nil {
		return nil, err
	}

	client, err := sshSetup(ctx, host, pureuserPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("SSH Client setup failed: %w", err)
	}

	adminUserPwd, err := generateAdminUserPwd()
	if err != nil {
		return nil, fmt.Errorf("failed to generate a random password: %+v", err)
	}

	const adminUser = "orchestration-user"
	const apiClient = "orchestration-client"

	err = executeSSHPureAdminCreate(ctx, client, adminUser, adminUserPwd)
	if err != nil {
		return nil, err
	}

	issuer, clientID, keyID, err := executeSSHPureAPIClientSetup(ctx, client, apiClient, string(publicPem))
	if err != nil {
		return nil, err
	}

	payload := SecretPayload{
		UserName:       adminUser,
		Issuer:         issuer,
		ClientID:       clientID,
		KeyID:          keyID,
		RestPrivateKey: string(privatePem),
	}
	credentials, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}

func optOutDefaultProtectionPolicyReal(
	ctx context.Context,
	host string, pureuserPrivateKey []byte,
) error {
	client, err := sshSetup(ctx, host, pureuserPrivateKey)
	if err != nil {
		return fmt.Errorf("SSH Client setup failed: %w", err)
	}

	if out, err := executeSSHPureArrayRemovePgroupsFromDefaultProtections(ctx, client); err != nil {
		if strings.Contains(out, "invalid choice: 'default-protection'") {
			log.Println("OptOutDefaultProtectionPolicy: RemoveDefaultProtections: missing CLI subcommand")
			return nil
		}

		return fmt.Errorf("OptOutDefaultProtectionPolicy: RemoveDefaultProtections: out=%s err=%w", out, err)
	}

	if out, err := executeSSHPurePgroupDestroy(ctx, client, rootPgroupName); err != nil {
		if strings.Contains(out, "Protection group does not exist") {
			log.Printf("OptOutDefaultProtectionPolicy: PgroupDestroy: pgroup %q does not exist\n", rootPgroupName)
			return nil
		}

		return fmt.Errorf("OptOutDefaultProtectionPolicy: PgroupDestroy: out=%s err=%w", out, err)
	}

	if out, err := executeSSHPurePgroupEradicate(ctx, client, rootPgroupName); err != nil {
		return fmt.Errorf("OptOutDefaultProtectionPolicy: PgroupEradicate: out=%s err=%w", out, err)
	}

	return nil
}

func generateKeyPair() ([]byte, []byte, error) {
	// generate private keys
	privateRSAKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	if err = privateRSAKey.Validate(); err != nil {
		return nil, nil, err
	}

	privateBytes := x509.MarshalPKCS1PrivateKey(privateRSAKey)
	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateBytes,
	}

	privatePem := pem.EncodeToMemory(privateBlock)

	// generate public keys
	publicBytes, err := x509.MarshalPKIXPublicKey(&privateRSAKey.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicBytes,
	}

	publicPem := pem.EncodeToMemory(publicKeyBlock)

	return privatePem, publicPem, nil
}

// generates a random password that we never intend to use
func generateAdminUserPwd() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// This is a callable that can be implemented to control the SSH interactive session.
// waitFor is a callable that will block until the remote sends the given string
type sshIOProcessor func(waitFor func(s string), in io.Writer)

func processPublicKeyPrompt(ctx context.Context, publicKey string) sshIOProcessor {
	return func(waitFor func(s string), in io.Writer) {
		waitFor("Please enter public key followed by Enter and then Ctrl-D:")
		in.Write([]byte(publicKey + "\n\x04"))
		waitFor("Name")
	}
}

func processPasswordPrompt(ctx context.Context, password string) sshIOProcessor {
	return func(waitFor func(s string), in io.Writer) {
		waitFor("Enter password:")
		tflog.Trace(ctx, "sending password first")
		io.WriteString(in, password+"\n")
		waitFor("Retype password:")
		tflog.Trace(ctx, "sending password second")
		io.WriteString(in, password+"\n")

		// Wait for the header of the results Table...
		// That should be a clear indication that we have succeeded...
		waitFor("Name")
	}
}

func executeSSHPureAdminCreate(ctx context.Context, client *ssh.Client, username, password string) error {
	tflog.Trace(ctx, "ExecuteSSHPureAdminCreate running pureadmin create")
	return executeSSHCommandWithInputProcessing(ctx, client, "pureadmin create --role array_admin "+username, processPasswordPrompt(ctx, password))
}

func executeSSHPureAPIClientSetup(ctx context.Context, client *ssh.Client, apiClientName, publicKeyPem string) (issuer, clientID, keyID string, err error) {
	tflog.Trace(ctx, "ExecuteSSHPureAPIClientSetup running pureapiclient create")
	err = executeSSHCommandWithInputProcessing(ctx, client, "pureapiclient create --max-role array_admin --public-key "+apiClientName, processPublicKeyPrompt(ctx, publicKeyPem))
	if err != nil {
		return
	}

	tflog.Trace(ctx, "ExecuteSSHPureAPIClientSetup running pureapiclient enable")
	if _, err = executeSSHCommandAndReturnOutput(client, "pureapiclient enable "+apiClientName); err != nil {
		return
	}

	tflog.Trace(ctx, "ExecuteSSHPureAPIClientSetup running pureapiclient list")
	output, err := executeSSHCommandAndReturnOutput(client, "pureapiclient list "+apiClientName+" --csv")
	if err != nil {
		return
	}
	if len(output) != 2 {
		err = fmt.Errorf("got an unexpected number of outputs from command: pureapiclient list " + apiClientName)
		return
	}
	tflog.Trace(ctx, "ExecuteSSHPureAPIClientSetup done executing commands")
	fields := strings.Split(output[0], ",")
	values := strings.Split(output[1], ",")

	apiClientMap := make(map[string]string)
	for i := 0; i < len(fields); i++ {
		apiClientMap[fields[i]] = values[i]
	}

	issuer = apiClientMap["Issuer"]
	clientID = apiClientMap["Client ID"]
	keyID = apiClientMap["Key ID"]

	return

}

func executeSSHPureArrayRemovePgroupsFromDefaultProtections(ctx context.Context, client *ssh.Client) (string, error) {
	return executeSSHCommandAndReturnCombinedOutput(client, "purearray default-protection set \"\" --pgroup \"\"")
}

func executeSSHPurePgroupDestroy(ctx context.Context, client *ssh.Client, pgroupName string) (string, error) {
	return executeSSHCommandAndReturnCombinedOutput(client, "purepgroup destroy "+pgroupName)
}

func executeSSHPurePgroupEradicate(ctx context.Context, client *ssh.Client, pgroupName string) (string, error) {
	return executeSSHCommandAndReturnCombinedOutput(client, "purepgroup eradicate "+pgroupName)
}

// Use this to ensure that something is closed (which usually does the actually
// work for aborting) when a context is cancelled
func closeOnCancel(ctx context.Context, closer io.Closer) {
	done := ctx.Done()
	if done != nil {
		go func() {
			<-done
			tflog.Trace(ctx, "processing cancel", map[string]interface{}{})
			closer.Close()
		}()
	}
}

// This is a helper, it consumes outBufferred one rune at a time, waiting until we see a match or hit an error
func waitForLineContainingText(ctx context.Context, expectedString string, errHad *error, processOutputTail *string, outBuffered *bufio.Reader) {
	ctx = tflog.SetField(ctx, "expectedString", expectedString)
	tflog.Trace(ctx, "waiting for line containing text")
	for {
		newRune, _, err := outBuffered.ReadRune()
		if err != nil {
			tflog.Error(ctx, "read failed before we found string", map[string]interface{}{
				"err": err,
			})
			*errHad = err
			break
		}
		*processOutputTail += string(newRune)

		// Constrain output buffer to 2K to which prevents
		// memory explosion on lots of unexpected data
		if len(*processOutputTail) > 1<<11 {
			// When we want to shrink, we pop off 64 bytes at a time to reduce the frequency of moves/copies
			*processOutputTail = (*processOutputTail)[1<<6:]
		}

		if strings.HasSuffix(*processOutputTail, expectedString) {
			tflog.Trace(ctx, "hit on waiting for line")
			break
		}
	}
}

func spawnProcessor(ctx context.Context, out io.Reader, in io.WriteCloser, processor sshIOProcessor) ( /* hadAnError */ <-chan error /* processOutputTail */, *string) {
	processOutputTail := ""
	processResult := make(chan error)
	go func() {
		outBuffered := bufio.NewReader(out) // Enable buffering, and also gives us utf8 rune decoding
		var errHad error

		waitFor := func(s string) {
			waitForLineContainingText(ctx, s, &errHad, &processOutputTail, outBuffered)
		}
		processor(waitFor, in)
		in.Close()
		processResult <- errHad
	}()

	return processResult, &processOutputTail
}

func executeSSHCommandAndReturnOutput(client *ssh.Client, cmd string) ([]string, error) {
	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to start a new session: %+v", err)
	}
	defer session.Close()

	out, err := session.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to setup stdout for session: %+v", err)
	}
	lines := make([]string, 0)
	go func() {
		reader := bufio.NewReader(out)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}
			lines = append(lines, string(line))
		}
	}()
	err = session.Run(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to execute command %q: %+v", cmd, err)
	}
	return lines, nil
}

func executeSSHCommandAndReturnCombinedOutput(client *ssh.Client, cmd string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to start a new session: %+v", err)
	}
	defer session.Close()

	out, err := session.CombinedOutput(cmd)

	return string(out), err
}

func executeSSHCommandWithInputProcessing(ctx context.Context, client *ssh.Client, cmd string, processor sshIOProcessor) error {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to start a new session: %+v", err)
	}
	defer session.Close()
	closeOnCancel(ctx, session)

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		session.Close()
		return fmt.Errorf("request for pseudo terminal failed: %+v", err)
	}

	in, err := session.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to setup stdin for session: %+v", err)
	}
	defer in.Close()

	out, err := session.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to setup stdout for session: %+v", err)
	}

	errHad, processOutputTail := spawnProcessor(ctx, out, in, processor)
	tflog.Trace(ctx, "run")
	err = session.Run(cmd)
	if err != nil {
		return fmt.Errorf("failed to execute command %q: %w  processOutputTail: %s", cmd, err, *processOutputTail)
	}

	err = <-errHad
	if err != nil {
		return fmt.Errorf("failed to execute command %q: %w  processOutputTail: %s", cmd, err, *processOutputTail)
	}
	return nil
}

func pureuserPublicKeyAuth(pureuserPrivateKey []byte) (ssh.AuthMethod, error) {
	key, err := ssh.ParsePrivateKey(pureuserPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %+v", err)
	}
	return ssh.PublicKeys(key), nil
}
