package service

import (
	"github.com/krinya32/pizzaApp"
	"github.com/krinya32/pizzaApp/pkg/repository"
)

type PizzaService struct {
	repo repository.Pizzas
}

func NewPizzaService(repo repository.Pizzas) *PizzaService {
	return &PizzaService{repo: repo}
}

func (s *PizzaService) GetAll() ([]pizzaApp.PizzaStruct, error) {
	return s.repo.GetAll()
}

func (s *PizzaService) GetById(id int) (pizzaApp.PizzaStruct, error) {
	return s.repo.GetById(id)
}
