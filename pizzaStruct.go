package pizzaApp

type Pizza struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Spicy       bool    `json:"spicy"`
	Size        string  `json:"size"`
	Available   bool    `json:"available"`
}
