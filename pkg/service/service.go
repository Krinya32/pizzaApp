package service

import "github.com/krinya32/pizzaApp/pkg/repository"

type Pizzas interface {
}

type Service struct {
	Pizzas
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
