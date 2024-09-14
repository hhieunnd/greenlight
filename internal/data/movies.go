package data

import "time"

type Movie struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"-"`
	Title    string    `json:"title"`
	Year     int32     `json:"year"`
	Runtime  Runtime   `json:"runtime,omitempty"`
	Genres   []string  `json:"genres"`  // (romance, comedy, etc..)
	Version  int32     `json:"version"` // Increate when update
}
