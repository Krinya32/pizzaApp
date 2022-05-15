package repository

type Pizzas interface {
}

type Repository struct {
	Pizzas
}

func NewRepository() *Repository {
	return &Repository{}
}
