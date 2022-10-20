package auth_test

import (
	"fmt"
	"testing"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/auth"
	"golang.org/x/crypto/ssh"
)

func TestPrivateKeyDerivePublicKeyMalformed(t *testing.T) {
	_, err := auth.PrivateKeyDerivePublicKey([]byte("adfasdfadsfadsasdfs"))
	if err == nil {
		panic(fmt.Errorf("expected error"))
	}
}
func TestPrivateKeyDerivePublicKeySuccess(t *testing.T) {
	derivedPubBytes, err := auth.PrivateKeyDerivePublicKey([]byte(key0Pvt))
	if err != nil {
		panic(err)
	}

	derivedPub, _, _, _, err := ssh.ParseAuthorizedKey(derivedPubBytes)
	if err != nil {
		panic(err)
	}

	knownPub, _, _, _, err := ssh.ParseAuthorizedKey([]byte(key0Pub))
	if err != nil {
		panic(err)
	}

	if ssh.FingerprintSHA256(derivedPub) != ssh.FingerprintSHA256(knownPub) {
		t.Error()
	}
}

const key0Pub = "" +
	"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC3E//FTyvCFb2tFq2s+aYJiEsS" +
	"UEFtIKpXbAMqiOAdrr/JJGTKtHOFIOTcsu4R2DSZVqgpZzz8rh9Y/Q4DGW9SFUcF" +
	"EdZK6XYZzYPUEz9WvcuyLelB2SlxN7GvbOiKJOBLLLYq9lDGfpWhSp/3Cdh8rLGg" +
	"ESuXnbmR+hr5POvXAhLKgluu70LUQ7tQDa6N/YOzaK/fFbjg0E9WQkKFhNDWywik" +
	"Sq14hBRG+K0tobpvLv4e3mrHlZE0V7EwqPNFz+k1hQwfSn6cJwrb2oyEsdVOmFem" +
	"WAe++6ruL0PQQS3Th9mjOJMt8ulHLhBygB1rEnOyY47CljoTMETm5LDrpwaX"

const key0Pvt = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAtxP/xU8rwhW9rRatrPmmCYhLElBBbSCqV2wDKojgHa6/ySRk
yrRzhSDk3LLuEdg0mVaoKWc8/K4fWP0OAxlvUhVHBRHWSul2Gc2D1BM/Vr3Lsi3p
QdkpcTexr2zoiiTgSyy2KvZQxn6VoUqf9wnYfKyxoBErl525kfoa+Tzr1wISyoJb
ru9C1EO7UA2ujf2Ds2iv3xW44NBPVkJChYTQ1ssIpEqteIQURvitLaG6by7+Ht5q
x5WRNFexMKjzRc/pNYUMH0p+nCcK29qMhLHVTphXplgHvvuq7i9D0EEt04fZoziT
LfLpRy4QcoAdaxJzsmOOwpY6EzBE5uSw66cGlwIDAQABAoIBAQCaBGsiNM6dQukF
GpUVdhim4FA3oejIw7hKP4YSXIAeuIqDzug0V8bvcpLW0HjT7k2hynNCEaYO9PVO
SeVl2hn6tge/Thg3gVxnrKuA2IhqktUwmssnKGhi5PXInRisTwWaeIzwa7PcqVV8
EWmtzEyh5i6weIFx1bDuC7hxgTzu7VUyqfXmWqaT/GhNrmgVOUPpY+WhLNWpz29U
D7pIHvSSGXYjonNC0QaHsu2Ju1d5CKFuzdYLEEnssKCTKEBEBsubVBV0T9E2P5fI
OdH5gRAfH47gEST2RoUyRrfCXEo0papC6NopNJaocC70Hrnv/CALRI7fb0HMoaYY
TEMLq3hhAoGBAN8F86W8vU3mIE/II5FIta2BaNNuNWAYgGcK8ve+CAai3A3D+vqe
ofQgA7n/9BPgFdndxG2huF4ELbX6trI4tEzSCNKA5VJl9zKx/crZqAnnjxYfll/l
IZRyP0dfN4LhQtXQ0IwxB3XWhYWeHfAhDSoJ/aQA5wfXrl411CmbTV/ZAoGBANIm
A3eC/xcoPTwtxtrPIcaGp+Ba1mDntr/todtCP7DkcPPGuXCAsduKV9Xza5TDrX2h
P2jNPeU0hEurb1FPfetLjjM1kE0fYoefcww5Fqh05v+J26mLN221TONCzNt7fKx9
Q0taywqz2fCZPU8zNBnFM7H16SzYyRs8FzCkwwPvAoGBAISeamsk18EJ1i77CUNs
ZDR/npETmQPCriAYH7D4PJeNoqNA5e7AA/hCTVT+geqLtxKt0A+NOjAV9gSfyKIk
G/sObpaWdUdRQPTRaOSGF2mEW87BC8+MVjKK2Vwcn8rhHrg8irbCtPf6j72L76uK
S+SyujQbIUbFl3eINNhk9FbJAoGAH2Ehv2TGJnURv7yuJtorFvPgeXXW8R1ognw/
YPghfo5990jDQ/NTm46q0v3IIKyfaVT8nO+YwHL5FjuBIckW6l2b3DWxWHLJSasb
iMW1hm3+WEYDkOuC2VTrKDw/Tr6/vRnvRH9INk7oyi53oy3oJ7j0oSwxJ0svrGtq
ow7XN+0CgYEAnHuBqqaVNTGbVrxF71bk99FdiMe5PZ2bQOtQUQswlsflmQzesRY3
PD7oxKlqhCYCyjDctUpZsCK02NPMtiHOk1fxtU7nKi1yiJFRbHFXp2X0ucq3xtKv
R95QQozRL4ZSk/4P1PTLBtpRBPsfqNTihZWOoHcMEENXm9i6qr5lpdw=
-----END RSA PRIVATE KEY-----`
