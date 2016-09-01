package todoist

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

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
// Response objects must know how to parse themselves from JSON and validate.
// Will short-circuit and return an error if one occurred during the request.
func (c Client) MakeRequest(req RequestParams, res Responser) (err error) {
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

	// Perform JSON validation.
	if err = res.ValidateResponse(); err != nil {
		return
	}

	return
}
