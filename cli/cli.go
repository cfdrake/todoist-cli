package cli

import (
	"fmt"
	"os"
	"os/user"
)

import "github.com/cfdrake/todoist-cli/todoist"

// Runs the command line interface with the given argument list.
func Execute(args []string) {
	// Fetch the current user's config.
	me, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Could not find current user!")
		os.Exit(1)
	}

	config, err := loadConfiguration(me)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}

	// Dummy test the networking code.
	c := todoist.Client{UserToken: config.userToken}
	resp, err := c.FetchAllData()

	fmt.Println(err)
	fmt.Println(resp)
}
