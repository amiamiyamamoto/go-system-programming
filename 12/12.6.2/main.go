package main

import "io"

import (
	colorable"github.com/mattn/go-colorable"
	isatty "github.com/mattn/go-isatty"
)
func main() {
	var out io.Writer
	if isatty.IsTerminal(os.Stdout.Fd())
}
