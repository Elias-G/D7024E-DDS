#!/bin/bash
docker-compose down
docker system prune && docker rmi kademlianodes:latest && docker rmi rootnode:latest && docker network prune
docker build -t rootnode .
docker build -t kademlianodes ./src
COMPOSE_PARALLEL_LIMIT=100 && docker-compose up --scale kademliaNodes=50 --scale rootnode=1