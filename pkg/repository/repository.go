package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/krinya32/pizzaApp"
)

type Pizzas interface {
	GetAll() ([]pizzaApp.PizzaStruct, error)
	GetById(id int) (pizzaApp.PizzaStruct, error)
}

type Repository struct {
	Pizzas
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Pizzas: NewPizzaPostgres(db),
	}
}
