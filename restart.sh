#!/bin/bash
docker-compose down
docker build -t rootnode .
docker build -t kademlianodes ./src
docker-compose up --scale kademliaNodes=10 --scale rootnode=1