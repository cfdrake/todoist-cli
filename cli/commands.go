package cli

import (
	"fmt"
)

import "github.com/ttacon/chalk"

func displayAllProjects() {
	projects, err := client.FetchProjects()
	if err != nil {
		die("ERROR: Could not fetch projects...")
	}

	fmt.Println(chalk.Bold.TextStyle("All Projects"))
	fmt.Println()

	for _, project := range *projects {
		fmt.Println("*", project)
	}
}

func displayProject(id int) {
	projects, items, err := client.FetchProjectsAndItems()
	if err != nil {
		die("ERROR: Could not fetch projects and/or items...")
	}

	for _, project := range *projects {
		if project.Id == id {
			projectHeader := fmt.Sprintf("Project: %s", project)
			fmt.Println(chalk.Bold.TextStyle(projectHeader))

			for _, item := range *items {
				if item.ProjectId == id {
					fmt.Println("*", item)
				}
			}

			break
		}
	}
}
