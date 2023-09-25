package main

import (
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/lewiscasewell/tolling/aggregator/client"
	"github.com/lewiscasewell/tolling/types"
	"github.com/sirupsen/logrus"
)

type KafkaConsumer struct {
	consumer         *kafka.Consumer
	isRunning        bool
	calcService      CalculatorServicer
	aggregatorClient *client.Client
}

func NewKafkaConsumer(topic string, svc CalculatorServicer, agClient client.Client) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "distance-calculator",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	c.SubscribeTopics([]string{topic}, nil)

	return &KafkaConsumer{
		consumer:         c,
		isRunning:        false,
		calcService:      svc,
		aggregatorClient: &agClient,
	}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("Starting Kafka Transport")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunning {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("Kafka consumer error: %v (%v)\n", err, msg)
			continue
		}

		var data types.OBUData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("Error unmarshalling message: %v", err)
			continue
		}

		distance, err := c.calcService.CalculateDistance(data)
		if err != nil {
			logrus.Errorf("Error calculating distance: %v", err)
			continue
		}

		req := types.Distance{
			OBUID: data.OBUID,
			Value: distance,
			Unix:  time.Now().UnixNano(),
		}

		if err := c.aggregatorClient.AggregateInvoice(req); err != nil {
			logrus.Errorf("Error aggregating invoice: %v", err)
			continue
		}
	}
}
