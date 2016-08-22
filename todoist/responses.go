package todoist

import (
	"fmt"
)

// Represents a read response from the Sync API.
type ReadResponse struct {
	Items    []*ItemResponse
	Projects []*ProjectResponse
}

// Returns the project with the given id, if any.
func (r *ReadResponse) ProjectWithId(id int) *ProjectResponse {
	for _, project := range r.Projects {
		if project.Id == id {
			return project
		}
	}

	return nil
}

// Returns the items associated with the given project.
func (r *ReadResponse) ItemsForProject(id int) []*ItemResponse {
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
func (r *ReadResponse) Denormalize() {
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
	Archived  int `json:"is_archived"`
	Deleted   int `json:"is_deleted"`
	ProjectId int `json:"project_id"`
	Project   *ProjectResponse
}

func (i ItemResponse) String() string {
	return i.Content
}

func (i ItemResponse) ShouldDisplay() bool {
	return i.Archived == 0 && i.Deleted == 0
}

// Represents a Project.
type ProjectResponse struct {
	Id       int
	Name     string
	Indent   int
	Items    []*ItemResponse
	Archived int `json:"is_archived"`
	Deleted  int `json:"is_deleted"`
}

func (p ProjectResponse) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Id)
}

func (p ProjectResponse) ShouldDisplay() bool {
	return p.Archived == 0 && p.Deleted == 0
}
