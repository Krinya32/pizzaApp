package repository

import "github.com/jmoiron/sqlx"

type Pizzas interface {
}

type Repository struct {
	Pizzas
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
