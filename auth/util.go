package auth

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func PrivateKeyDerivePublicKey(pvtKeyBytes []byte) ([]byte, error) {
	pvtKey, err := ssh.ParsePrivateKey(pvtKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %+v", err)
	}
	return ssh.MarshalAuthorizedKey(pvtKey.PublicKey()), nil
}
