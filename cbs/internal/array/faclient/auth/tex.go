package auth

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type TokenExchanger interface {
	ExchangeToken(string) (string, error)
}

// implements the TokenExchanger interface
type Tex struct {
	TexHost string
}

const (
	oauthEndpoint         = "/oauth2/1.0/token"
	oauthGrantType        = "urn:ietf:params:oauth:grant-type:token-exchange"
	oauthSubjectTokenType = "urn:ietf:params:oauth:token-type:jwt"
)

// TexResponse represents the response for the token exchange request.
type TexResponse struct {
	AccessToken     string `json:"access_token,omitempty"`
	IssuedTokenType string `json:"issued_token_type,omitempty"`
	TokenType       string `json:"token_type,omitempty"`
	ExpiresIn       uint32 `json:"expires_in,omitempty"`
}

// Exchanges an identity token for an access token by calling the OAuth endpoint on the array
func (t *Tex) ExchangeToken(identityToken string) (string, error) {

	//TODO make cert verification a parameter
	httpClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	texURL := url.URL{
		Scheme: "https",
		Host:   t.TexHost,
		Path:   oauthEndpoint,
	}

	resp, err := httpClient.PostForm(texURL.String(), url.Values{
		"grant_type":         {oauthGrantType},
		"subject_token":      {identityToken},
		"subject_token_type": {oauthSubjectTokenType},
	})
	if err != nil {
		return "", fmt.Errorf("Error exchanging tokens: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Array returned HTTP response code %d when attempting to exchange tokens: %s",
			resp.StatusCode, resp.Status)
	}

	decoder := json.NewDecoder(resp.Body)
	texResponse := new(TexResponse)
	err = decoder.Decode(texResponse)
	if err != nil {
		return "", fmt.Errorf("Error decoding token exchange response: %s", err)
	}

	return texResponse.AccessToken, nil
}
