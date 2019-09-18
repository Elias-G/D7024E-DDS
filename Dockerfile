FROM larjim/kademlialab:latest
LABEL maintainer = "G4"
WORKDIR /app
EXPOSE 5000
RUN apt-get install net-tools
COPY . .
#CMD ["/bin/bash/go", "run", "./main.go", "1"]
#ENTRYPOINT ["/bin/bash/go", "run", "./main.go", "1"]
