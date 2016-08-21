package todoist

// Represents a response from the Sync API.
type Response struct {
  Items []Item
  Projects []Project
}

// Represents an Item.
type Item struct {
  ID int
  Content string
  Indent int
  ProjectID int
}

// Represents a Project.
type Project struct {
  ID int
  Name string
  Indent int
}
