#FROM larjim/kademlialab:latest
FROM golang:onbuild
LABEL maintainer = "G4"
WORKDIR /app
EXPOSE 5000
RUN apt-get install net-tools
COPY . .
#RUN "go run main.go"