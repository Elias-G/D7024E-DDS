#!/bin/bash
docker swarm leave -f
docker swarm init
docker build -t test .
docker build -t test2 ./RootNode
docker stack deploy up -c docker-compose.yml

