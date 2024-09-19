package data

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Movie struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"-"`
	Title    string    `json:"title"`
	Year     int32     `json:"year"`
	Runtime  Runtime   `json:"runtime,omitempty"`
	Genres   []string  `json:"genres"`  // (romance, comedy, etc..)
	Version  int32     `json:"version"` // Increate when update
}

type MovieModel struct {
	DB *sqlx.DB
}

func (m MovieModel) Insert(movie *Movie) error {
	return nil
}

func (m MovieModel) Get(id int64) (*Movie, error) {
	return nil, nil
}

func (m MovieModel) Update(movie *Movie) error {
	return nil
}

func (m MovieModel) Delete(id int64) error {
	return nil
}
