#!/bin/bash
docker swarm leave -f
docker swarm init
docker build -t test .
docker stack deploy up -c docker-compose.yml

