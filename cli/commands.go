package cli

import (
	"fmt"

	"github.com/ttacon/chalk"
)

func displayAllProjects() {
	res, err := client.FetchProjects()
	if err != nil {
		die("Could not fetch projects...")
	}

	fmt.Println(chalk.Bold.TextStyle("All Projects"))
	fmt.Println()

	for _, project := range res.Projects {
		if project.ShouldDisplay() {
			fmt.Println("*", project)
		}
	}
}

func displayProject(id int) {
	res, err := client.FetchProjectsAndItems()
	if err != nil {
		die("Could not fetch projects and/or items...")
	}

	for _, project := range res.Projects {
		if project.Id == id {
			projectHeader := fmt.Sprintf("Project: %s", project)
			fmt.Println(chalk.Bold.TextStyle(projectHeader))

			for _, item := range project.Items {
				if item.ShouldDisplay() {
					fmt.Println("*", item)
				}
			}

			break
		}
	}
}
