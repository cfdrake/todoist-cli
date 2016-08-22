package main

import (
	"os"

	"github.com/cfdrake/todoist-cli/cli"
)

func main() {
	config := cli.NewFileConfig()
	c := cli.NewApp(config)
	c.Run(os.Args)
}
