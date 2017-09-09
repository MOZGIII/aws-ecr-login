package ael // import "github.com/MOZGIII/aws-ecr-login"

import (
	"fmt"
)

func do() error {
	authorizationData, err := getECRAuthorizationData()
	if err != nil {
		return err
	}

	tmpl, err := parseTemplate(DockerLoginTemplate)
	if err != nil {
		return err
	}

	str, err := evaluateTemplate(authorizationData, tmpl)
	if err != nil {
		return err
	}

	fmt.Println(str)
	return nil
}

// RunCLI executes the app.
func RunCLI() error {
	return do()
}
