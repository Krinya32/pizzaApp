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

func (s *PizzaService) Create(pizza pizzaApp.PizzaStruct) (int, error) {
	return s.repo.Create(pizza)
}

func (s *PizzaService) GetAll() ([]pizzaApp.PizzaStruct, error) {
	return s.repo.GetAll()
}

func (s *PizzaService) GetById(id int) (pizzaApp.PizzaStruct, error) {
	return s.repo.GetById(id)
}

func (s *PizzaService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *PizzaService) Update(id int, input pizzaApp.UpdatePizzaInput) error {
	if err := input.Validate(); err != nil {
		return nil
	}

	return s.repo.Update(id, input)
}
