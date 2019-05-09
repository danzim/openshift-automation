package main

import (
	"fmt"
	"net/http"
)

func setHeaders(req *http.Request, token string) *http.Request {
	auth := fmt.Sprintf("bearer %s", token)
	accept := "application/json"
	contentType := accept

	req.Header.Add("Authorization", auth)
	req.Header.Add("Accept", accept)
	req.Header.Add("Content-Type", contentType)

	return req
}
