package fzf

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"
)

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
	return strings.Split(string(result), "\n")
}

type fzfOptions struct {
	query string
	multi bool
}

func (opts *fzfOptions) String() string {
	var buffer bytes.Buffer
	if opts.multi {
		buffer.WriteString("-m ")
	}

	if opts.query != "" {
		buffer.WriteString("-q ")
		buffer.WriteString(opts.query)
	}
	return buffer.String()
}

type fzfCmd struct {
	name string
	opts *fzfOptions
}

func (cmd *fzfCmd) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(cmd.name)
	buffer.WriteString(" ")
	buffer.WriteString(cmd.opts.String())
	return buffer.String()
}

func newFzfCmd(opts *fzfOptions) *fzfCmd {
	return &fzfCmd{
		name: "fzf",
		opts: opts,
	}
}

// FilterOne interactively filters one item from a list
func FilterOne(query string, input func(in io.WriteCloser)) string {
	opts := &fzfOptions{
		query: query,
	}
	cmd := newFzfCmd(opts)
	result := withFilter(cmd.String(), input)

	if len(result) == 0 {
		return ""
	}
	return result[0]
}

// FilterMany interactively filters many items from a list
func FilterMany(query string, input func(in io.WriteCloser)) []string {
	opts := &fzfOptions{
		query: query,
		multi: true,
	}
	cmd := newFzfCmd(opts)
	result := withFilter(cmd.String(), input)
	return result
}
