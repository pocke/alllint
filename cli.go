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
	res := make([]Failure, 0)
	// TODO: parallel
	for _, cmd := range c.Commands {
		f, err := cmd.Exec()
		if err != nil {
			return nil, err
		}
		res = append(res, f...)
	}
	return res, nil
}

type Commands []Command

// for Impement pflag.Value

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
