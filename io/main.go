package main

import (
	"os"
)

func usage() {
	println("Usage: cmd input/output <contest> <problem>")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}
	if os.Args[1] == "input" {
		Examples(os.Stdout, os.Args[2], os.Args[3], ".input")
	} else if os.Args[2] == "output" {
		Examples(os.Stdout, os.Args[2], os.Args[3], ".output")
	} else {
		usage()
	}
}
