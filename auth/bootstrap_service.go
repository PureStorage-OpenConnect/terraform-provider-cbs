// +build !mock

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
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net"
	"strings"

	"golang.org/x/crypto/ssh"
)

type bootstrapService struct{}

func NewBootstrapService() Bootstrapper {
	return &bootstrapService{}
}

func (b *bootstrapService) GenerateSecretPayload(host string, pureuserPrivateKey []byte) ([]byte, error) {
	privatePem, publicPem, err := generateKeyPair()
	if err != nil {
		return nil, err
	}

	authMethod, err := pureuserPublicKeyAuth(pureuserPrivateKey)
	if err != nil {
		return nil, err
	}

	sshConfig := &ssh.ClientConfig{
		User:            "pureuser",
		Auth:            []ssh.AuthMethod{authMethod},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := net.Dial("tcp", net.JoinHostPort(host, "22"))
	if err != nil {
		return nil, fmt.Errorf("dial error: %+v", err)
	}

	sshConn, chans, reqs, err := ssh.NewClientConn(conn, host, sshConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create new ssh client session: %+v", err)
	}

	client := ssh.NewClient(sshConn, chans, reqs)

	adminUserPwd, err := generateAdminUserPwd()
	if err != nil {
		return nil, fmt.Errorf("failed to generate a random password: %+v", err)
	}

	const adminUser = "orchestration-user"
	const apiClient = "orchestration-client"
	if err = executeCmdPwdInput(client, "pureadmin create --role array_admin "+adminUser, adminUserPwd); err != nil {
		return nil, err
	}

	if err = executeCmdPubKeyInput(client, "pureapiclient create --max-role array_admin --public-key "+apiClient, string(publicPem)); err != nil {
		return nil, err
	}

	if _, err := executeCmd(client, "pureapiclient enable "+apiClient); err != nil {
		return nil, err
	}

	output, err := executeCmd(client, "pureapiclient list "+apiClient+" --csv")
	if err != nil {
		return nil, err
	}
	if len(output) != 2 {
		return nil, fmt.Errorf("got an unexpected number of outputs from command: pureapiclient list " + apiClient)
	}
	fields := strings.Split(output[0], ",")
	values := strings.Split(output[1], ",")

	apiClientMap := make(map[string]string)
	for i := 0; i < len(fields); i++ {
		apiClientMap[fields[i]] = values[i]
	}

	payload := SecretPayload{
		UserName:       adminUser,
		Issuer:         apiClientMap["Issuer"],
		ClientID:       apiClientMap["Client ID"],
		KeyID:          apiClientMap["Key ID"],
		RestPrivateKey: string(privatePem),
	}
	credentials, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return credentials, nil
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

func executeCmd(client *ssh.Client, cmd string) ([]string, error) {
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

func executeCmdPubKeyInput(client *ssh.Client, cmd string, input string) error {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to start a new session: %+v", err)
	}
	defer session.Close()

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

	out, err := session.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to setup stdout for session: %+v", err)
	}

	go func(in io.WriteCloser, out io.Reader) {
		reader := bufio.NewReader(out)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}

			if string(line) == "Please enter public key followed by Enter and then Ctrl-D:" {
				in.Write([]byte(input + "\n\x04"))
				in.Close()
				break
			}
		}
	}(in, out)
	err = session.Run(cmd)
	if err != nil {
		return fmt.Errorf("failed to execute command %q: %+v", cmd, err)
	}

	return nil
}

func executeCmdPwdInput(client *ssh.Client, cmd string, password string) error {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to start a new session: %+v", err)
	}
	defer session.Close()

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

	out, err := session.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to setup stdout for session: %+v", err)
	}

	go func(in io.WriteCloser, out io.Reader) {
		reader := bufio.NewReader(out)
		var line []byte

		for {
			b, err := reader.ReadByte()
			if err != nil {
				break
			}
			line = append(line, b)

			if string(line) == "Enter password:" {
				in.Write([]byte(password + "\n"))
				line = line[:0]
			}
			if strings.Contains(string(line), "Retype password:") {
				in.Write([]byte(password + "\n"))
				in.Close()
				break
			}
		}
	}(in, out)
	err = session.Run(cmd)
	if err != nil {
		return fmt.Errorf("failed to execute command %q: %+v", cmd, err)
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
