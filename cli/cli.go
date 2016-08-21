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
		fmt.Fprintln(os.Stderr, "ERROR: Could not find current user!\n")
		os.Exit(1)
	}

	config, err := loadConfiguration(me)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err, "\n")
		os.Exit(1)
	}

	// Setup client access.
	client := todoist.Client{UserToken: config.userToken}

	// Determine command to execute.
	if len(args) == 0 {
		fmt.Println("Usage...")
		return
	}

	var commandArgs []string
	command := args[0]

	if len(args) > 1 {
		commandArgs = args[1:]
	}

	switch command {
	case "i", "items", "item":
		break
	case "p", "projects", "project":
		handleProjectsCommand(client, commandArgs)
	case "s", "summary":
		break
	}
}
