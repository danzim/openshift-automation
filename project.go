package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
)

func project(paramsProject ParamsProject, apiURL string, method string, contentType string) string {
	//apiURL := fmt.Sprintf("%soapi/v1/projectrequests", baseURL)
	token := getToken()

	var buffer bytes.Buffer
	var data io.Reader

	cache, _ := ioutil.ReadFile("templates/project.json")
	projectTemplate := string(cache)
	tmplProject := template.Must(template.New("projectTemplate").Parse(projectTemplate))
	tmplProject.Execute(&buffer, paramsProject)
	data = &buffer
	s := buffer.String()
	fmt.Printf(s)

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
	return message
}
