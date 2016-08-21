package cli

import (
	"fmt"
	"os"
	"strconv"
)

import "github.com/cfdrake/todoist-cli/todoist"
import "github.com/ttacon/chalk"

func handleProjectsCommand(client todoist.Client, args []string) {
	if len(args) == 0 {
		displayAllProjects(client)
	} else if i, err := strconv.Atoi(args[0]); err == nil {
		fmt.Printf("Project ID: %d\n", i)
	} else {
		fmt.Fprintf(os.Stderr, "ERROR: No such subcommand...\n")
		os.Exit(1)
	}
}

func displayAllProjects(client todoist.Client) {
	projects, err := client.FetchProjects()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Could not fetch projects...\n")
		os.Exit(1)
	}

	fmt.Println(chalk.Bold.TextStyle("All Projects"))
	fmt.Println()

	for _, project := range *projects {
		id := fmt.Sprintf("(%d)", project.Id)
		fmt.Println(
			"*",
			chalk.Underline.TextStyle(project.Name),
			chalk.White.Color(id),
			chalk.Reset,
		)
	}
}

func displayProject() {
}
