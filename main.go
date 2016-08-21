package main

import (
	"flag"
)

import "github.com/cfdrake/todoist-cli/cli"

func main() {
	flag.Parse()
	args := flag.Args()
	cli.Execute(args)
}
