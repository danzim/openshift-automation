package main

import (
	"fmt"
	"net/http"
)

const (
	defaultTimeFormat = "2006-01-02 03:04:05.000 MST"
	baseURL           = "https://192.168.64.4:8443/"
	port              = ":8080"
)

func main() {
	setLog("info", "starting the application...")
	http.HandleFunc("/", rootSite)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		setLog("critical", fmt.Sprintf("starting the web server failed: %s", err))
	}
}
