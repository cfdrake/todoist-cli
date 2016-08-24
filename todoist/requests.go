package todoist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/satori/go.uuid"
)

// Represents a Sync API request type.
type RequestParams url.Values

// Builds a comma separated list string, for example: ["comma", "separated"].
func buildListStr(types []string) string {
	output := "["

	for i, t := range types {
		if i > 0 {
			output += ","
		}

		elem := fmt.Sprintf("\"%s\"", t)
		output += elem
	}

	return output + "]"
}

// Creates a new parameter set to query all of the items in the given types array.
func defaultParamsForAllResources(types []string) RequestParams {
	allObjectsSyncToken := "*"
	return RequestParams{
		"sync_token":     {allObjectsSyncToken},
		"resource_types": {buildListStr(types)},
	}
}

// Request for all data.
var AllDataRequest = defaultParamsForAllResources([]string{"all"})

// Request for all projects and items.
var AllProjectsAndItemsRequest = defaultParamsForAllResources([]string{"projects", "items"})

// Request for all items.
var AllItemsRequest = defaultParamsForAllResources([]string{"items"})

// Request for user account info.
var UserRequest = defaultParamsForAllResources([]string{"user"})

// Represents a command request.
type command struct {
	Kind string                 `json:"type"`
	Uuid string                 `json:"uuid"`
	Args map[string]interface{} `json:"args"`
}

// Request to complete an item.
func CompleteItemRequest(id int) RequestParams {
	uuid := uuid.NewV4().String()
	idStr := strconv.Itoa(id)
	cmd := []command{
		{
			Kind: "item_complete",
			Uuid: uuid,
			Args: map[string]interface{}{
				"ids": []string{idStr},
			},
		},
	}

	buf := new(bytes.Buffer)
	e := json.NewEncoder(buf)
	e.Encode(cmd)
	jsonStr := buf.String()
	fmt.Println(jsonStr)

	return RequestParams{"commands": {jsonStr}}
}
