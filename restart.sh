#!/bin/bash

#Hur man kor:
docker swarm leave -f
docker swarm init
#docker build -t test . #borde kunna byggas automatiskt i compose
docker stack deploy up -c docker-compose.yml

# Automatiskt?
#docker exec -it "container id" /bin/bash
#borde inte behovas om automatiskt
#app container 1> go run CreateRoot.go
#Borde kunna bli entrypoint/cmd i dockerfile
#app container 2> go run main.go
# -||-

# Dockerfile:
#CMD ["go", "run", "./main.go"]

