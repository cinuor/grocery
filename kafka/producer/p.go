package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"time"
)

const TOPIC = "test"

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1",
	})

	if err != nil {
		panic(err)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				fmt.Println(ev)
			default:
				fmt.Println(ev)
			}
		}
	}()

	topic := "test"

	for i := 0; i < 100; i++ {
		key := uuid.NewString()
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(fmt.Sprintf("message value from %d", i)),
			Key:            []byte(key),
		}
		// p.ProduceChannel() <- msg
		err := p.Produce(msg, nil)
		if err != nil {
			fmt.Println(err)
		}
	}

	time.Sleep(5 * time.Second)
}
