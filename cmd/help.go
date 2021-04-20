package cmd

import (
	"github.com/spf13/cobra"
)

const helpMsg = `vtask is a simple CLI.
Usage:
vtask [command]
		
Available Commands:
	version                   Prints current version of the CLI
	help                      Views more information about this program
	run [-f | --file]         Starts the HTTP web server and serves HTML file provided as an argument
`

// NewHelpCmd returns help command to be executed
func NewHelpCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "help",
		Short: "vtask is a simple CLI",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleResponse(cmd, helpMsg)
		},
	}
}

func init() {
	rootCmd.AddCommand(NewHelpCmd())
}
