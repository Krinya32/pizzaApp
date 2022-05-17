package pizzaApp

import "errors"

type PizzaStruct struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Price       int    `json:"price" db:"price"`
	Description string `json:"description" db:"description"`
	Spicy       bool   `json:"spicy" db:"spicy"`
	Available   bool   `json:"available" db:"available"`
}

type UpdatePizzaInput struct {
	Title       *string `json:"title"`
	Price       *int    `json:"price"`
	Description *string `json:"description"`
	Spicy       *bool   `json:"spicy"`
	Available   *bool   `json:"available"`
}

func (i UpdatePizzaInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Price == nil && i.Spicy == nil && i.Available == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
