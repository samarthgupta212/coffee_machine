package models

import (
	"fmt"
	"sync"
)

// CoffeeMachine contains fields relevant to coffee machine
type CoffeeMachine struct {
	Outlets       int
	mu            sync.Mutex
	IngredientMap map[string]int
	LowCount      int
}

// NewCoffeeMachine contains the coffee machine instance
func NewCoffeeMachine(outlets int, lowCount int) CoffeeMachine {
	ingredientMap := make(map[string]int)
	return CoffeeMachine{Outlets: outlets, IngredientMap: ingredientMap, LowCount: lowCount}
}

// AddIngredient adds ingredients to coffee machine
func (c *CoffeeMachine) AddIngredient(name string, quantity int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.IngredientMap[name] = quantity
}

func (c *CoffeeMachine) deductQuantity(ingredient Ingredient) {
	c.IngredientMap[ingredient.Name] -= ingredient.Quantity
}

// ServeBeverage takes beverage as input and deducts quantity of ingredients if possible
func (c *CoffeeMachine) ServeBeverage(wg *sync.WaitGroup, jobs <-chan Beverage) {
	defer wg.Done()

	for beverage := range jobs {
		// Lock so only one goroutine at a time can access the map c.IngredientMap
		c.mu.Lock()
		isQuantityAvailable := true
		isQuantitySufficient := true
		for _, ingredient := range beverage.Ingredients {
			if _, ok := c.IngredientMap[ingredient.Name]; !ok {
				fmt.Printf("%s cannot be prepared because %s is not available\n", beverage.Name, ingredient.Name)
				isQuantityAvailable = false
				break
			}
		}
		if isQuantityAvailable {
			for _, ingredient := range beverage.Ingredients {
				if c.IngredientMap[ingredient.Name] < ingredient.Quantity {
					fmt.Printf("%s cannot be prepared because item %s is not sufficient\n", beverage.Name, ingredient.Name)
					isQuantitySufficient = false
					break
				}
			}
		}
		if isQuantitySufficient && isQuantityAvailable {
			for i := 0; i < len(beverage.Ingredients); i++ {
				c.deductQuantity(beverage.Ingredients[i])
			}
			fmt.Printf("%s is prepared\n", beverage.Name)
		}
		c.mu.Unlock()
	}
}

// RefillIngredient refills ingredient with given quantity
func (c *CoffeeMachine) RefillIngredient(name string, quantity int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.IngredientMap[name] += quantity
}

// GetLowIngredients returns ingredients which are running low
func (c *CoffeeMachine) GetLowIngredients() []string {
	ingredients := make([]string, 0)
	for key, value := range c.IngredientMap {
		if value <= c.LowCount {
			ingredients = append(ingredients, key)
			fmt.Printf("Running Low on Ingredient %s: Count: %d\n", key, value)
		}
	}
	return ingredients
}
