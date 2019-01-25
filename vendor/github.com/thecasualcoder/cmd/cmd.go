package cmd

import (
	"fmt"
	"sort"
	"strings"
)

// Cmd represents a command and its options
type Cmd struct {
	Name string
	Opts
}

// String converts Cmd to a string suitable for executing it in a shell
func (cmd Cmd) String() string {
	var buffer strings.Builder

	buffer.WriteString(cmd.Name)
	buffer.WriteString(" ")
	buffer.WriteString(cmd.Opts.String())

	return buffer.String()
}

// OverrideOpt override value for opt
func (cmd *Cmd) OverrideOpt(flag, value string) {
	cmd.Opts.Override(flag, value)
}

// Opts represents a collection of command options
type Opts []Opt

// String converts Opts to string that can used with a command when executing it in a shell
func (opts Opts) String() string {
	var buffer strings.Builder

	sort.Slice(opts, func(i, j int) bool {
		return opts[i].Flag < opts[j].Flag
	})

	for _, opt := range opts {
		optStr := opt.String()

		if strings.TrimSpace(optStr) != "" {
			buffer.WriteString(opt.String())
			buffer.WriteString(" ")
		}
	}

	return strings.TrimSpace(buffer.String())
}

// Override overrides value for opt identified by given flag
func (opts *Opts) Override(flag, value string) {
	for i := range *opts {
		if (*opts)[i].Flag == flag {
			(*opts)[i].Override(value)
			return
		}
	}
}

// Opt represents a command option
type Opt struct {
	Type  string
	Flag  string
	Value string
}

// String converts an Opt to string that can used with a command when executing it in a shell
func (opt Opt) String() string {
	var buffer strings.Builder
	value := strings.TrimSpace(opt.Value)

	if opt.Type == "string" && value != "" {
		buffer.WriteString(opt.Flag)
		buffer.WriteString("=")
		buffer.WriteString(opt.Value)
	} else if opt.Type == "bool" && value == "true" {
		buffer.WriteString(opt.Flag)
	}

	return buffer.String()
}

// Override overrides an opt's value
func (opt *Opt) Override(value string) error {
	if opt.Type == "bool" && (value != "true" && value != "false") {
		return fmt.Errorf("%s: is not a bool value", value)
	}

	opt.Value = value

	return nil
}

// New creates a new command
func New(name string, opts Opts) *Cmd {
	return &Cmd{
		Name: name,
		Opts: opts,
	}
}
