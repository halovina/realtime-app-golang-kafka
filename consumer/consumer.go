package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
	"github.com/halovina/realtime-app-golang-kafka/config"
)

func main() {
	newConfig := config.NewConfig()
	consumer, err := sarama.NewConsumer(config.KafkaBrokers, newConfig)
	if err != nil {
		log.Fatalln("Failed to start Sarama consumer: ", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Println("Failed to close Sarama consumer: ", err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(config.KafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("Failed to start consumer for partition 0", err)
	}
	defer partitionConsumer.AsyncClose()

	go func() {
		for msg := range partitionConsumer.Messages() {
			fmt.Printf("Received message: %s\n", string(msg.Value))
		}
	}()

	time.Sleep(time.Second * 60)
}
