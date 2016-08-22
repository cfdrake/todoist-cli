package todoist

import (
	"fmt"
)

// A write command representation.
type WriteCommand struct {
	Type   string
	Args   map[string]string
	Uuid   string
	TempId string `json:"temp_id"`
}

// A request to perform an array of WriteCommands.
type WriteRequest struct {
	Commands []WriteCommand
}

// A write result.
type WriteResult struct {
	// TODO: ...
}

// Represents a read response from the Sync API.
type ReadResult struct {
	Items    []*Item
	Projects []*Project
}

// Returns the item with the given id, if any.
func (r *ReadResult) ItemWithId(id int) *Item {
	for _, item := range r.Items {
		if item.Id == id {
			return item
		}
	}

	return nil
}

// Returns the project with the given id, if any.
func (r *ReadResult) ProjectWithId(id int) *Project {
	for _, project := range r.Projects {
		if project.Id == id {
			return project
		}
	}

	return nil
}

// Returns the items associated with the given project.
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
		item.Project = r.ProjectWithId(item.ProjectId)
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

// Predicate indicating whether or not the item should be shown.
func (i Item) ShouldDisplay() bool {
	return i.Archived == 0 && i.Deleted == 0
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

// Predicate indicating whether or not the project should be shown.
func (p Project) ShouldDisplay() bool {
	return p.Archived == 0 && p.Deleted == 0
}
