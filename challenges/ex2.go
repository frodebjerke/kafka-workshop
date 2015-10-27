package challenges

import (
    "fmt"

    "github.com/Shopify/sarama"

    "github.com/frodebjerke/kafka-workshop/setup"
)

func Ex2() {
    fmt.Println("Create topic for exercise 2")
    producer := setup.GetProducer()

    for _, jedi := range Jedis {
        producer.SendMessage(&sarama.ProducerMessage{
            Topic: "starwars",
            Partition: 0,
            Value: sarama.StringEncoder(jedi),
        })
    }

    producer.Close()
}
