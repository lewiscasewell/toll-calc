package main

import "github.com/lewiscasewell/tolling/aggregator/client"

const (
	kafkaTopic         = "obu-data"
	aggregatorEndpoint = "http://127.0.0.1:3030/aggregate"
)

func main() {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, *client.NewClient(aggregatorEndpoint))
	if err != nil {
		panic(err)
	}
	kafkaConsumer.Start()
}
