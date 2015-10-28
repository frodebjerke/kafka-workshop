package main

import (
	sarama "gopkg.in/Shopify/sarama.v1"
	"log"
	"math/rand"
	"time"
	"sync"
)

var kafkaHosts = []string{
	"192.168.59.103:9092",
}

const starwarsTopic = "starwars"

const randomSleep = 5000

var jedis = []string{
	"Darth Maul",
	"Palpatine",
	"Emperor",
	"Qui-Gon Jinn",
	"Anakin Skywalker",
	"Obi-Wan Kenobi",
	"Mace Windu",
	"Yoda",
	"Luke Skywalker",
}

var wg = sync.WaitGroup{}

func main() {
	config := sarama.NewConfig()


	done := make(chan bool)

	wg.Add(1)
	go RunKafkaConsumer(done, config)
	wg.Add(1)
	go randomKafkaProducer(done, config)

	// wait for done
	wg.Wait()
}

func randomKafkaProducer(done chan bool, config *sarama.Config) {
	for {
		sleep := time.Millisecond * time.Duration(rand.Int31n(randomSleep))
		select {
		case <- time.After(sleep):
			random := rand.Int() % len(jedis)

			go func(t time.Time) {
				RunKafkaProducer(done, config, jedis[random])
				log.Printf("producing took %v", time.Since(t))
			}(time.Now())
		case <- done:
			wg.Done()
		}
	}
}

// RunKafkaProducer creates a new producer for every call.
func RunKafkaProducer(done chan bool, config *sarama.Config, value string) {
	prod, err := sarama.NewSyncProducer(kafkaHosts, config)
	if err != nil {
		log.Fatal(err)
	}

	msg := sarama.ProducerMessage{
		Topic: starwarsTopic,
		//Key: // apparently don't need this!
		Value: sarama.StringEncoder(value),
	}
	part, offset, err := prod.SendMessage(&msg)
	if err != nil {
		log.Printf("error sending kafka msg: %v", err)
	}
	log.Printf("produced: part: %v offset: %v", part, offset)
}

//
func RunKafkaConsumer(done chan bool, config *sarama.Config) {
	cons, err := sarama.NewConsumer(kafkaHosts, config)
	if err != nil {
		log.Fatal(err)
	}
	defer cons.Close()

	topics, err := cons.Topics()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Found topics:")
	for _, topic := range topics {
		log.Printf("\t- %v", topic)
	}

	ids, err := cons.Partitions(starwarsTopic)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Partitions for topic: %v", starwarsTopic)
	for _, id := range ids {
		log.Printf("\t- %v", id)
	}

	partConsumer, err := cons.ConsumePartition(starwarsTopic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}
	defer partConsumer.Close()
	consume(done, partConsumer.Messages())

}

//
func consume(done chan bool, c <-chan *sarama.ConsumerMessage) {
	count := 0
	for {
		select {
		case msg := <-c:
			key := string(msg.Key)
			value := string(msg.Value)
			log.Printf("got msg key: %v value: %v", key, value)
			count = count + 1
		case <-done:
			log.Printf("no more messages. Got: %v", count)
			return
		}
	}
}
