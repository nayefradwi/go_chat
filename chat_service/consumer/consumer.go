package consumer

import "github.com/Shopify/sarama"

func NewConsumer(brokerList []string) sarama.ConsumerGroup {
	return newConsumerConn(brokerList)
}
