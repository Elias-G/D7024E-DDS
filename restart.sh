#!/bin/bash
docker swarm leave -f
docker swarm init
docker build -t test ./src
docker build -t test2 .
docker stack deploy up -c docker-compose.yml