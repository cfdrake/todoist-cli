package todoist

import (
	"fmt"
)

// Represents a response from the Sync API.
type SyncResponse struct {
	Items    []ItemResponse
	Projects []ProjectResponse
}

// Represents an Item.
type ItemResponse struct {
	Id        int
	Content   string
	Indent    int
	ProjectId int `json:"project_id"`
}

func (i ItemResponse) String() string {
	return i.Content
}

// Represents a Project.
type ProjectResponse struct {
	Id     int
	Name   string
	Indent int
}

func (p ProjectResponse) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Id)
}
