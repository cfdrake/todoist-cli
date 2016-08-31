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

func closeItem(id int, client *todoist.Client) bool {
	items := fetchItems(client)
	item := todoist.ItemWithId(items, id)
	if item == nil {
		return false
	}

	res := new(todoist.WriteResult)
	err := client.MakeRequest(todoist.CompleteItemRequest(id), res)
	return err == nil
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
