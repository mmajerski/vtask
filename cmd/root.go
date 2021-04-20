package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const rootMsg = "vtask is a simple CLI"

var isTest bool = false    // for testing (used primarily for preventing server to start and to terminate test)
var rootCmd = NewRootCmd() // rootCmd is used as a parent, every other commands are just added to it

// runCmd represents the root command
func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "vtask",
		Short: "vtask is a simple CLI",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleResponse(cmd, rootMsg)
		},
	}
}

// Execute executes command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// handleReponse takes command and message and puts the message to stdout or redirects it to the other buffer
func handleResponse(cmd *cobra.Command, msg string) error {
	fmt.Fprintln(cmd.OutOrStdout(), msg)
	return nil
}
