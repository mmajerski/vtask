package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

const (
	noFileMsg     = "Please check help and provide a HTML file as an argument with --file flag"
	cannotReadMsg = "Cannot read provided file(file might not exist)"
	successMsg    = "Server with file has started on http://localhost:9090 (^C to exit)"
)

// NewRunCmd returns run command to be executed
func NewRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Starts the HTTP web server and serves HTML file provided as an argument",
		RunE: func(cmd *cobra.Command, args []string) error {
			// parse the --file flag
			file, err := cmd.Flags().GetString("file")
			if err != nil {
				return handleResponse(cmd, "There was an error during parsing flag")
			}
			// file is required
			if file == "" {
				return handleResponse(cmd, noFileMsg)
			}
			// read file from argument provided as a string
			data, err := ioutil.ReadFile(file)
			if err != nil {
				return handleResponse(cmd, cannotReadMsg)
			}
			// serve that file on root endpoint
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				w.Write(data)
			})
			fmt.Fprintln(cmd.OutOrStdout(), successMsg)
			if !isTest {
				// start server
				err = http.ListenAndServe(":9090", nil)
				if err != nil {
					return handleResponse(cmd, "Server closed because of an error")
				}
			}
			return err
		},
	}
	// enable the --file flag
	cmd.Flags().StringP("file", "f", "", "HTML file")
	return cmd
}

func init() {
	rootCmd.AddCommand(NewRunCmd())
}
