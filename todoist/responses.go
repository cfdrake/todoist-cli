package todoist

import (
	"fmt"
)

// Represents a response from the Sync API.
type SyncResponse struct {
	Items    []*ItemResponse
	Projects []*ProjectResponse
}

// Returns the project with the given id, if any.
func (r *SyncResponse) ProjectWithId(id int) *ProjectResponse {
	for _, project := range r.Projects {
		if project.Id == id {
			return project
		}
	}

	return nil
}

// Returns the items associated with the given project.
func (r *SyncResponse) ItemsForProject(id int) []*ItemResponse {
	items := []*ItemResponse{}

	for _, item := range r.Items {
		if item.ProjectId == id {
			items = append(items, item)
		}
	}

	return items
}

// Denormalizes the dataset.
//
// Adds pointers to projects for each item and an array of pointers to
// items for each project.
func (r *SyncResponse) Denormalize() {
	// Associate items with their project.
	for _, item := range r.Items {
		item.Project = r.ProjectWithId(item.ProjectId)
	}

	// Associate projects with their items.
	for _, project := range r.Projects {
		project.Items = r.ItemsForProject(project.Id)
	}
}

// Represents an Item.
type ItemResponse struct {
	Id        int
	Content   string
	Indent    int
	ProjectId int `json:"project_id"`
	Project   *ProjectResponse
}

func (i ItemResponse) String() string {
	return i.Content
}

// Represents a Project.
type ProjectResponse struct {
	Id     int
	Name   string
	Indent int
	Items  []*ItemResponse
}

func (p ProjectResponse) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Id)
}
