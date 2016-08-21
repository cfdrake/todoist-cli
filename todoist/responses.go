package todoist

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
	ProjectID int
}

// Represents a Project.
type ProjectResponse struct {
	Id     int
	Name   string
	Indent int
}
