# Commandline Challenges

Use the commandline tool [kafkacat](https://github.com/edenhill/kafkacat#install) to solved the following challenges.

[Install kafkacat](https://github.com/edenhill/kafkacat#install)

## First challenge
*In this challenge you will learn how to first write and then read one message with Kafka, using the cli-tool Kafkacat.*

1. Write a single message containing just a simple 'hello' to a topic called ex1.
1. Read all messages from the ex1 topic.

   *Should yield*
   ````
   hello
   ````

## Second challenge
*In this challenge you will learn to lookup information via kafkacat and actively controlling your consumers offset-pointer.*

1. Find all topics existing on the Kafka broker.
1. Read the two last messages from the topic named after a movie.

    *Should yield*
    ````
Yoda
Luke Skywalker
    ````
1. Read the first message from the topic named after a movie.

    *Should yield*
    ````
    Obi-Wan Kenobi
    ````

## Third challenge
*In this challenge you will get a feel of the implications of partitioning.*

1. Write the message 'hello' to partition 0 of the topic ex3
1. Write the message 'world' to partition 1 of the topic ex3
1. Read all messages on the topic ex3 without specifying partition

    *Should yield*
    ````
hello
world
    ````
1. Read all messages in partition 1 on the topic ex3

    *Should yield*
    ````
    world
    ````
