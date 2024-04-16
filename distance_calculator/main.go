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
	httpClient :=  client.NewHTTPClient(aggregatorEndpoint)
	// grpcClient, err :=  client.NewGRPCClient(aggregatorEndpoint)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	
	KafkaConsumer , err := NewKafkaConsumer(kafkaTopic, svc, httpClient) 
	if err !=nil{
		log.Fatal(err)
	}
	KafkaConsumer.Start()
	
}