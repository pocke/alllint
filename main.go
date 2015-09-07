package main

import (
	"fmt"
	"os"
)

func main() {
	c := &CLI{}
	err := c.Parse(os.Args)
	if err != nil {
		os.Exit(1)
	}
	f, err := c.Exec()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(f)
}
