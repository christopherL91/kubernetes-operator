package operator

import (
	"github.com/christopherL91/k8s-operator/pkg/controller"
	"github.com/christopherL91/k8s-operator/pkg/operator/options"
)

func Run(options *options.OperatorOptions) error {
	c := controller.New()
	c.Run()
	return nil
}
