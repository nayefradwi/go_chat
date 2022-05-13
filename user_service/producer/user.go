package producer

import (
	"github.com/Shopify/sarama"
)

const (
	UserRegisteredTopic = "userRegistered"
)

type IUserProducer interface {
	CreateJsonEvent(topic string, value []byte)
	getPartition() int32
}

type UserProducer struct {
	producerConn sarama.SyncProducer
}

func NewUserProducer(producerConn sarama.SyncProducer) UserProducer {
	return UserProducer{
		producerConn: producerConn,
	}
}
func (producer UserProducer) CreateJsonEvent(topic string, value []byte) {
	partion := producer.getPartition()
	producer.producerConn.SendMessage(&sarama.ProducerMessage{
		Topic:     topic,
		Partition: partion,
		Value:     sarama.ByteEncoder(value),
	})
}

func (UserProducer) getPartition() int32 {
	return 1
}
