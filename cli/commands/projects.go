package commands

import (
	"fmt"

	"github.com/cfdrake/todoist-cli/todoist"
	"github.com/urfave/cli"
)

func ProjectCommands(client *todoist.Client) cli.Command {
	var displayAll = func(c *cli.Context) error {
		fmt.Println("display all")
		return nil
	}

	var displayOne = func(c *cli.Context) error {
		fmt.Println("display one")
		return nil
	}

	var create = func(c *cli.Context) error {
		fmt.Println("create")
		return nil
	}

	return cli.Command{
		Name:    "project",
		Aliases: []string{"p", "projects"},
		Usage:   "Commands for projects",
		Action:  displayAll,
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all projects",
				Action:  displayAll,
			},
			{
				Name:    "show",
				Aliases: []string{"s"},
				Usage:   "Show project details",
				Action:  displayOne,
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create new project",
				Action:  create,
			},
		},
	}
}
