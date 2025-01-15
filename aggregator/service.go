package main

import "traffic-toll-calculator/types"

type Aggregator interface {
	AggregateDistance(types.Distance) error
	CalculateInvoice(int) (*types.Invoice, error)
}

type AggregatorService struct {
	store Storer
}

func NewAggregatorService(store Storer) Aggregator {
	return &AggregatorService{
		store: store,
	}
}

func (agg *AggregatorService) AggregateDistance(distance types.Distance) error {
	return nil
}

func (agg *AggregatorService) CalculateInvoice(obuID int) (*types.Invoice, error) {
	return nil, nil
}
