package ael // import "github.com/MOZGIII/aws-ecr-login"

import (
	"bytes"
	"text/template"

	"github.com/aws/aws-sdk-go/service/ecr"
)

// DockerLoginTemplate renders standard Docker login command
const DockerLoginTemplate = `docker login -u {{ username .AuthorizationToken }} -p {{ password .AuthorizationToken }} {{ .ProxyEndpoint }}`

// JSONTemplate renders simple JSON output
const JSONTemplate = `{{ print . }}`

func parseTemplate(input string) (*template.Template, error) {
	funcMap := template.FuncMap{
		"base64decode": base64decode,
		"split":        split,
		"credentials":  credentials,
		"username":     username,
		"password":     password,
	}
	return template.New("format").Funcs(funcMap).Parse(input)
}

func evaluateTemplate(data *ecr.AuthorizationData, tmpl *template.Template) (string, error) {
	buff := bytes.NewBuffer(make([]byte, 0, 4096))
	if err := tmpl.Execute(buff, data); err != nil {
		return "", err
	}
	return buff.String(), nil
}
