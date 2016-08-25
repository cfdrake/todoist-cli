package commands

import (
	"fmt"

	"github.com/cfdrake/todoist-cli/todoist"
	"github.com/urfave/cli"
)

func displayProject(project *todoist.Project, indented bool) {
	var whitespace string
	if indented {
		whitespace = indent(project.Indent)
	} else {
		whitespace = ""
	}
	checked := len(project.Items) == 0
	fmt.Printf("%s%s %s\n", whitespace, checkmark(checked), project)
}

func displayProjectItems(project *todoist.Project) {
	for _, item := range project.Items {
		displayItem(item)
	}
}

func ProjectCommands(client *todoist.Client) cli.Command {
	var displayAll = func(c *cli.Context) error {
		projects, _ := fetchProjectsAndItems(client)
		for _, project := range projects {
			displayProject(project, true)
		}
		return nil
	}

	var displayOne = func(c *cli.Context) error {
		id := parseInt(c.Args().Get(0))
		projects, _ := fetchProjectsAndItems(client)
		project := todoist.ProjectWithId(projects, id)
		if project != nil {
			displayProject(project, false)
			displayProjectItems(project)
		} else {
			die("No such project")
		}
		return nil
	}

	var create = func(c *cli.Context) error {
		fmt.Println("create")
		return nil
	}

	return cli.Command{
		Name:    "project",
		Aliases: []string{"p", "projects"},
		Usage:   "Commands for projects",
		Action:  displayAll,
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all projects",
				Action:  displayAll,
			},
			{
				Name:    "show",
				Aliases: []string{"s"},
				Usage:   "Show project details",
				Action:  displayOne,
			},
			{
				Name:    "new",
				Aliases: []string{"n"},
				Usage:   "Create new project",
				Action:  create,
			},
		},
	}
}
