package cmd

import (
	"flag"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// RootCmd is the Cobra root command for clam-scanner.  All subcommands will be
// added to RootCmd.
var RootCmd = &cobra.Command{Use: "clam-scanner"}

func init() {
	// Add flags set with the standard flag library to the pflag library's
	// flag set.
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}
