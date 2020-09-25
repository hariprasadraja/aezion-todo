package model

import "time"

// ListItem represents a single task entry or an item in the Todo List
type ListItem struct {
	Index        int       `json:"index,omitempty"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	CreatedTime  time.Time `json:"created_time,omitempty"`
	ModifiedTime time.Time `json:"modified_time,omitempty"`
}

// TodoList represents a single List of Todo Items
type TodoList struct {
	ID    int32      `json:"id,omitempty"`
	Name  string     `json:"name,omitempty"`
	Items []ListItem `json:"tasks,omitempty"`
}
