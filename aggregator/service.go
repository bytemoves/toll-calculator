package main

import (
	"fmt"

	"github.com/bytemoves/toll-calculator/types"
)



type Aggregator interface{
	 AggregateDistance(types.Distance) error
	
}

type Storer interface{
	Insert(types.Distance) error
}

type InvoiceAggregator struct{
	store Storer

}

func NewInvoiceAggreagator(store Storer) *InvoiceAggregator{
	return &InvoiceAggregator{
		store: store,
	}
}

func(i *InvoiceAggregator) AggregateDistance(distance types.Distance) error{
	fmt.Println("Processing and inserting distance in the storage: ",distance)

	return i.store.Insert(distance)
	
}