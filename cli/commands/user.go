package commands

import (
	"fmt"

	"github.com/cfdrake/todoist-cli/todoist"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

func UserCommands(client *todoist.Client) cli.Command {
	var info = func(c *cli.Context) error {
		res := new(todoist.ReadResult)
		err := client.MakeRequest(todoist.UserRequest, res)

		if err == nil {
			u := res.User
			fmt.Printf(chalk.Bold.TextStyle("%s\n\n"), u.FullName)
			printAttr(" Email", "%s", u.Email)
			printAttr(" Completed Today", "%d", u.CompletedToday)
			printAttr(" Completed Total", "%d", u.CompletedCount)
			printAttr(" Karma", "%.1f", u.Karma)
		} else {
			die("Could not load user information... (%s)", err)
		}

		return nil
	}

	return cli.Command{
		Name:        "user",
		Aliases:     []string{"u"},
		Usage:       "Displays user account information",
		Action:      info,
		Subcommands: []cli.Command{},
	}
}
