package cmd

import (
	"github.com/spf13/cobra"
)

const versionMsg = "Current version: v0.1"

// NewVersionCmd returns version command to be executed
func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints current version of the CLI",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleResponse(cmd, versionMsg)
		},
	}
}

func init() {
	rootCmd.AddCommand(NewVersionCmd())
}
