package pizzaApp

type PizzaStruct struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Spicy       bool    `json:"spicy"`
	Available   bool    `json:"available"`
}
