package main

import (
	"log"

	"github.com/openshift/clam-scanner/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatalf(err.Error())
	}
}
