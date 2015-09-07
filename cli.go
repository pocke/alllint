package main

import "strings"

type StringSlice []string

func (v *StringSlice) String() string {
	return strings.Join(*v, ", ")
}

func (v *StringSlice) Set(s string) error {
	*v = append(*v, s)
	return nil
}
