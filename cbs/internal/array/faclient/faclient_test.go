package faclient_test

import (
	"testing"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/auth"
)

const testPrivateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAumcxoptOzGLwQuTAWZXovlc0KO79Zp/MYFlsF5XMWS8q2/ff
dpnknyfFNUIy6uEYf3kKHl1D15KuyGZJ2Opiv2nDc1XJAl4LWQU2YlliKrJ2yURq
oxRxQfoW0T/Zi1+v7wqyfxq6uYqUMxitBRRgstIGihTBGgjc7obP9mUp+3MH3NPx
qXT4CVsQDKlDWJC7pR0sZmeGc/YeDanD4iXl9ldk4j45ow8eaO6yeJgRx39hW7kT
NYYrDaWrlwVAt96OQHGmFmNyVIaikvm5tSid4nr3GKioyvU3a4aFIRK9s0iJ4Pa1
Xkae0/M+LVnwqUzSPtOKB4ZOsxOkcG22QV5aQQIDAQABAoIBABSG4qdmdPH6/zxO
loQHBx9W2Q6c6tjPRiFiF44tV9iGYjDhBgG4hr2kEop/5L2K1jjLanrXiG4H4Cl/
Yih5/y/XBMgBrWYOVy+RbGC+ORo8luopF5kn2iUK1lJqtpkri2NGiUuH9QITqahK
0lrZ2KA4krHIAU/NpA68V9BpaYsf+oUek43iZrZyjABp9sv4DOm1JeA61vUFTBPQ
euZCcxbdFo/5+9k2AAEr29o79fnLw8yeAckczke9Xpp++XAC8HQ9/ZjSeCo9NEEE
iAx7T1ahaT2LzRf5K/Uj3n4Fuo11XnMjqDyet46/yk821QB2H0epm+Z3iaMjSCYW
0mnCKbECgYEA6nnkIx5keh9BT5SdHHAKePCJBo/v7Ag6E2Q9yMN/QQu0t+wQZGne
+4P8Xts1arigWg85LOdrAdFq6DYuT3ja1JKwahQ/XlWYqpbLwuHnq3LvQm2pLxYA
gqnb7GPrfKm3bj10Cezy3Zv4EwpDRiXrVSEqlXNsVQV+Pnnusx/n7XUCgYEAy4Oa
VNK4Ey9nPeYL95de0fL4sI35EZNY9yX6n4pv+EOgvrsQSoWQb0JQLyV/Tmw064vb
70Z1bEHz8wWM2BjedqFuKkdzUpQTTZIXS5ozyxmAAS6h4MvBPSuFOO7yqhkkSKJo
McO61FteoR26LvRapt+x1r55iFtvCwkVm5FkJB0CgYA6w8vKhW53MOgkcsGhg+8L
+nTNITvnMvSjMYdOjriQ68ciJVbCY8pPzPduKpBLq/P8Pj59I46tCPg7NIEMx+RI
TG9MVsC++sLlVh/BOu7eCFMwmd1CAMil9r44k55MQxjG1z4C0tDXe6SD2RmdNhmx
3zsV87Sd5l+KdvK9D+0HlQKBgDjiKer3kvfZ0hOdD08/AgPQ0+4VYL6m3sEF3o1l
VnKgBHgLNTx/JKXUdTEYXAMBf7EuwGSa3wtJS/RrYrisCtJBwNcUbYlxVgvif5xk
F4H3OK4b6Kc6jGKanXwSXcVpjZi3vEPcn4XnnAWQl4+0QPpPoBeT2chhNiJxgZag
BsuJAoGBAIt/VO7U1Dea29weR3nk5nNblhKUAG5kcZ82iCUaOg5tPMy+E09ClaxF
6b85KZRT66rSlVFcPfNETD9ZEdwrtM/YTNmhsUf3KdW+p1IbHF3VwhHWRwrrSkyz
HYz83jsJVBLVMwmUltHmsGLKSGwpaAJHqHh7OOgka9bmiuGnSDD/
-----END RSA PRIVATE KEY-----`

const testPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAumcxoptOzGLwQuTAWZXo
vlc0KO79Zp/MYFlsF5XMWS8q2/ffdpnknyfFNUIy6uEYf3kKHl1D15KuyGZJ2Opi
v2nDc1XJAl4LWQU2YlliKrJ2yURqoxRxQfoW0T/Zi1+v7wqyfxq6uYqUMxitBRRg
stIGihTBGgjc7obP9mUp+3MH3NPxqXT4CVsQDKlDWJC7pR0sZmeGc/YeDanD4iXl
9ldk4j45ow8eaO6yeJgRx39hW7kTNYYrDaWrlwVAt96OQHGmFmNyVIaikvm5tSid
4nr3GKioyvU3a4aFIRK9s0iJ4Pa1Xkae0/M+LVnwqUzSPtOKB4ZOsxOkcG22QV5a
QQIDAQAB
-----END PUBLIC KEY-----`

type testParams struct {
	host       string
	issuer     string
	clientID   string
	keyID      string
	privateKey string
}

var params testParams

// TODO use env vars to specify test params
func setup() {
	params.host = "vm-ejacobs2"
	params.issuer = "test-client"
	params.clientID = "90b1fb8a-f15e-4a59-98ca-b220b53e7e10"
	params.keyID = "aa837c07-d1c8-4c11-9ab6-c307aa913aca"
	params.privateKey = testPrivateKey
}

func TestAuth(t *testing.T) {
	setup()

	tokenConfig := auth.TokenConfig{
		Issuer:     params.issuer,
		ClientID:   params.clientID,
		KeyID:      params.keyID,
		User:       "pureuser",
		PrivateKey: params.privateKey,
	}

	_, err := faclient.New(params.host, tokenConfig)
	if err != nil {
		t.Error(err)
	}
}
