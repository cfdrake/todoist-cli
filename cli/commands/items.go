package commands

import (
	"fmt"

	"github.com/cfdrake/todoist-cli/todoist"
	"github.com/urfave/cli"
)

func displayItem(item *todoist.Item) {
	fmt.Printf("%s%s %s\n", indent(item.Indent), checkmark(false), item)
}

func ItemCommands(client *todoist.Client) cli.Command {
	var displayAll = func(c *cli.Context) error {
		items := fetchItems(client)
		for _, item := range items {
			displayItem(item)
		}
		return nil
	}

	var displayOne = func(c *cli.Context) error {
		id := parseInt(c.Args().Get(0))
		items := fetchItems(client)
		item := todoist.ItemWithId(items, id)
		if item != nil {
			displayItem(item)
		} else {
			die("No such item")
		}
		return nil
	}

	var toggle = func(c *cli.Context) error {
		id := parseInt(c.Args().Get(0))
		if ok := toggleItem(id, client); ok {
			fmt.Println("Toggled item...")
		} else {
			fmt.Println("Could not toggle item...")
		}
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
