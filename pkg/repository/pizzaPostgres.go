package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/krinya32/pizzaApp"
)

type PizzaPostgres struct {
	db *sqlx.DB
}

func NewPizzaPostgres(db *sqlx.DB) *PizzaPostgres {
	return &PizzaPostgres{db: db}
}

func (r *PizzaPostgres) GetAll() ([]pizzaApp.PizzaStruct, error) {
	var pizzas []pizzaApp.PizzaStruct
	query := fmt.Sprintf("SELECT * FROM pizzas")
	err := r.db.Select(&pizzas, query)
	return pizzas, err
}

func (r *PizzaPostgres) GetById(id int) (pizzaApp.PizzaStruct, error) {
	var pizza pizzaApp.PizzaStruct
	query := fmt.Sprintf("SELECT pizzas.id FROM pizzas")
	err := r.db.Get(&pizza, query, id)
	return pizza, err
}
