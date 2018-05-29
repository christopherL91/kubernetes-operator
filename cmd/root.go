package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "operator-k8s",
	Short:            "This is an example of a simple k8s operator",
	TraverseChildren: true,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return errors.Wrap(err, "Could not execute command")
	}
	return nil
}
