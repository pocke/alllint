package main

import (
	"testing"

	"github.com/ogier/pflag"
)

func TestImplementPFlagValue(t *testing.T) {
	s := make(Commands, 0)
	var _ pflag.Value = &s
}
