package service

import (
	"github.com/krinya32/pizzaApp"
	"github.com/krinya32/pizzaApp/pkg/repository"
)

type Pizzas interface {
	Create(pizza pizzaApp.PizzaStruct) (int, error)
	GetAll() ([]pizzaApp.PizzaStruct, error)
	GetById(id int) (pizzaApp.PizzaStruct, error)
	Delete(id int) error
	Update(id int, input pizzaApp.UpdatePizzaInput) error
}

type Service struct {
	Pizzas
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Pizzas: NewPizzaService(repos.Pizzas),
	}
}
