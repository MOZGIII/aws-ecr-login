package main

import (
	"log"

	ael "github.com/MOZGIII/aws-ecr-login"
)

func main() {
	if err := ael.RunCLI(); err != nil {
		log.Fatal(err)
	}
}
