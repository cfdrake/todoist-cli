package cli

import (
	"github.com/cfdrake/todoist-cli/cli/commands"
	"github.com/cfdrake/todoist-cli/todoist"
	"github.com/urfave/cli"
)

// Command line application.
type Application struct {
	app *cli.App // Wrapping cli.App so as not to expose CLI utility library to callers.
}

// Creates a new application with the given configuration.
func NewApp(config Configurer) *Application {
	// Capture client to inject into commands.
	client := todoist.NewDefaultClient(config.UserToken())

	app := cli.NewApp()
	app.Usage = "Todoist.com command line client"
	app.Version = "0.1.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Colin Drake",
			Email: "todoist-support@colinfdrake.fastmail.com",
		},
	}
	app.Commands = []cli.Command{
		commands.ItemCommands(client),
		commands.ProjectCommands(client),
	}

	return &Application{app}
}

// Runs the application with the given arguments.
func (c *Application) Run(args []string) {
	c.app.Run(args)
}
