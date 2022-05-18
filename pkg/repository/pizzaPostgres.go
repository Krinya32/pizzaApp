package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/krinya32/pizzaApp"
	"strings"
)

type PizzaPostgres struct {
	db *sqlx.DB
}

func NewPizzaPostgres(db *sqlx.DB) *PizzaPostgres {
	return &PizzaPostgres{db: db}
}

func (r *PizzaPostgres) Create(pizza pizzaApp.PizzaStruct) (int, error) {
	var lastInsertId = 0

	createPizzaQuery := fmt.Sprintf("INSERT INTO %s (title, price, description, spicy, available) VALUES ($1, $2, $3, $4, $5) RETURNING id", pizzaVariable)

	err := r.db.QueryRow(createPizzaQuery, pizza.Title, pizza.Price, pizza.Description, pizza.Spicy, pizza.Available).Scan(&lastInsertId)

	fmt.Printf("Created pizza with id: %d", lastInsertId)
	return lastInsertId, err
}

func (r *PizzaPostgres) GetAll() ([]pizzaApp.PizzaStruct, error) {
	var pizzas []pizzaApp.PizzaStruct
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC", pizzaVariable)
	err := r.db.Select(&pizzas, query)
	return pizzas, err
}

func (r *PizzaPostgres) GetById(id int) (pizzaApp.PizzaStruct, error) {
	var pizza pizzaApp.PizzaStruct
	query := fmt.Sprintf("SELECT pizzas.id, pizzas.title, pizzas.price, pizzas.description, pizzas.spicy, pizzas.available FROM %s WHERE id = $1", pizzaVariable)
	err := r.db.Get(&pizza, query, id)
	return pizza, err
}

func (r *PizzaPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", pizzaVariable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PizzaPostgres) Update(id int, input pizzaApp.UpdatePizzaInput) error {
	value := make([]string, 0)
	arguments := make([]interface{}, 0)
	argumentsId := 1

	if input.Title != nil {
		value = append(value, fmt.Sprintf("title=$%d", argumentsId))
		arguments = append(arguments, *input.Title)
		argumentsId++
	}

	if input.Price != nil {
		value = append(value, fmt.Sprintf("price=$%d", argumentsId))
		arguments = append(arguments, *input.Price)
		argumentsId++
	}

	if input.Description != nil {
		value = append(value, fmt.Sprintf("description=$%d", argumentsId))
		arguments = append(arguments, *input.Description)
		argumentsId++
	}

	if input.Spicy != nil {
		value = append(value, fmt.Sprintf("spicy=$%d", argumentsId))
		arguments = append(arguments, *input.Spicy)
		argumentsId++
	}

	if input.Available != nil {
		value = append(value, fmt.Sprintf("available=$%d", argumentsId))
		arguments = append(arguments, *input.Available)
		argumentsId++
	}

	SetQuery := strings.Join(value, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", pizzaVariable, SetQuery, argumentsId)
	arguments = append(arguments, id)

	_, err := r.db.Exec(query, arguments...)
	return err
}
