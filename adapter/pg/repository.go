package pg

import (
	"fmt"
)

type Repository struct {
	*DB
}

func NewRepository(db *DB) (*Repository, error) {
	if db == nil {
		return nil, fmt.Errorf("db is nil")
	}
	return &Repository{db}, nil
}
