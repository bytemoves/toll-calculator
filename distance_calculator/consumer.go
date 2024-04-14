package main

import (
	"encoding/json"
	"time"

	// "time"

	"github.com/bytemoves/toll-calculator/aggregator/client"
	"github.com/bytemoves/toll-calculator/types"
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/sirupsen/logrus"
)

//	type DataConsumer interface{
//		ConsumeData ()
//	}
//
// this can be called kafka transport
type KafkaConsumer struct {
	consumer    *kafka.Consumer
	isRunning   bool
	calcService CalculatorServicer
	aggClient *client.HTTPClient
}

func NewKafkaConsumer(topic string, svc CalculatorServicer,aggClient *client.HTTPClient) (*KafkaConsumer, error) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	c.SubscribeTopics([]string{topic}, nil)

	return &KafkaConsumer{
		consumer:    c,
		calcService: svc,
		aggClient: aggClient,
	}, nil
}
func (c *KafkaConsumer) Start() {
	logrus.Info("Kafka tranport started")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {

	for c.isRunning {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("kafka consumer error %s", err)
			continue
		}
		var data types.OBUData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("JSON serialization error: %s", err)
			continue
		}
		distance, err := c.calcService.calculateDistance(data)
		if err != nil {
			logrus.Errorf("Calculation error: %s", err)
			continue
		}
		req := types.Distance{
			Value:distance,
			Unix: time.Now().UnixNano(),
			OBUID: data.OBUID,
		}
		if err := c.aggClient.AggregateInvoice(req); err != nil{
			logrus.Errorf("aggregate error: %s",err)
			continue

		}
	}

}


