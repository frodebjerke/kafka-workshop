package setup

import (
    "log"
    "time"

    "github.com/Shopify/sarama"
)

var (
    brokerList = []string{"192.168.99.100:9092"}
)

func GetProducer() sarama.SyncProducer {
    config := sarama.NewConfig()

    config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond

    producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	return producer
}

func GetConsumer() sarama.Consumer {

    config := sarama.NewConfig()

    consumer, err := sarama.NewConsumer(brokerList, config)
    if err != nil {
        log.Fatalln("Failed to start Sarama consumer", err)
    }

    return consumer
}
