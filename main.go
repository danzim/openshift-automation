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

//ReqData - parameters for json request
type ReqData struct {
	Name         string
	DisplayName  string
	Description  string
	Pods         string
	LimitsCPU    string
	LimitsMemory string
}

//ParamsProject - parameters for project creation
type ParamsProject struct {
	Name        string
	Description string
	DisplayName string
}

//ParamsQuota - parameters for quota config
type ParamsQuota struct {
	Name         string
	Pods         string
	LimitsCPU    string
	LimitsMemory string
}

func main() {
	setLog("info", "starting the application...")
	http.HandleFunc("/", rootSite)
	http.HandleFunc("/request/", projectRequest)
	http.HandleFunc("/modify/", modify)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		setLog("critical", fmt.Sprintf("starting the web server failed: %s", err))
	}
}
