package todoist

import (
	"encoding/json"
	"fmt"
)

// Represents a read request result.
type ReadResult struct {
	Items    []*Item
	Projects []*Project
}

// Unpacks a read result from a JSON representation.
// Ensures this type adheres to JsonUnmarshaler.
func (r *ReadResult) UnmarshalJson(bytes []byte) error {
	if err := json.Unmarshal(bytes, r); err != nil {
		return err
	}

	r.denormalize()
	return nil
}

// Returns the item with the given id, if any.
func ItemWithId(items []*Item, id int) *Item {
	for _, item := range items {
		if item.Id == id {
			return item
		}
	}

	return nil
}

// Returns the project with the given id, if any.
func ProjectWithId(projects []*Project, id int) *Project {
	for _, project := range projects {
		if project.Id == id {
			return project
		}
	}

	return nil
}

// Returns the items associated with the given project.
// This is meant for internal use: clients calling this code should use the denormalized
// representation to traverse project structure.
func (r *ReadResult) itemsForProject(id int) []*Item {
	items := []*Item{}

	for _, item := range r.Items {
		if item.ProjectId == id {
			items = append(items, item)
		}
	}

	return items
}

// Denormalizes the dataset.
// Adds pointers to projects for each item and an array of pointers to items for each project.
func (r *ReadResult) denormalize() {
	// Associate items with their project.
	for _, item := range r.Items {
		item.Project = ProjectWithId(r.Projects, item.ProjectId)
	}

	// Associate projects with their items.
	for _, project := range r.Projects {
		project.Items = r.itemsForProject(project.Id)
	}
}

// Represents an Item.
type Item struct {
	Id        int
	Content   string
	Indent    int
	Archived  int `json:"is_archived"`
	Deleted   int `json:"is_deleted"`
	ProjectId int `json:"project_id"`
	Project   *Project
}

// Adhere to Stringer interface.
func (i Item) String() string {
	return fmt.Sprintf("%s (%d)", i.Content, i.Id)
}

// Represents a Project.
type Project struct {
	Id       int
	Name     string
	Indent   int
	Items    []*Item
	Archived int `json:"is_archived"`
	Deleted  int `json:"is_deleted"`
}

// Adhere to Stringer interface.
func (p Project) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Id)
}
