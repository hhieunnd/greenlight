package data

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Movies interface {
		Insert(movie *Movie) error
		Get(id int64) (*Movie, error)
		Update(movie *Movie) error
		Delete(id int64) error
	}
}

func NewModels(db *sqlx.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
