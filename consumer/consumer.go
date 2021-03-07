package consumer

import (
	"coffee_machine/models"
	"sync"
)

type Consumer struct {
	jobs chan models.Beverage
}

func NewConsumer(jobs chan models.Beverage) Consumer {
	return Consumer{jobs: jobs}
}

func (c *Consumer) Consume(coffeeMachine models.CoffeeMachine) {
	wg := sync.WaitGroup{}
	// This makes sure that code is run parallel with number of outlets defined
	for i := 1; i <= coffeeMachine.Outlets; i++ {
		wg.Add(1)
		go coffeeMachine.ServeBeverage(&wg, c.jobs)
	}
	wg.Wait()
}
