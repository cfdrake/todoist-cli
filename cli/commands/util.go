package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cfdrake/todoist-cli/todoist"
	"github.com/ttacon/chalk"
)

func parseInt(intStr string) int {
	n, err := strconv.Atoi(intStr)
	if err != nil {
		die("Could not parse '%s' as an ID", intStr)
	}
	return n
}

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

func printAttr(name string, format string, a ...interface{}) {
	fmt.Printf(chalk.Dim.TextStyle("%s: "), name)
	fmt.Printf(format, a...)
	fmt.Println()
}

func die(format string, a ...interface{}) {
	formatStr := fmt.Sprintf("ERROR: %s\n", format)
	fmt.Fprintf(os.Stderr, formatStr, a...)
	os.Exit(1)
}

func requireWriteResult(c *todoist.Client, req todoist.RequestParams) *todoist.WriteResult {
	res := new(todoist.WriteResult)
	if err := c.MakeRequest(req, res); err != nil {
		die("Networking error (%s)...", err)
	}
	return res
}

func requireReadResult(c *todoist.Client, req todoist.RequestParams) *todoist.ReadResult {
	res := new(todoist.ReadResult)
	if err := c.MakeRequest(req, res); err != nil {
		die("Networking error (%s)...", err)
	}
	return res
}

func fetchProjectsAndItems(client *todoist.Client) ([]*todoist.Project, []*todoist.Item) {
	res := new(todoist.ReadResult)
	err := client.MakeRequest(todoist.AllProjectsAndItemsRequest, res)
	if err != nil {
		die("Could not fetch data... (%s)", err)
	}
	return res.Projects, res.Items
}
