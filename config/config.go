package config

import "github.com/IBM/sarama"

var (
	KafkaBrokers = []string{"localhost:29092"}
	KafkaTopic   = "test-kafka"
)

func NewConfig() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	return config
}
