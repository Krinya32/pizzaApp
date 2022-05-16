package service

import (
	"github.com/krinya32/pizzaApp"
	"github.com/krinya32/pizzaApp/pkg/repository"
)

type Pizzas interface {
	GetAll() ([]pizzaApp.PizzaStruct, error)
	GetById(id int) (pizzaApp.PizzaStruct, error)
}

type Service struct {
	Pizzas
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Pizzas: NewPizzaService(repos.Pizzas),
	}
}
