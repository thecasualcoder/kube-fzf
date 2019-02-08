package fzf

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/thecasualcoder/cmd"
)

type inputFunc func(io.WriteCloser)

func newFzfCmd() *cmd.Cmd {
	defaultFzfOpts := cmd.Opts([]cmd.Opt{
		{
			Type:  "string",
			Flag:  "--height",
			Value: "12",
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
			Flag:  "--reverse",
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

// Filter interactively filters one or more items from a list
func Filter(query string, multi bool, input inputFunc) []string {
	fzfCmd := newFzfCmd()
	fzfCmd.OverrideOpt("--query", query)
	if multi {
		fzfCmd.OverrideOpt("--multi", "true")
	}
	shell := os.Getenv("SHELL")
	if len(shell) == 0 {
		shell = "sh"
	}
	command := exec.Command(shell, "-c", fzfCmd.String())
	command.Stderr = os.Stderr
	in, _ := command.StdinPipe()
	go func() {
		input(in)
		in.Close()
	}()
	result, _ := command.Output()
	filtered := strings.Split(string(result), "\n")
	filtered = filtered[:len(filtered)-1]
	return filtered
}
