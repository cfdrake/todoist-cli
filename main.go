package main

import (
	"os"
)

import "github.com/cfdrake/todoist-cli/cli"

func main() {
	cli.Run(os.Args)
}
