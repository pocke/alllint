package main

import (
	"bufio"
	"strconv"
)
import "bytes"
import "os/exec"

type Command []string

func (c *Command) Exec(files []string) ([]Failure, error) {
	b, err := exec.Command(c.Name(), c.Args(files)...).Output()
	if _, execError := err.(*exec.Error); err != nil && execError {
		return nil, err
	}
	return c.parse(b)
}

func (c *Command) parse(b []byte) ([]Failure, error) {
	lint, err := config.FindLint(c.Name())
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(b)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)

	res := make([]Failure, 0)

	for sc.Scan() {
		matches := lint.Re().FindStringSubmatch(sc.Text())
		if len(matches) == 0 {
			continue
		}

		// TODO: refactor
		line := -1
		if i, err := strconv.Atoi(matches[lint.Line]); err == nil {
			line = i
		}

		col := -1
		if i, err := strconv.Atoi(matches[lint.Col]); err == nil {
			col = i
		}

		f := Failure{
			FName:    matches[lint.FName],
			Line:     line,
			Col:      col,
			Message:  matches[lint.Message],
			LintName: lint.Name,
		}
		res = append(res, f)
	}
	return res, nil
}

func (c *Command) Args(files []string) []string {
	args := (*c)[1:]
	return append(args, files...)
}

func (c *Command) Name() string {
	return (*c)[0]
}
