package main

type Failure struct {
	FName    string
	Line     int
	Col      int
	Message  string
	LintName string
}
