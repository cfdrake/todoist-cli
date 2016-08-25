package commands

import (
	"fmt"

	"github.com/cfdrake/todoist-cli/todoist"
	"github.com/urfave/cli"
)

func UserCommands(client *todoist.Client) cli.Command {
	var info = func(c *cli.Context) error {
		res := new(todoist.ReadResult)
		err := client.MakeRequest(todoist.UserRequest, res)

		if err == nil {
			u := res.User
			fmt.Printf("%s (%s)\n", u.FullName, u.Email)
			fmt.Printf("* Completed Today: %d\n", u.CompletedToday)
			fmt.Printf("* Total Completed: %d\n", u.CompletedCount)
			fmt.Printf("* Karma: %.1f\n", u.Karma)
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
