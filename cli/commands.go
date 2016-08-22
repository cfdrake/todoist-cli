package cli

import (
	"fmt"

	"github.com/ttacon/chalk"
)

func displayAllItems() {
	res, err := client.FetchItems()
	if err != nil {
		die("Could not fetch items...")
	}

	fmt.Println(chalk.Bold.TextStyle("All Items"))
	fmt.Println()

	for _, item := range res.Items {
		if item.ShouldDisplay() {
			fmt.Println("*", item)
		}
	}
}

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

	project := res.ProjectWithId(id)
	projectHeader := fmt.Sprintf("Project: %s", project)
	fmt.Println(chalk.Bold.TextStyle(projectHeader))

	for _, item := range project.Items {
		if item.ShouldDisplay() {
			fmt.Println("*", item)
		}
	}
}
