package main

import (
	"fmt"

	"github.com/bytemoves/toll-calculator/types"
	"github.com/sirupsen/logrus"
)


const basePrice = 3.15
type Aggregator interface{
	 AggregateDistance(types.Distance) error
	 CalculateInvoice(int) (*types.Invoice , error)
}

type Storer interface{
	Insert(types.Distance) error
	Get(int) (float64, error)
}

type InvoiceAggregator struct{
	store Storer

}

func NewInvoiceAggreagator(store Storer) Aggregator{
	return &InvoiceAggregator{
		store: store,
	}
}

func(i *InvoiceAggregator) AggregateDistance(distance types.Distance) error{
	fmt.Println("Processing and inserting distance in the storage: ",distance)
	logrus.WithFields(logrus.Fields{
		"obuid": distance.OBUID,
		"distance": distance.Value,
		"unix": distance.Unix,

	}).Info("aggrrgating distance")

	return i.store.Insert(distance)
	
}

func(i *InvoiceAggregator) CalculateInvoice(obuID int) (*types.Invoice,error){
	//fmt.Println("Processing and inserting distance in the storage: ",distance)
	dist,err := i.store.Get(obuID)
	if err != nil {
		return nil , err
	}
	inv := &types.Invoice{
		OBUID: obuID,
		TotalDistance: dist,
		TotalAmount:basePrice*dist ,
	}
	return inv , nil
	
}