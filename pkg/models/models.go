package models

import "time"

//Snippet struct holding the DB entry
type Snippet struct {
	ID         int
	Title      string
	Content    string
	CreateTime time.Time
	ExpTime    time.Time
}
