package models

type Ingredient struct {
	Name     string
	Quantity int
}

func NewIngredient(name string, quantity int) Ingredient {
	return Ingredient{Name: name, Quantity: quantity}
}
