package main

import "strings"

type Signal struct {
	defs   []string // definitions of 10 unique signal patterns
	output []string // actual output
}

func NewSignal(s string) Signal {
	parts := strings.Split(s, "|")
	defs, output := strings.Split(strings.TrimSpace(parts[0]), " "), strings.Split(strings.TrimSpace(parts[1]), " ")
	return Signal{
		defs:   defs,
		output: output,
	}
}