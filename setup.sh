#! /bin/sh

docker-compose -f kafka-docker/docker-compose.yml up -d

docker build -t kafka-workshop-setupscripts .
docker run kafka-workshop-setupscripts kafka-workshop 
