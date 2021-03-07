package main

import (
	"coffee_machine/models"
	"reflect"
	"sort"
	"testing"
)

func TestAddIngredient(t *testing.T) {
	coffeeMachine := models.NewCoffeeMachine(2, 10)
	coffeeMachine.AddIngredient("water", 10)
	coffeeMachine.AddIngredient("milk", 20)

	if coffeeMachine.IngredientMap["water"] != 10 {
		t.Errorf("Expected water Quantity to be %d", 10)
	}
	if coffeeMachine.IngredientMap["milk"] != 20 {
		t.Errorf("Expected milk Quantity to be %d", 20)
	}
}

func TestServeBeverage(t *testing.T) {
	coffeeMachine := models.NewCoffeeMachine(2, 10)
	coffeeMachine.AddIngredient("water", 100)
	coffeeMachine.AddIngredient("milk", 200)
	coffeeMachine.AddIngredient("syrup", 100)
	coffeeMachine.AddIngredient("mixture", 100)

	ingredient1 := models.NewIngredient("water", 50)
	ingredient2 := models.NewIngredient("milk", 100)
	hotTeaIngredients := []models.Ingredient{ingredient1, ingredient2}
	beverage1 := models.NewBeverage("hot_tea", hotTeaIngredients)

	ingredient3 := models.NewIngredient("syrup", 100)
	greenTeaIngredients := []models.Ingredient{ingredient1, ingredient2, ingredient3}
	beverage2 := models.NewBeverage("green_tea", greenTeaIngredients)

	ingredient4 := models.NewIngredient("black_mixture", 100)
	blackTeaIngredients := []models.Ingredient{ingredient1, ingredient2, ingredient4}
	beverage3 := models.NewBeverage("green_tea", blackTeaIngredients)

	beverages := []models.Beverage{beverage1, beverage2, beverage3}
	serverBeverages(coffeeMachine, beverages)
	if coffeeMachine.IngredientMap["water"] != 0 {
		t.Errorf("Expected water Quantity to be %d", 0)
	}
	if coffeeMachine.IngredientMap["milk"] != 0 {
		t.Errorf("Expected milk Quantity to be %d", 0)
	}
	if coffeeMachine.IngredientMap["syrup"] != 0 {
		t.Errorf("Expected syrup Quantity to be %d", 0)
	}
	if coffeeMachine.IngredientMap["mixture"] != 100 {
		t.Errorf("Expected mixture Quantity to be %d", 100)
	}

	coffeeMachine = models.NewCoffeeMachine(2, 10)
	coffeeMachine.AddIngredient("water", 100)
	coffeeMachine.AddIngredient("milk", 200)
	coffeeMachine.AddIngredient("syrup", 100)
	coffeeMachine.AddIngredient("mixture", 100)

	ingredient5 := models.NewIngredient("mixture", 50)
	mixtureIngredients := []models.Ingredient{ingredient5}
	beverage4 := models.NewBeverage("green_mixture", mixtureIngredients)
	beverages = append(beverages, beverage4)

	ingredient6 := models.NewIngredient("mixture", 50)
	redTeaIngredients := []models.Ingredient{ingredient1, ingredient6}
	beverage5 := models.NewBeverage("red_tea", redTeaIngredients)
	beverages = append(beverages, beverage5)

	serverBeverages(coffeeMachine, beverages)
	if coffeeMachine.IngredientMap["water"] != 0 {
		t.Errorf("Expected water Quantity to be %d", 0)
	}
	if coffeeMachine.IngredientMap["milk"] != 0 {
		t.Errorf("Expected milk Quantity to be %d", 0)
	}
	if coffeeMachine.IngredientMap["syrup"] != 0 {
		t.Errorf("Expected syrup Quantity to be %d", 0)
	}
	if coffeeMachine.IngredientMap["mixture"] != 50 {
		t.Errorf("Expected mixture Quantity to be %d", 50)
	}

	lowIngredients := coffeeMachine.GetLowIngredients()
	sort.Strings(lowIngredients)
	expectedLowIngredients := []string{"milk", "syrup", "water"}
	if !reflect.DeepEqual(lowIngredients, expectedLowIngredients) {
		t.Errorf("mis-match: %v %v", lowIngredients, expectedLowIngredients)
	}
}

func TestRefillIngredient(t *testing.T) {
	coffeeMachine := models.NewCoffeeMachine(2, 10)
	coffeeMachine.AddIngredient("water", 10)
	coffeeMachine.AddIngredient("milk", 20)

	coffeeMachine.RefillIngredient("syrup", 20)
	coffeeMachine.RefillIngredient("water", 20)

	if coffeeMachine.IngredientMap["water"] != 30 {
		t.Errorf("Expected water Quantity to be %d", 10)
	}
	if coffeeMachine.IngredientMap["milk"] != 20 {
		t.Errorf("Expected milk Quantity to be %d", 20)
	}
	if coffeeMachine.IngredientMap["syrup"] != 20 {
		t.Errorf("Expected syrup Quantity to be %d", 20)
	}
}