package consumer

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	Client         sarama.ConsumerGroup
	ConsumedEvents chan []byte
}

func NewConsumer(ctx context.Context, topics []string, client sarama.ConsumerGroup) *Consumer {
	eventsChannel := make(chan []byte)
	consumer := Consumer{
		Client:         client,
		ConsumedEvents: eventsChannel,
	}
	go processEvents(ctx, &consumer, topics)
	return &consumer
}

func processEvents(ctx context.Context, consumer *Consumer, topics []string) {
	log.Printf("consuming the following topics: %s", topics)
	for {
		if ctx.Err() != nil {
			log.Printf("server context has been cancelled; stopping consumption")
			return
		}
		if err := consumer.Client.Consume(context.Background(), topics, consumer); err != nil {
			log.Panicf("Error from consumer: %v", err)
		}
	}
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	close(consumer.ConsumedEvents)
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	log.Printf("listening for events in topic: %s; offset: %d; high-watermark offset: %d",
		claim.Topic(),
		claim.InitialOffset(), claim.InitialOffset(),
	)
	for message := range claim.Messages() {
		log.Printf("message received: %v", message.Value)
		consumer.ConsumedEvents <- message.Value
		session.MarkMessage(message, "")
	}
	return nil
}
