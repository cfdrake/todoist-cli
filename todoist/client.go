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

	for _, t := range resourceTypes {
		elem := fmt.Sprintf("\"%s\"", t)
		output += elem
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

// Calls the Sync service with the given inputs and returns the body decoded into a SyncResponse type.
func (c *Client) callSyncService(syncToken string, resourceTypes []ResourceTyper) (*SyncResponse, error) {
	params := generateBaseParams(c.UserToken, syncToken, resourceTypes)
	resp, err := http.PostForm(endpointUrl, params)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	model := &SyncResponse{}
	if err = json.Unmarshal(responseBytes, model); err != nil {
		return nil, err
	}

	return model, nil
}

func (c *Client) FetchAllData() (*SyncResponse, error) {
	types := []ResourceTyper{allDataResourceType}
	return c.callSyncService(initialSyncToken, types)
}

func (c *Client) FetchProjectsAndItems() (*[]ProjectResponse, *[]ItemResponse, error) {
	types := []ResourceTyper{projectsResourceType, itemsResourceType}
	resp, err := c.callSyncService(initialSyncToken, types)

	if err != nil {
		return nil, nil, err
	}

	return &resp.Projects, &resp.Items, nil
}

func (c *Client) FetchProjects() (*[]ProjectResponse, error) {
	types := []ResourceTyper{projectsResourceType}
	resp, err := c.callSyncService(initialSyncToken, types)

	if err != nil {
		return nil, err
	}

	return &resp.Projects, nil
}
