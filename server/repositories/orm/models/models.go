package models

import (
	"gorm.io/gorm/schema"
)

// Model definition.
type Model interface {
	schema.Tabler
}

func ListModels() []Model {
	return []Model{
		&Account{},
		&VerifyAccount{},
		&Sessions{},
		&Product{},
	}
}
