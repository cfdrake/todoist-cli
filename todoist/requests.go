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

// Represents a command in a request.
type command struct {
	Kind   string                 `json:"type"`
	Uuid   string                 `json:"uuid"`
	TempId *string                `json:"temp_id"`
	Args   map[string]interface{} `json:"args"`
}

// Returns the command represented as a JSON formatted string.
func (c command) JsonString() string {
	commands := []command{c}
	buf := new(bytes.Buffer)
	e := json.NewEncoder(buf)
	e.Encode(commands)
	jsonStr := buf.String()
	return jsonStr
}

// Request for all data.
var AllDataRequest = defaultParamsForAllResources([]string{"all"})

// Request for all projects and items.
var AllProjectsAndItemsRequest = defaultParamsForAllResources([]string{"projects", "items"})

// Request for all items.
var AllItemsRequest = defaultParamsForAllResources([]string{"items"})

// Request for user account info.
var UserRequest = defaultParamsForAllResources([]string{"user"})

// Request to complete an item.
func CompleteItemRequest(id int) RequestParams {
	operationUuid := uuid.NewV4().String()
	idStr := strconv.Itoa(id)
	cmd := command{
		Kind:   "item_complete",
		Uuid:   operationUuid,
		TempId: nil,
		Args: map[string]interface{}{
			"ids": []string{idStr},
		},
	}

	return RequestParams{"commands": {cmd.JsonString()}}
}

// Request to create an item.
func CreateItemRequest(content string, projectId int) RequestParams {
	operationUuid := uuid.NewV4().String()
	tempId := uuid.NewV4().String()
	cmd := command{
		Kind:   "item_add",
		Uuid:   operationUuid,
		TempId: &tempId,
		Args: map[string]interface{}{
			"content":    content,
			"project_id": projectId,
		},
	}

	return RequestParams{"commands": {cmd.JsonString()}}
}
