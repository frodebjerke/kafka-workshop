# Kafka workshop

## Pre-requisites
1. docker ([install](https://docs.docker.com/installation/))
1. kafkacat (for the cmd-challenges) ([install](https://github.com/edenhill/kafkacat#install))

## Run

!! Before running `setup.sh` verify that your docker-ip matches the KAFKA_ADVERTISED_HOST_NAME variable found in the file `kafka-docker/docker-compose.yml`

Start kafka and populate
````
./setup.sh
````

Reset kafka data
````
./reset.sh
````

## Challenges

- Commandline challenges in file `cmd-challenges.md`
- Programming challenges in file `programming-challenges.md`
