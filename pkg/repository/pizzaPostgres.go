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

func (r *PizzaPostgres) Create(pizza pizzaApp.PizzaStruct) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createPizzaQuery := fmt.Sprintf("INSERT INTO %s (title, price, description, spicy, available) VALUES ($1, $2, $3, $4, $5) RETURNING id", pizzaVariable)
	row := tx.QueryRow(createPizzaQuery, pizza.Title, pizza.Price, pizza.Description, pizza.Spicy, pizza.Available)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *PizzaPostgres) GetAll() ([]pizzaApp.PizzaStruct, error) {
	var pizzas []pizzaApp.PizzaStruct
	query := fmt.Sprintf("SELECT * FROM %s", pizzaVariable)
	err := r.db.Select(&pizzas, query)
	return pizzas, err
}

func (r *PizzaPostgres) GetById(id int) (pizzaApp.PizzaStruct, error) {
	var pizza pizzaApp.PizzaStruct
	query := fmt.Sprintf("SELECT pizzas.id, pizzas.title, pizzas.price, pizzas.description, pizzas.spicy, pizzas.available FROM %s WHERE id = $1", pizzaVariable)
	err := r.db.Get(&pizza, query, id)
	return pizza, err
}
