package options

import (
	"fmt"
	"github.com/spf13/pflag"
	"strings"
)

type OperatorOptions struct {
	Kubeconfig  string
	MasterURL   string
	Threadiness int
}

func New() *OperatorOptions {
	return &OperatorOptions{}
}

func (o *OperatorOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.MasterURL, "master", "",
		`The url of the Kubernetes API server,
		 will overrides any value in kubeconfig, only required if out-of-cluster.`)
	fs.IntVar(&o.Threadiness, "threadiness", 2,
		`How many threads to process the main logic`)
}

func (o *OperatorOptions) String() string {
	return strings.Replace(fmt.Sprintf("%#v", o), ", ", "\n", -1)
}
