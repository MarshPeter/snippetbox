package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
    ID int
    Title string
    Content string
    Created time.Time
    Expires time.Time
}

type SnippetModel struct {
    DB *sql.DB
}

// insert new snippet into the database.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
    return 0, nil
}

// returns a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (Snippet, error) {
    return Snippet{}, nil
}

// returns the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
    return nil, nil
}


