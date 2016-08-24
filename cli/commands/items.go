package commands

import (
	"fmt"

	"github.com/cfdrake/todoist-cli/todoist"
	"github.com/urfave/cli"
)

func ItemCommands(client *todoist.Client) cli.Command {
	var displayAll = func(c *cli.Context) error {
		items := fetchItems(client)
		for _, i := range items {
			fmt.Printf("%s%s %s\n", indent(i.Indent), checkmark(false), i)
		}
		return nil
	}

	var displayOne = func(c *cli.Context) error {
		fmt.Println("display one")
		return nil
	}

	var toggle = func(c *cli.Context) error {
		fmt.Println("toggle")
		return nil
	}

	var create = func(c *cli.Context) error {
		fmt.Println("create")
		return nil
	}

	return cli.Command{
		Name:    "item",
		Aliases: []string{"i", "items"},
		Usage:   "Commands for items",
		Action:  displayAll,
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all items",
				Action:  displayAll,
			},
			{
				Name:    "show",
				Aliases: []string{"s"},
				Usage:   "Show item details",
				Action:  displayOne,
			},
			{
				Name:    "toggle",
				Aliases: []string{"t"},
				Usage:   "Toggle item checked status",
				Action:  toggle,
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "Create new item",
				Action:  create,
			},
		},
	}
}
