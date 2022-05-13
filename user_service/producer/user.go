package producer

import (
	"log"

	"github.com/Shopify/sarama"
)

const (
	UserRegisteredTopic = "userRegistered"
)

type IUserProducer interface {
	CreateJsonEvent(topic string, value []byte)
	getPartition() int32
	Close()
}

type UserProducer struct {
	producerConn sarama.SyncProducer
}

func NewUserProducer(brokers []string) *UserProducer {
	producer := newProducer(brokers)
	return &UserProducer{
		producerConn: producer,
	}
}
func (producer UserProducer) CreateJsonEvent(topic string, value []byte) {
	partion := producer.getPartition()
	resultPartion, offSet, err := producer.producerConn.SendMessage(&sarama.ProducerMessage{
		Topic:     topic,
		Partition: partion,
		Value:     sarama.ByteEncoder(value),
	})
	log.Printf("record sent with partion: %d; Offset: %d", resultPartion, offSet)
	if err != nil {
		log.Printf("record sent with error: %s", err.Error())
	}
}

func (UserProducer) getPartition() int32 {
	return 1
}

func (producer *UserProducer) Close() {
	producer.producerConn.Close()
}
