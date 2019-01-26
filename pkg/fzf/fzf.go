package fzf

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/thecasualcoder/cmd"
)

func defaultFzfCmd() *cmd.Cmd {
	defaultFzfOpts := cmd.Opts([]cmd.Opt{
		{
			Type:  "bool",
			Flag:  "--height",
			Value: "10",
		},
		{
			Type:  "bool",
			Flag:  "--ansi",
			Value: "true",
		},
		{
			Type:  "bool",
			Flag:  "--select-1",
			Value: "true",
		},
		{
			Type:  "bool",
			Flag:  "--multi",
			Value: "false",
		},
		{
			Type:  "string",
			Flag:  "--query",
			Value: "",
		},
	})

	return cmd.New("fzf", defaultFzfOpts)
}

func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func withFilter(command string, input func(in io.WriteCloser)) []string {
	shell := os.Getenv("SHELL")
	if len(shell) == 0 {
		shell = "sh"
	}
	cmd := exec.Command(shell, "-c", command)
	cmd.Stderr = os.Stderr
	in, _ := cmd.StdinPipe()
	go func() {
		input(in)
		in.Close()
	}()
	result, _ := cmd.Output()
	filtered := strings.Split(string(result), "\n")
	filtered = filtered[:len(filtered)-1]
	return filtered
}

func inputFunc(items []string) func(io.WriteCloser) {
	return func(in io.WriteCloser) {
		for _, item := range items {
			fmt.Fprintln(in, item)
		}
	}
}

// Filter interactively filters one or more items from a list
func Filter(query string, multi bool, items []string) []string {
	fzfCmd := defaultFzfCmd()
	fzfCmd.OverrideOpt("--query", query)
	if multi {
		fzfCmd.OverrideOpt("--multi", "true")
	}

	return withFilter(fzfCmd.String(), inputFunc(items))
}
