// Package main provides ...
package main

import (
	"fmt"
	// "os"
	// "strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func DemoRebalanceCb(c *kafka.Consumer, ev kafka.Event) error {
	fmt.Println(ev)
	return nil
}

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "127.0.0.1",
		"group.id":           "test",
		"auto.offset.reset":  "latest",
		"enable.auto.commit": "false",
	})

	if err != nil {
		panic(err)
	}

	if err = c.Subscribe("test", DemoRebalanceCb); err != nil {
		panic(err)
	}

	/*
	 * topic := os.Args[1]
	 * partition, _ := strconv.Atoi(os.Args[2])
	 * tp := kafka.TopicPartition{
	 *     Topic:     &topic,
	 *     Partition: int32(partition),
	 *     Offset:    kafka.OffsetStored,
	 * }
	 * partitions := []kafka.TopicPartition{tp}
	 * if err := c.Assign(partitions); err != nil {
	 *     panic(err)
	 * }
	 */

	partitions, err := c.Assignment()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c.String())

	for {
		// msg, err := c.ReadMessage(-1)
		ev := c.Poll(-1)

		switch e := ev.(type) {
		case *kafka.Message:
			if e.TopicPartition.Error != nil {
				fmt.Println(e.TopicPartition.Error)
			}
			fmt.Printf("Partition %s, Topic %s, Key: %s, Value: %s \t", partitions, e.TopicPartition, string(e.Key), string(e.Value))
		case kafka.Error:
			fmt.Printf("Error %s", e.Error())
		default:
			// Ignore other event types
		}

		partitions, err := c.Commit()
		fmt.Printf("Partition %s\n", partitions)
		if err != nil {
			fmt.Println(err)
		}
	}

	c.Close()

}
