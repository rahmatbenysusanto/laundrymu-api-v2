package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strconv"
	"time"
)

func Producer() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	topic := "helloworld"

	for i := 0; i < 100; i++ {
		key := strconv.Itoa(i)
		value := fmt.Sprintf("Hello %d", i)
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Key:   []byte(key),
			Value: []byte(value),
		}

		err := producer.Produce(msg, nil)
		if err != nil {
			panic(err)
		}
	}

	producer.Flush(5 * 1000)
}

func Consumer() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "golang",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}
	defer func(consumer *kafka.Consumer) {
		err := consumer.Close()
		if err != nil {
			panic(err)
		}
	}(consumer)

	err = consumer.Subscribe("helloworld", nil)
	if err != nil {
		panic(err)
	}

	for {
		message, err := consumer.ReadMessage(1 * time.Second)
		if err == nil {
			fmt.Printf("Receive message : %s\n", message.Value)
		}
	}
}
