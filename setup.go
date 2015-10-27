package main

import (
    "os"

    "github.com/frodebjerke/kafka-workshop/challenges"
)

func main() {
    args := os.Args

    kafkaHost := "192.168.99.100:9092"

    if len(args) == 2 {
        kafkaHost = args[1] + ":9092"
    }

    challenges.Ex2(kafkaHost)
    // challenges.Ex3()
}
