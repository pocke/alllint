package main

import "os/exec"

type Command []string

func (c *Command) Exec() ([]Failure, error) {
	_, err := exec.Command((*c)[0], (*c)[1:]...).Output()
	if err != nil {
		return nil, err
	}
	// TODO: parse
	return nil, nil
}
