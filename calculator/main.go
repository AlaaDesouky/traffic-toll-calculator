package main

import (
	"log"
	"traffic-toll-calculator/types"
)

func main() {
	var (
		err error
		svc CalculatorServicer
	)

	svc = NewCalculatorService()
	kafkaConsumer, err := NewKafkaConsumer(types.OBUKafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
