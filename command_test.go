package main

import (
	"reflect"
	"testing"
)

func TestCommandName(t *testing.T) {
	c := Command{"tsc", "--noImplicitAny"}
	if c.Name() != "tsc" {
		t.Errorf("Name should be tsc, but got %s", c.Name())
	}
}

func TestCommandArgs(t *testing.T) {
	c := Command{"tsc", "--noImplicitAny", "--noEmitOnError"}
	f := []string{"hoge.ts", "fuga.ts"}

	e := []string{"--noImplicitAny", "--noEmitOnError", "hoge.ts", "fuga.ts"}
	if !reflect.DeepEqual(c.Args(f), e) {
		t.Errorf("Expected: %+v, but got %+v", e, c.Args(f))
	}
}
