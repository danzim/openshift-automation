package main

import (
	"crypto/tls"
	"net/http"
)

func skipTLS() *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	return client
}
