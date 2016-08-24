package todoist

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/satori/go.uuid"
)

// Represents a Sync API request type.
type RequestParams url.Values

// Describes a type that may be unmarshaled via JSON response data.
type ResponseUnmarshaler interface {
	UnmarshalJson(bytes []byte) error
}

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

func CompleteItemRequest(id int) RequestParams {
	uuid := uuid.NewV4().String()
	idString := strconv.Itoa(id)
	jsonEncoded := "[{\"type\": \"item_complete\", \"uuid\": \"" + uuid + "\", \"args\": {\"ids\": [\"" + idString + "\"]}}]"
	return RequestParams{
		"commands": {jsonEncoded},
	}
}

// Todoist Sync API client.
// Carries a token corresponding to a user's session and acts on a base API URL.
type Client struct {
	Token      string
	SyncApiUrl string
}

// Default production endpoint URL for the Todoist Sync API.
const defaultSyncUrl string = "https://todoist.com/API/v7/sync"

// Creates a client pointed to the default Todoist Sync API URL.
func NewDefaultClient(token string) *Client {
	return &Client{token, defaultSyncUrl}
}

// The base parameter values used by the API client.
func (c Client) baseParams() RequestParams {
	return RequestParams{
		"token": {c.Token},
	}
}

// Performs the given request, providing the response in the passed in object.
// Response objects must know how to parse themselves from JSON.
// Will short-circuit and return an error if one occurred during the request.
func (c Client) MakeRequest(req RequestParams, res ResponseUnmarshaler) (err error) {
	// Build parameter set.
	params := c.baseParams()

	for k, v := range req {
		params[k] = v
	}

	// Make HTTP request.
	resp, err := http.PostForm(c.SyncApiUrl, url.Values(params))
	if err != nil {
		return
	}

	// Parse JSON.
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if err = res.UnmarshalJson(body); err != nil {
		return
	}

	return
}
