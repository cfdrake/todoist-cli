package cli

import (
	"fmt"
	"os/user"
	"strconv"

	"github.com/cfdrake/todoist-cli/todoist"
	"github.com/urfave/cli"
)

var (
	// Shared app instance.
	app *cli.App

	// Shared API client.
	client *todoist.Client
)

func init() {
	app = cli.NewApp()

	app.Usage = "Todoist.com command line client"
	app.Version = "0.1.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Colin Drake",
			Email: "todoist-support@colinfdrake.fastmail.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "item",
			Aliases: []string{"i", "items"},
			Usage:   "Commands for items",
			Action: func(c *cli.Context) error {
				fmt.Println("List all items")
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "List all items",
					Action: func(c *cli.Context) error {
						fmt.Println("List all items")
						return nil
					},
				},
				{
					Name:    "show",
					Aliases: []string{"s"},
					Usage:   "Show item details",
					Action: func(c *cli.Context) error {
						fmt.Println("Show item details")
						return nil
					},
				},
			},
		},
		{
			Name:    "project",
			Aliases: []string{"p", "projects"},
			Usage:   "Commands for projects",
			Action: func(c *cli.Context) error {
				displayAllProjects()
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "List all projects",
					Action: func(c *cli.Context) error {
						displayAllProjects()
						return nil
					},
				},
				{
					Name:    "show",
					Aliases: []string{"s"},
					Usage:   "Show project details",
					Action: func(c *cli.Context) error {
						idStr := c.Args().First()
						if id, err := strconv.Atoi(idStr); err != nil {
							die("Could not parse project ID...")
						} else {
							displayProject(id)
						}
						return nil
					},
				},
			},
		},
	}
}

// Runs the command line interface given input arguments.
func Run(args []string) {
	me, err := user.Current()
	if err != nil {
		die("Could not find current user!")
	}

	config, err := loadConfiguration(me)
	if err != nil {
		die(err)
	}

	client = &todoist.Client{UserToken: config.userToken}

	app.Run(args)
}
