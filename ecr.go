package ael // import "github.com/MOZGIII/aws-ecr-login"

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecr"
)

func getECRAuthorizationData() (*ecr.AuthorizationData, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	svc := ecr.New(sess)

	input := &ecr.GetAuthorizationTokenInput{}
	result, err := svc.GetAuthorizationToken(input)
	if err != nil {
		return nil, err
	}
	if len(result.AuthorizationData) != 1 {
		return nil, fmt.Errorf("unexpected response cardinality %d", len(result.AuthorizationData))
	}
	authorizationData := result.AuthorizationData[0]
	return authorizationData, nil
}
