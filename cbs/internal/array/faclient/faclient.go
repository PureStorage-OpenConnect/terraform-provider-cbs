package faclient

import (
	"crypto/tls"
	"fmt"
	"net/http"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/auth"
)

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
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	transport := httptransport.NewWithClient(cfg.Host, cfg.BasePath, cfg.Schemes, &httpClient)
	// TODO implement token refreshing
	transport.DefaultAuthentication = httptransport.BearerToken(accessToken)

	return client.New(transport, strfmt.Default), nil
}
