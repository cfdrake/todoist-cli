package main

import (
	"flag"
)

import "github.com/cfdrake/todoist-cli/cli"

func main() {
	args := flag.Args()
	cli.Execute(args)
}
