package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// Wraps all data models, so it can be easily passed around within the application struct
type Models struct {
	Games GameModel
}

// initialises all Models into the wrapper
func NewModels(db *sql.DB) Models {
	return Models{
		Games: GameModel{DB: db},
	}
}
