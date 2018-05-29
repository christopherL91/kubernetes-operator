package main

import (
	"github.com/christopherL91/k8s-operator/cmd"
	"github.com/pkg/errors"
	"log"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("[main]: %v\n", errors.Wrap(err, "Could not execute"))
		os.Exit(-1)
	}
}
