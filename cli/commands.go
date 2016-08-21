package cli

import (
	"fmt"
	"strconv"
)

import "github.com/cfdrake/todoist-cli/todoist"
import "github.com/ttacon/chalk"

func handleProjectsCommand(client todoist.Client, args []string) {
	if len(args) == 0 {
		displayAllProjects(client)
	} else if id, err := strconv.Atoi(args[0]); err == nil {
		displayProject(id, client)
	} else {
		die("ERROR: No such subcommand...")
	}
}

func displayAllProjects(client todoist.Client) {
	projects, err := client.FetchProjects()
	if err != nil {
		die("ERROR: Could not fetch projects...")
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

func displayProject(id int, client todoist.Client) {
	resp, err := client.FetchAllData()
	if err != nil {
		die("ERROR: Could not fetch projects and/or items...")
	}

	for _, project := range resp.Projects {
		if project.Id == id {
			projectHeader := fmt.Sprintf("Project: %s (%d)", project.Name, project.Id)
			fmt.Println(chalk.Bold.TextStyle(projectHeader))

			for _, item := range resp.Items {
				if item.ProjectId == id {
					fmt.Println("*", item.Content)
				}
			}

			break
		}
	}
}
