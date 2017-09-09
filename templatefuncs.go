package ael // import "github.com/MOZGIII/aws-ecr-login"

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func base64decode(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

func split(input, sep string) ([]string, error) {
	return strings.Split(input, sep), nil
}

func credentials(authorizationToken string) ([]string, error) {
	str, err := base64decode(authorizationToken)
	if err != nil {
		return nil, err
	}
	slice := strings.SplitN(str, ":", 2)
	if len(slice) != 2 {
		return nil, fmt.Errorf("invalid credentials pair")
	}
	return slice, nil
}

func username(authorizationToken string) (string, error) {
	slice, err := credentials(authorizationToken)
	if err != nil {
		return "", err
	}
	return slice[0], nil
}

func password(authorizationToken string) (string, error) {
	slice, err := credentials(authorizationToken)
	if err != nil {
		return "", err
	}
	return slice[1], nil
}
