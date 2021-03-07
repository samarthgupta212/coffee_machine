package main

import (
	"coffee_machine/consumer"
	"coffee_machine/models"
	"coffee_machine/producer"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Outlets struct {
	CountN int `json:"count_n"`
}

// MachineInfo struct - expected structure in input file
type MachineInfo struct {
	Outlets            Outlets                   `json:"outlets"`
	LowCount           int                       `json:"low_count"`
	TotalItemsQuantity map[string]int            `json:"total_items_quantity"`
	Beverages          map[string]map[string]int `json:"beverages"`
}

type Machine struct {
	Machine MachineInfo `json:"machine"`
}

func getMachineInfoFromJSON(fileName string) Machine {
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var machine Machine
	// we unmarshal our byteArray which contains our machineInfo
	_ = json.Unmarshal(byteValue, &machine)

	return machine
}

func getBeverageFromMachine(machine Machine) []models.Beverage {
	beverages := make([]models.Beverage, 0)
	for beverageName, value := range machine.Machine.Beverages {
		ingredients := make([]models.Ingredient, 0)
		for ingredientName, qty := range value {
			ingredients = append(ingredients, models.NewIngredient(ingredientName, qty))
		}
		beverage := models.NewBeverage(beverageName, ingredients)
		beverages = append(beverages, beverage)
	}
	return beverages
}

func serverBeverages(coffeeMachine models.CoffeeMachine, beverages []models.Beverage) {
	jobs := make(chan models.Beverage)
	newProducer := producer.NewProducer(jobs)
	newConsumer := consumer.NewConsumer(jobs)
	go newProducer.Produce(beverages)
	newConsumer.Consume(coffeeMachine)
}

// Files are run in order - input.json and input1.json
// input.json represents first inputs and input1.json uses refill method to refill the ingredients and then serve the beverages again
func main() {
	machine := getMachineInfoFromJSON("input.json")
	var coffeeMachine models.CoffeeMachine
	coffeeMachine = models.NewCoffeeMachine(machine.Machine.Outlets.CountN, machine.Machine.LowCount)
	// Add ingredients
	for key, value := range machine.Machine.TotalItemsQuantity {
		coffeeMachine.AddIngredient(key, value)
	}

	beverages := getBeverageFromMachine(machine)
	// Serve beverages
	serverBeverages(coffeeMachine, beverages)

	coffeeMachine.GetLowIngredients()
	// Input1.json contains items which will be used to refill
	machine1 := getMachineInfoFromJSON("input1.json")

	for key, value := range machine1.Machine.TotalItemsQuantity {
		coffeeMachine.RefillIngredient(key, value)
	}
	beverages = getBeverageFromMachine(machine1)
	// Serve beverages again
	serverBeverages(coffeeMachine, beverages)

}
