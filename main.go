package main

import "os"

func main() {
	c := &CLI{}
	err := c.Parse(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
