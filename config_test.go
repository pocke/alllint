package main

import "testing"

func TestConfigFindLint(t *testing.T) {
	c := &Config{
		Lints: []Lint{
			{Name: "hoge"},
		},
	}

	if _, err := c.FindLint("hoge"); err != nil {
		t.Error(err)
	}

	if _, err := c.FindLint("fuga"); err == nil {
		t.Error("fuga should not be found, but found")
	}
}
