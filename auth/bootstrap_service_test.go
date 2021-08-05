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
	"context"
	"io"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tfsdklog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"golang.org/x/crypto/ssh"
)

func setupTestCtx(t *testing.T) context.Context {
	ctx := context.Background()

	// This is needed to make tflog work at early stages of unit tests or
	// for tests that don't use the terraform helpers
	ctx = tfsdklog.RegisterTestSink(ctx, t)
	ctx = tfsdklog.NewRootProviderLogger(ctx)
	return ctx
}

func testSetupSSHToRealArray(ctx context.Context, t *testing.T) *ssh.Client {
	privateKeyPath := os.Getenv("TF_TEST_PRIVATE_KEY_PATH")
	if privateKeyPath == "" {
		t.Skipf("TF_TEST_PRIVATE_KEY_PATH not set, skipping test")
	}
	host := os.Getenv("TF_TEST_HOST_ADDRESS")
	if host == "" {
		t.Skipf("TF_TEST_HOST_ADDRESS not set, skipping test")
	}

	privateKeyBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		t.Errorf("Failed to read privateKeyPath: %s, err: %s", privateKeyPath, err)
	}

	client, err := sshSetup(ctx, host, privateKeyBytes)
	if err != nil {
		t.Errorf("SSH Client setup failed: %s", err)
	}
	return client
}

func Test__ExecuteSSHPureAdminCreate(t *testing.T) {
	t.Parallel()
	ctx := setupTestCtx(t)
	client := testSetupSSHToRealArray(ctx, t)

	adminUser := acctest.RandomWithPrefix("test-user-")
	adminUserPwd := acctest.RandomWithPrefix("test-password-")

	err := executeSSHPureAdminCreate(ctx, client, adminUser, adminUserPwd)
	if err != nil {
		t.Errorf("command failed %s", err)
	}

}

func Test__ExecuteSSHPureAPIClientCreate(t *testing.T) {
	t.Parallel()
	ctx := setupTestCtx(t)
	client := testSetupSSHToRealArray(ctx, t)

	apiClientName := acctest.RandomWithPrefix("api-client-")

	_, pub, err := generateKeyPair()
	if err != nil {
		t.Errorf("auth.GenerateKeyPair err:%s", err)
	}

	t.Log(string(pub))

	issuer, clientID, keyID, err := executeSSHPureAPIClientSetup(ctx, client, apiClientName, string(pub))
	if err != nil {
		t.Errorf("command failed %s", err)
	}
	if issuer == "" {
		t.Errorf("missing issuer")
	}
	if clientID == "" {
		t.Errorf("missing clientID")
	}
	if keyID == "" {
		t.Errorf("missing keyID")
	}

}

func processInteractiveTestHelper(t *testing.T, processor sshIOProcessor, simulateSSHRemote func(assertStringWritten func(checkValue string), outWriter *os.File)) {
	ctx, cancel := context.WithTimeout(setupTestCtx(t), time.Second*10)
	defer cancel()

	out, outWriter, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe err: %s", err)
	}
	inReader, in, err := os.Pipe()
	if err != nil {
		t.Errorf("os.Pipe err: %s", err)
	}
	closeOnCancel(ctx, out)
	closeOnCancel(ctx, in)
	closeOnCancel(ctx, outWriter)
	closeOnCancel(ctx, inReader)

	inReaderByte := make(chan byte)
	inReaderErr := make(chan error)
	go func() {
		for {
			newData := [1]byte{}
			readCount, err := inReader.Read(newData[:])
			if err != nil {
				inReaderErr <- err
				break
			}
			if readCount > 0 {
				inReaderByte <- newData[0]
			}
		}
	}()

	assertStringWritten := func(checkValue string) {
		timeoutTimer := time.NewTimer(time.Millisecond * 20)
		inReaderData := []byte{}
	loop:
		for {
			select {
			case newData := <-inReaderByte:
				inReaderData = append(inReaderData, newData)
			case err := <-inReaderErr:
				if err == io.EOF {
					break loop
				} else {
					t.Fatalf("got inReaderErr err: %s", err)
					break loop
				}
			case <-timeoutTimer.C:
				break loop
			}
		}
		inReaderString := string(inReaderData)
		if checkValue != inReaderString {
			t.Errorf("expected: %s, got: %s", checkValue, inReaderString)
		}
	}

	errHad, _ := spawnProcessor(ctx, out, in, processor)

	simulateSSHRemote(assertStringWritten, outWriter)

	err = <-errHad
	if err != nil {
		t.Errorf("had an error: %s", err)
	}

}

func Test__ProcessInteractivePasswordPromptBasic(t *testing.T) {
	t.Parallel()
	ctx := setupTestCtx(t)
	processInteractiveTestHelper(t, processPasswordPrompt(ctx, "testPassword"), func(assertStringWritten func(checkValue string), outWriter *os.File) {
		assertStringWritten("")
		outWriter.WriteString("Enter password:")
		assertStringWritten("testPassword\n")
		outWriter.WriteString("Retype password:")
		assertStringWritten("testPassword\n")
		outWriter.WriteString(".... Name ....")
		assertStringWritten("")
	})

}

func Test__ProcessInteractivePasswordPromptWithJunk(t *testing.T) {
	t.Parallel()
	ctx := setupTestCtx(t)
	processInteractiveTestHelper(t, processPasswordPrompt(ctx, "testPassword"), func(assertStringWritten func(checkValue string), outWriter *os.File) {
		assertStringWritten("")
		outWriter.WriteString("junk0\n")
		assertStringWritten("")
		outWriter.WriteString("Enter password:")
		assertStringWritten("testPassword\n")
		outWriter.WriteString("junk1\n")
		outWriter.WriteString("Retype password:")
		assertStringWritten("testPassword\n")
		outWriter.WriteString(".... Name ....")
		assertStringWritten("")
	})
}

func Test__ProcessInteractivePublicKeyPrompt(t *testing.T) {
	t.Parallel()
	ctx := setupTestCtx(t)
	processInteractiveTestHelper(t, processPublicKeyPrompt(ctx, "testKey0\nline\nline\nline\n"), func(assertStringWritten func(checkValue string), outWriter *os.File) {
		assertStringWritten("")
		outWriter.WriteString("Please enter public key followed by Enter and then Ctrl-D:\n")
		assertStringWritten("testKey0\nline\nline\nline\n\n\x04")
		outWriter.WriteString(".... Name ....")
		assertStringWritten("")
	})
}
