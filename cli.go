package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ogier/pflag"
)

type CLI struct {
	Commands Commands
	Files    []string
}

func (c *CLI) Parse(args []string) error {
	cmds := make(Commands, 0)
	c.Commands = cmds

	fset := pflag.NewFlagSet(args[0], pflag.ContinueOnError)
	fset.VarP(&c.Commands, "command", "c", "command")

	err := fset.Parse(args[1:])
	c.Files = fset.Args()
	return err
}

func (c *CLI) Exec() ([]Failure, error) {
	if len(c.Commands) == 0 {
		return nil, fmt.Errorf("cmd is required")
	}

	errCh := make(chan error)
	failureCh := make(chan []Failure)
	wg := &sync.WaitGroup{}
	res := make([]Failure, 0)

	for _, cmd := range c.Commands {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f, err := cmd.Exec(c.Files)
			if err != nil {
				errCh <- err
				return
			}
			failureCh <- f
		}()
	}

	waitCh := wgCh(wg)
	for {
		select {
		case <-waitCh:
			return res, nil
		case err := <-errCh:
			return nil, err
		case f := <-failureCh:
			res = append(res, f...)
		}
	}
}

func wgCh(wg *sync.WaitGroup) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		wg.Wait()
		ch <- struct{}{}
	}()
	return ch
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
