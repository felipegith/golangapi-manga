package entity

import "github.com/google/uuid"

type Manga struct {
	ID          string
	Name        string
	Description string
}

func Constructor(name, description string) *Manga {
	return &Manga{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
	}
}
