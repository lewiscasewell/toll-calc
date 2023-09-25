package main

import (
	"fmt"

	"github.com/lewiscasewell/tolling/types"
)

const basePrice = 3.15

type Aggregator interface {
	AggregateDistance(distance types.Distance) error
	CalculateInvoice(int) (types.Invoice, error)
}

type Storer interface {
	Insert(types.Distance) error
	Get(int) (float64, error)
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	fmt.Println("processing and inserting distance in the storage: ", distance)
	return i.store.Insert(distance)
}

func (i *InvoiceAggregator) CalculateInvoice(obuID int) (types.Invoice, error) {
	fmt.Println("calculating invoice for obuID: ", obuID)
	dist, err := i.store.Get(obuID)
	if err != nil {
		return types.Invoice{}, err
	}
	return types.Invoice{
		OBUID:         obuID,
		TotalDistance: dist,
		TotalCost:     dist * basePrice,
	}, nil
}
