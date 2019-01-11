package fzf

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type fzfOptions struct {
	query       string
	multi       bool
	height      int
	ansi        bool
	reverse     bool
	selectIfOne bool
}

func (opts *fzfOptions) String() string {
	var buffer bytes.Buffer
	if opts.multi {
		buffer.WriteString("--multi ")
	}

	if opts.ansi {
		buffer.WriteString("--ansi ")
	}

	if opts.reverse {
		buffer.WriteString("--reverse ")
	}

	if opts.selectIfOne {
		buffer.WriteString("--select-1 ")
	}

	if opts.height >= 0 {
		buffer.WriteString("--height ")
		buffer.WriteString(strconv.Itoa(opts.height))
		buffer.WriteString(" ")
	}

	if opts.query != "" {
		buffer.WriteString("--query ")
		buffer.WriteString(opts.query)
		buffer.WriteString(" ")
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

func defaultfzfCmd() *fzfCmd {
	return &fzfCmd{
		name: "fzf",
		opts: &fzfOptions{
			query:       "",
			multi:       false,
			height:      10,
			ansi:        true,
			reverse:     true,
			selectIfOne: true,
		},
	}
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
	return strings.Split(string(result), "\n")
}

func inputFunc(items []string) func(io.WriteCloser) {
	return func(in io.WriteCloser) {
		for _, item := range items {
			fmt.Fprintln(in, item)
		}
	}
}

// FilterOne interactively filters one item from a list
func FilterOne(query string, items []string) string {
	cmd := defaultfzfCmd()
	cmd.opts.query = query
	fmt.Println(cmd)
	result := withFilter(cmd.String(), inputFunc(items))

	if len(result) == 0 {
		return ""
	}
	return result[0]
}

// FilterMany interactively filters many items from a list
func FilterMany(query string, items []string) []string {
	cmd := defaultfzfCmd()
	cmd.opts.query = query
	cmd.opts.multi = true
	result := withFilter(cmd.String(), inputFunc(items))

	return result
}
