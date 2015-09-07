package main

type Command []string

func (c *Command) Exec() ([]Failure, error) {
	return nil, nil
}
