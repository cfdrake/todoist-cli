package cli

import (
	"os"
	"os/user"
	"strconv"
)

import "github.com/urfave/cli"
import "github.com/cfdrake/todoist-cli/todoist"

var app *cli.App
var client *todoist.Client

func init() {
	// Fetch the current user's config and initialize shared client.
	me, err := user.Current()
	if err != nil {
		die("ERROR: Could not find current user!")
		os.Exit(1)
	}

	config, err := loadConfiguration(me)
	if err != nil {
		die("ERROR:", err)
		os.Exit(1)
	}

	client = &todoist.Client{UserToken: config.userToken}

	// Setup app interface.
	app = cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:    "project",
			Aliases: []string{"p", "projects"},
			Usage:   "Commands for projects",
			Subcommands: []cli.Command{
				{
					Name:    "show",
					Aliases: []string{"s"},
					Usage:   "Show project details",
					Action: func(c *cli.Context) error {
						idStr := c.Args().First()
						if id, err := strconv.Atoi(idStr); err != nil {
							die("ERROR: Could not parse project ID...")
						} else {
							displayProject(id, client)
						}
						return nil
					},
				},
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "List all projects",
					Action: func(c *cli.Context) error {
						displayAllProjects(client)
						return nil
					},
				},
			},
		},
	}
}

func Run(args []string) {
	app.Run(args)
}
