#! /bin/sh

docker-compose -f kafka-docker/docker-compose.yml kill
docker-compose -f kafka-docker/docker-compose.yml rm
