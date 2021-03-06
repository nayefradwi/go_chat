package producer

import (
	"crypto/tls"
	"log"

	"github.com/Shopify/sarama"
)

func newProducerConn(brokerList []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.ClientID = "gochat-user-service"
	config.Producer.RequiredAcks = sarama.WaitForLocal // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                     // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true
	tlsConfig := createTlsConfiguration()
	if tlsConfig != nil {
		config.Net.TLS.Config = tlsConfig
		config.Net.TLS.Enable = true
	}
	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	log.Printf("connected to brokers %v", brokerList)
	return producer
}

func createTlsConfiguration() *tls.Config {
	return nil
}
