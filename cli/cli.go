package cli

import (
	"fmt"
	"strconv"

	"github.com/cfdrake/todoist-cli/todoist"
	"github.com/urfave/cli"
)

// Parses an integer ID from a string, or fails.
func parseId(idString string) int {
	id, err := strconv.Atoi(idString)
	if err != nil {
		die("Could not parse '" + idString + "' as an ID")
	}
	return id
}

// Command line application.
type CliApplication struct {
	app    *cli.App
	client *todoist.Client
}

// Creates a new application with the given configuration.
func NewApp(config Configurer) *CliApplication {
	// Setup client and app instances.
	client := &todoist.Client{UserToken: config.UserToken()}
	app := cli.NewApp()

	app.Usage = "Todoist.com command line client"
	app.Version = "0.1.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Colin Drake",
			Email: "todoist-support@colinfdrake.fastmail.com",
		},
	}

	// Configure possible subcommands.
	app.Commands = []cli.Command{
		{
			Name:    "item",
			Aliases: []string{"i", "items"},
			Usage:   "Commands for items",
			Action: func(c *cli.Context) error {
				displayAllItems(client)
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "List all items",
					Action: func(c *cli.Context) error {
						displayAllItems(client)
						return nil
					},
				},
				{
					Name:    "show",
					Aliases: []string{"s"},
					Usage:   "Show item details",
					Action: func(c *cli.Context) error {
						id := parseId(c.Args().First())
						displayItem(id, client)
						return nil
					},
				},
				{
					Name:    "toggle",
					Aliases: []string{"t"},
					Usage:   "Toggle item checked status",
					Action: func(c *cli.Context) error {
						fmt.Println("Toggle item checked status")
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
				displayAllProjects(client)
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "List all projects",
					Action: func(c *cli.Context) error {
						displayAllProjects(client)
						return nil
					},
				},
				{
					Name:    "show",
					Aliases: []string{"s"},
					Usage:   "Show project details",
					Action: func(c *cli.Context) error {
						id := parseId(c.Args().First())
						displayProject(id, client)
						return nil
					},
				},
			},
		},
	}

	return &CliApplication{app: app, client: client}
}

// Runs the application with the given arguments.
func (c *CliApplication) Run(args []string) {
	c.app.Run(args)
}
