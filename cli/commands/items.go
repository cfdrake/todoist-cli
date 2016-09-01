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
		res := requireReadResult(client, todoist.AllItemsRequest)
		for _, item := range res.Items {
			displayItem(item)
		}
		return nil
	}

	var displayOne = func(c *cli.Context) error {
		id := parseInt(c.Args().Get(0))
		res := requireReadResult(client, todoist.AllItemsRequest)
		item := todoist.ItemWithId(res.Items, id)
		if item != nil {
			displayItem(item)
		} else {
			die("No such item")
		}
		return nil
	}

	var close = func(c *cli.Context) error {
		id := parseInt(c.Args().Get(0))
		requireWriteResult(client, todoist.CompleteItemRequest(id))
		fmt.Println("Closing item...")
		return nil
	}

	var create = func(c *cli.Context) error {
		args := c.Args()
		contents := args.Get(0)
		projectId := parseInt(args.Get(1))
		requireWriteResult(client, todoist.CreateItemRequest(contents, projectId))
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
				Name:    "close",
				Aliases: []string{"c"},
				Usage:   "Close unchecked item",
				Action:  close,
			},
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "Create new item",
				Action:  create,
			},
		},
	}
}
