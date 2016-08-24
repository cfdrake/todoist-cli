package main

import (
	"os"

	"github.com/cfdrake/todoist-cli/cli"
)

func main() {
	c := cli.NewApp(Config())
	c.Run(os.Args)
}
