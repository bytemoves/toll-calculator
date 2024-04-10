package main

import (
	
	"log"
	// "time"
)

const kafkaTopic = "obudata"
/// Transport (http,grpc kafka) --> business logic to transport


func main () {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)
	
	KafkaConsumer , err := NewKafkaConsumer(kafkaTopic,svc)
	if err !=nil{
		log.Fatal(err)
	}
	KafkaConsumer.Start()
	
}