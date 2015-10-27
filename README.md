# Kafka workshop

## Pre-requisites
1. docker ([install](https://docs.docker.com/installation/))
1. kafkacat (for the cmd-challenges) ([install](https://github.com/edenhill/kafkacat#install))

## Run

!! Before running `setup.sh` verify that your docker-ip matches the KAFKA_ADVERTISED_HOST_NAME variable found in the file `kafka-docker/docker-compose.yml`

**Default docker ip is set to 192.168.99.100. Set DOCKER_IP to specify another IP.**

Start kafka and populate
````
DOCKER_IP=127.0.0.1 ./setup.sh
````

Reset kafka data
````
DOCKER_IP=127.0.0.1 ./reset.sh
````

## Challenges

- Commandline challenges in file `cmd-challenges.md`
- Programming challenges in file `programming-challenges.md`
