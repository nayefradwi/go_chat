package producer

import (
	"log"

	"github.com/Shopify/sarama"
)

type IProducer interface {
	CreateJsonEvent(topic string, value []byte)
	getPartition() int32
	Close()
}

type Producer struct {
	producerConn sarama.SyncProducer
}

func NewProducer(brokers []string) *Producer {
	producer := newProducerConn(brokers)
	return &Producer{
		producerConn: producer,
	}
}

func (producer Producer) CreateJsonEvent(topic string, value []byte) {
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

func (Producer) getPartition() int32 {
	return 1
}

func (producer *Producer) Close() {
	producer.producerConn.Close()
}
