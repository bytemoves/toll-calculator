package main

import (
	"time"

	"github.com/bytemoves/toll-calculator/types"
	"github.com/sirupsen/logrus"
)


type LogMiddleware struct{
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator{

	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) AggregateDistance(distance types.Distance) (err error){
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err, 
			
		}).Info("Aggregate Distance")
	}(time.Now())
	err = m.next.AggregateDistance(distance)
	return
}


func (m *LogMiddleware) CalculateInvoice(obuID int) (inv *types.Invoice, err error){
	defer func(start time.Time) {
		var (
			distance float64
			amount float64
		)
		if inv != nil{
			distance = inv.TotalDistance
			amount = inv.TotalAmount
		}
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err, 
			"obuID":obuID,
			"distance":distance,
			"amount":amount,
			
		}).Info("Calculate invoice")
	}(time.Now())
	inv,err = m.next.CalculateInvoice(obuID)
	return
}