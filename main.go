package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"

	"transwarp/tos-app-market/cmd"
)

var globalUsage = `The Tos App Market Mirco Service

To begin working with tos-app-market, run the 'tos-app-market serv' command:

    $ tos-app-market serv
`

func newRootCmd(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:          "tos-app-market",
		Short:        "Tos App Market Mirco Service.",
		Long:         globalUsage,
		SilenceUsage: true,
	}
	flags := rootCmd.PersistentFlags()

	rootCmd.AddCommand(
		cmd.NewServCmd(),
		cmd.NewVersionCmd(),
	)

	flags.Parse(args)

	return rootCmd
}

func main() {
	logrus.Infof("TOS App Market Start...")
	rootCmd := newRootCmd(os.Args[1:])
	if err := rootCmd.Execute(); err != nil {
		logrus.Errorln(err)
		os.Exit(1)
	}
}