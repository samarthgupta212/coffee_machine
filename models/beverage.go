package models

type Beverage struct {
	Name        string
	Ingredients []Ingredient
}

func NewBeverage(name string, ingredients []Ingredient) Beverage {
	return Beverage{Name: name, Ingredients: ingredients}
}
