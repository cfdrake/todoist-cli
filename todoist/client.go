package todoist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Endpoint URL for the Todoist Sync API.
const endpointUrl string = "https://todoist.com/API/v7/sync"

// Resource type indicating what data to pull down.
type resourceType string

const (
	itemsResourceType    resourceType = "items"
	projectsResourceType resourceType = "projects"
	allDataResourceType  resourceType = "all"
)

type ResourceTyper interface {
	ResourceType() resourceType
}

func (t resourceType) ResourceType() resourceType {
	return t
}

// Client data type.
type Client struct {
	UserToken string
}

// Default sync token to fetch all data.
const initialSyncToken string = "*"

// Builds the resource types string expected by the Sync endpoint.
func buildResourceTypesString(resourceTypes []ResourceTyper) string {
	output := "["
	first := true

	for _, t := range resourceTypes {
		if !first {
			output += ","
		}

		elem := fmt.Sprintf("\"%s\"", t)
		output += elem
		first = false
	}

	output += "]"
	return output
}

// Generates the basic required URL parameters to use for an API request.
func generateBaseParams(userToken string, syncToken string, resourceTypes []ResourceTyper) url.Values {
	typeString := buildResourceTypesString(resourceTypes)
	return url.Values{
		"token":          {userToken},
		"sync_token":     {syncToken},
		"resource_types": {typeString},
	}
}

// Performs a read request and returns the result from the Sync service.
func (c *Client) performReadRequest(syncToken string, resourceTypes []ResourceTyper) (res *ReadResult, err error) {
	params := generateBaseParams(c.UserToken, syncToken, resourceTypes)
	resp, err := http.PostForm(endpointUrl, params)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	res = &ReadResult{}
	if err = json.Unmarshal(responseBytes, res); err != nil {
		return
	}
	res.denormalize()

	return
}

// Fetches all data for the given user.
func (c *Client) FetchAllData() (*ReadResult, error) {
	types := []ResourceTyper{allDataResourceType}
	return c.performReadRequest(initialSyncToken, types)
}

// Fetches project and item data for the user.
func (c *Client) FetchProjectsAndItems() (*ReadResult, error) {
	types := []ResourceTyper{projectsResourceType, itemsResourceType}
	return c.performReadRequest(initialSyncToken, types)
}

// Fetches item data for the user.
func (c *Client) FetchItems() (*ReadResult, error) {
	types := []ResourceTyper{itemsResourceType}
	return c.performReadRequest(initialSyncToken, types)
}

// Fetches project data for the user.
func (c *Client) FetchProjects() (*ReadResult, error) {
	types := []ResourceTyper{projectsResourceType}
	return c.performReadRequest(initialSyncToken, types)
}
