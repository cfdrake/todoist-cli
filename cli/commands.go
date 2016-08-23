package cli

import (
	"fmt"
	"strings"

	"github.com/cfdrake/todoist-cli/todoist"
	"github.com/ttacon/chalk"
)

func indent(indentLevel int) {
	whitespace := strings.Repeat(" ", indentLevel-1)
	fmt.Print(whitespace)
}

func fetchProjectsAndItemsOrFail(client *todoist.Client) *todoist.ReadResult {
	res, err := client.FetchProjectsAndItems()
	if err != nil {
		die("Could not fetch the requested data...")
	}
	return res
}

func displayAllItems(client *todoist.Client) {
	res, err := client.FetchItems()
	if err != nil {
		die("Could not fetch items...")
	}

	fmt.Println(chalk.Bold.TextStyle("All Items"))

	for _, item := range res.Items {
		if item.ShouldDisplay() {
			indent(item.Indent)
			fmt.Println("✘", item)
		}
	}
}

func displayItem(id int, client *todoist.Client) {
	res := fetchProjectsAndItemsOrFail(client)
	item := res.ItemWithId(id)
	itemHeader := fmt.Sprintf("Item: %s", item)
	fmt.Println(chalk.Bold.TextStyle(itemHeader))
	fmt.Print("* Project:")
	fmt.Println(item.Project)
}

func displayAllProjects(client *todoist.Client) {
	res := fetchProjectsAndItemsOrFail(client)
	fmt.Println(chalk.Bold.TextStyle("All Projects"))

	for _, project := range res.Projects {
		if project.ShouldDisplay() {
			indent(project.Indent)
			if len(project.Items) > 0 {
				fmt.Println("✘", project)
			} else {
				fmt.Println("✔", project)
			}
		}
	}
}

func displayProject(id int, client *todoist.Client) {
	res := fetchProjectsAndItemsOrFail(client)
	project := res.ProjectWithId(id)
	projectHeader := fmt.Sprintf("Project: %s", project)
	fmt.Println(chalk.Bold.TextStyle(projectHeader))

	for _, item := range project.Items {
		if item.ShouldDisplay() {
			fmt.Println("✘", item)
		}
	}
}
