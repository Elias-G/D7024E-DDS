#!/bin/bash
docker swarm leave -f
docker swarm init
docker network create --driver overlay --subnet 10.0.0.0/24 --gateway 10.0.0.1 kademlia-network
docker build -t test ./src
docker build -t test2 .
docker stack deploy up -c docker-compose.yml