package main

import (
	"strings"

	"github.com/ogier/pflag"
)

type CLI struct {
	Commands []string
}

func (c *CLI) Parse(args []string) error {
	cmds := StringSlice(c.Commands)
	fset := pflag.NewFlagSet(args[0], pflag.ContinueOnError)
	fset.VarP(&cmds, "commands", "c", "commands")

	return fset.Parse(args[1:])
}

type StringSlice []string

func (v *StringSlice) String() string {
	return strings.Join(*v, ", ")
}

func (v *StringSlice) Set(s string) error {
	*v = append(*v, s)
	return nil
}
