#!/bin/bash
docker-compose down
docker build -t rootnode .
docker build -t kademlianodes ./src
docker-compose up --scale kademliaNodes=1 --scale rootnode=1