package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/cfdrake/todoist-cli/todoist"
)

func checkmark(checked bool) string {
	if checked {
		return "✔"
	} else {
		return "✘"
	}
}

func indent(indentLevel int) string {
	return strings.Repeat(" ", indentLevel-1)
}

func die(format string, a ...interface{}) {
	formatStr := fmt.Sprintf("ERROR: %s\n", format)
	fmt.Fprintf(os.Stderr, formatStr, a...)
	os.Exit(1)
}

func fetchItems(client *todoist.Client) []*todoist.Item {
	res := new(todoist.ReadResult)
	err := client.MakeRequest(todoist.AllItemsRequest, res)
	if err != nil {
		die("Could not fetch items... (%s)", err)
	}
	return res.Items
}

func fetchProjectsAndItems(client *todoist.Client) ([]*todoist.Project, []*todoist.Item) {
	res := new(todoist.ReadResult)
	err := client.MakeRequest(todoist.AllProjectsAndItemsRequest, res)
	if err != nil {
		die("Could not fetch data... (%s)", err)
	}
	return res.Projects, res.Items
}
