package main

import (
	"reflect"
	"testing"

	"github.com/ogier/pflag"
)

func TestCommandsImplementPFlagValue(t *testing.T) {
	s := make(Commands, 0)
	var _ pflag.Value = &s
}

func TestParse(t *testing.T) {
	opts := []string{"alllint", "--command=tslint", "-c", "tsc --noImplicitAny"}
	c := &CLI{}
	err := c.Parse(opts)
	if err != nil {
		t.Error(err)
	}

	expected := Commands{
		{"tslint"},
		{"tsc", "--noImplicitAny"},
	}
	if !reflect.DeepEqual(c.Commands, expected) {
		t.Errorf("\n%+v !=\n%+v", c.Commands, expected)
	}
}
