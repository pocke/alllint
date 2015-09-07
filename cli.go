package main

import (
	"strings"

	"github.com/ogier/pflag"
)

type CLI struct {
	Commands Commands
}

func (c *CLI) Parse(args []string) error {
	cmds := make(Commands, 0)
	c.Commands = cmds

	fset := pflag.NewFlagSet(args[0], pflag.ContinueOnError)
	fset.VarP(&c.Commands, "command", "c", "command")

	return fset.Parse(args[1:])
}

func (c *CLI) Exec() ([]Failure, error) {
	// TODO
	return nil, nil
}

type Commands []Command

func (v *Commands) String() string {
	cmds := make([]string, 0, len(*v))
	for _, c := range *v {
		cmds = append(cmds, strings.Join(c, " "))
	}
	return strings.Join(cmds, ", ")
}

func (v *Commands) Set(s string) error {
	cmd := strings.Split(s, " ")
	*v = append(*v, cmd)
	return nil
}
