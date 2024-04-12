package main

import (
	"log"

	"github.com/bytemoves/toll-calculator/aggregator/client"
	// "time"
)

const (
	kafkaTopic = "obudata"
	aggregatorEndpoint = "http://127.0.0.1:3000/aggregate"




)
/// Transport (http,grpc kafka) --> business logic to transport


func main () {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)
	
	KafkaConsumer , err := NewKafkaConsumer(kafkaTopic,svc,client.NewClient(aggregatorEndpoint))
	if err !=nil{
		log.Fatal(err)
	}
	KafkaConsumer.Start()
	
}