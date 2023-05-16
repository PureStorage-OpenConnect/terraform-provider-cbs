package faclient

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/auth"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/version"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
)

type customUserAgentTransport struct {
	innerTransport http.RoundTripper
	userAgent      string
}

func (t *customUserAgentTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("User-Agent", t.userAgent)
	return t.innerTransport.RoundTrip(r)
}

func New(host string, tokenConfig auth.TokenConfig) (*client.Flasharray, error) {

	identityToken, err := tokenConfig.BuildToken()
	if err != nil {
		return nil, fmt.Errorf("Error creating client: %s", err)
	}

	tex := auth.Tex{TexHost: host}
	accessToken, err := tex.ExchangeToken(identityToken)
	if err != nil {
		return nil, fmt.Errorf("Error exchanging tokens: %s", err)
	}

	cfg := client.DefaultTransportConfig().WithHost(host).WithSchemes([]string{"https"})
	httpClient := http.Client{
		Transport: &customUserAgentTransport{
			innerTransport: logging.NewLoggingHTTPTransport(
				&http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				}),
			userAgent: fmt.Sprintf("%s/%s", "terraform-provider-cbs", version.ProviderVersion),
		},
	}
	transport := httptransport.NewWithClient(cfg.Host, cfg.BasePath, cfg.Schemes, &httpClient)
	// TODO implement token refreshing
	transport.DefaultAuthentication = httptransport.BearerToken(accessToken)

	return client.New(transport, strfmt.Default), nil
}
