package cmd

import (
	"github.com/christopherL91/k8s-operator/pkg/operator"
	op "github.com/christopherL91/k8s-operator/pkg/operator/options"
	"github.com/spf13/cobra"
)

var (
	options = op.New()
)

func init() {
	rootCmd.AddCommand(operatorCmd)
	options.AddFlags(operatorCmd.Flags())
}

var operatorCmd = &cobra.Command{
	Use:   "operator",
	Short: "simple k8s operator",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cmd.SilenceErrors = true
		cmd.SilenceUsage = true
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return operator.Run(options)
	},
}
