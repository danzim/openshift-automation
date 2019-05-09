package main

import "net/http"

func rootSite(w http.ResponseWriter, r *http.Request) {
	message := "OpenShift Automated Project"
	w.Write([]byte(message))
}
