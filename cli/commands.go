package cli

import (
	"fmt"
)

//import "github.com/cfdrake/todoist-cli/todoist"
import "github.com/ttacon/chalk"

func displayAllProjects() {
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

func displayProject(id int) {
	projects, items, err := client.FetchProjectsAndItems()
	if err != nil {
		die("ERROR: Could not fetch projects and/or items...")
	}

	for _, project := range *projects {
		if project.Id == id {
			projectHeader := fmt.Sprintf("Project: %s (%d)", project.Name, project.Id)
			fmt.Println(chalk.Bold.TextStyle(projectHeader))

			for _, item := range *items {
				if item.ProjectId == id {
					fmt.Println("*", item.Content)
				}
			}

			break
		}
	}
}
