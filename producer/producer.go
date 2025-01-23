package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/halovina/realtime-app-golang-kafka/config"
)

func main() {
	newConfig := config.NewConfig()
	producer, err := sarama.NewSyncProducer(config.KafkaBrokers, newConfig)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer: ", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Println("Failed to close Sarama producer: ", err)
		}
	}()

	message := &sarama.ProducerMessage{
		Topic: config.KafkaTopic,
		Value: sarama.StringEncoder("Hello, Kafka!"),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Println("Failed to send message: ", err)
	} else {
		fmt.Printf("Message is stored in partition %d at offset %d\n", partition, offset)
	}
}
