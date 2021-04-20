package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/spf13/cobra"
)

func TestTableTests(t *testing.T) {
	type testStruct struct {
		name string
		cmd  *cobra.Command
		args []string
		want string
	}

	testCases := []testStruct{
		{name: "rootCmd", cmd: NewRootCmd(), args: []string{""}, want: rootMsg},
		{name: "versionCmd", cmd: NewVersionCmd(), args: []string{""}, want: versionMsg},
		{name: "helpCmd", cmd: NewHelpCmd(), args: []string{""}, want: helpMsg},
		{name: "runNoFileProvided", cmd: NewRunCmd(), args: []string{""}, want: noFileMsg},
		{name: "runFileNotExist", cmd: NewRunCmd(), args: []string{"--file", "notexist.html"}, want: cannotReadMsg},
		{name: "runFileSuccess", cmd: NewRunCmd(), args: []string{"--file", "testfiles/index.html"}, want: successMsg},
	}

	for _, tc := range testCases {
		got := runCmd(tc.cmd, tc.args)
		if string(got) != tc.want {
			t.Fatalf("%s: expected \"%s\" got \"%s\"", tc.name, tc.want, string(got))
		}
	}
}

// runCmd executes the command while testing
func runCmd(cmd *cobra.Command, args []string) string {
	isTest = true
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs(args)
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		fmt.Println(err)
	}
	out = out[:len(out)-1] // remove LF from buffer(LF causes problem when compring strings inside tests)
	isTest = false
	return string(out)
}
