package masters

import "github.com/google/uuid"

type Master struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Phone     string
	Position  string
}

type MastersReader interface {
	GetMasters() ([]Master, error)
}
