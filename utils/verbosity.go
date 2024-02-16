package utils

import (
	"github.com/spf13/pflag"
)

var Verbose bool

func init() {
	pflag.BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
