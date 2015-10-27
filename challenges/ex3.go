package challenges

import (
    "fmt"

    "github.com/Shopify/sarama"

    "github.com/frodebjerke/kafka-workshop/setup"
)

func Ex3() {
    fmt.Println("Create topic for exercise 3")
    producer := setup.GetProducer()

    for _, jedi := range Jedis {
        p, o, _ := producer.SendMessage(&sarama.ProducerMessage{
            Topic: "ex3",
            Key: sarama.StringEncoder("jedis"),
            Value: sarama.StringEncoder(jedi),
        })

        fmt.Println(p, o)
    }

    for _, sith := range Siths {
        p, o, _ := producer.SendMessage(&sarama.ProducerMessage{
            Topic: "ex3",
            Key: sarama.StringEncoder("siths"),
            Value: sarama.StringEncoder(sith),
        })
        
        fmt.Println(p, o)
    }

    producer.Close()
}
