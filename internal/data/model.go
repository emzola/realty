package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

// Models is a 'container' struct to wrap all models of the application.
type Models struct {
	Properties PropertyModel
}

// NewModels returns a models struct containing an uninitialised property model.
func NewModels(db *sql.DB) Models {
	return Models{
		Properties: PropertyModel{DB: db},
	}
}