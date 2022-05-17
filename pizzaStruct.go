package pizzaApp

type PizzaStruct struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Price       int    `json:"price" db:"price"`
	Description string `json:"description" db:"description"`
	Spicy       bool   `json:"spicy" db:"spicy"`
	Available   bool   `json:"available" db:"available"`
}
