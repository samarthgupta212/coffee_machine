package producer

import "coffee_machine/models"

type Producer struct {
	jobs chan models.Beverage
}

func NewProducer(jobs chan models.Beverage) Producer {
	return Producer{jobs: jobs}
}

func (p *Producer) Produce(beverages []models.Beverage) {
	for _, beverage := range beverages {
		p.jobs <- beverage
	}
	close(p.jobs)
}
