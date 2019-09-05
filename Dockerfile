FROM larjim/kademlialab:latest
LABEL maintainer = "G4"
WORKDIR /app
EXPOSE 5000
COPY . .

#CMD ["./network.go"]