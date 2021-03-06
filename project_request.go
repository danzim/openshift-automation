package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
)

func projectRequest(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var reqData ReqData
	var buffer bytes.Buffer
	var data io.Reader

	err := decoder.Decode(&reqData)
	if err != nil {
		setLog("critical", fmt.Sprintf("decoding request body failed: %s", err))
	}

	paramsProject := ParamsProject{reqData.Name, reqData.Description, reqData.DisplayName}
	contentType := "application/json"
	method := "POST"
	apiURL := fmt.Sprintf("%soapi/v1/projects", baseURL)
	token := getToken()

	cache, _ := ioutil.ReadFile("templates/projectrequest.json")
	projectTemplate := string(cache)
	tmplProject := template.Must(template.New("projectTemplate").Parse(projectTemplate))
	tmplProject.Execute(&buffer, paramsProject)
	data = &buffer
	//s := buffer.String()
	//fmt.Printf(s)
	req, err := http.NewRequest(method, apiURL, data)
	if err != nil {
		setLog("critical", fmt.Sprintf("error reading request: %s", err))
	}

	req = setHeaders(req, token, contentType)
	client := skipTLS()

	setLog("info", fmt.Sprintf("creating project with request %s", apiURL))
	responseData, err := client.Do(req)
	if err != nil {
		setLog("critical", fmt.Sprintf("failed to send request to %s: %s", apiURL, err))
	}

	response, _ := ioutil.ReadAll(responseData.Body)
	message := fmt.Sprintf("Project succesfully created:\n%s", string(response))
	setLog("info", fmt.Sprintf("project %s succesfully created", paramsProject.Name))
	//message := project(paramsProject, apiURL, method, contentType)
	fmt.Fprint(w, message)
}
