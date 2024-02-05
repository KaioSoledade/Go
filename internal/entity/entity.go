package entity

import "github.com/google/uuid"

type Categoty struct {
	ID string
	Name string
}

func NewCategory(name string) *Category {
	return &Category{
		ID: uuid.New().String(),
		Name: name,
	}
}