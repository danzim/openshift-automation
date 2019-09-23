package main

import (
	"fmt"
	"net/http"
)

func setHeaders(req *http.Request, token string, contentType string) *http.Request {
	auth := fmt.Sprintf("bearer %s", token)
	accept := "application/json"

	req.Header.Add("Authorization", auth)
	req.Header.Add("Accept", accept)
	req.Header.Add("Content-Type", contentType)

	return req
}
