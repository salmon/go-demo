package cmd

import (
	"github.com/spf13/cobra"

	"transwarp/tos-app-market/pkg/version"
)

const versionDesc = `
This command print version of installer.`

type vCmd struct {
}

func NewVersionCmd() *cobra.Command {
	vc := &vCmd{}

	cmd := &cobra.Command{
		Use:   "version",
		Short: "print version",
		Long:  versionDesc,

		RunE: func(cmd *cobra.Command, args []string) error {
			defer vc.run()
			return nil
		},
	}
	return cmd
}

func (vc *vCmd) run() {
	version.PrintVersionInfo()
}
