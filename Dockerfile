FROM golang:latest
LABEL maintainer = "G4"
EXPOSE 5000
ADD src /go/src/
#RUN go get github.com/golang/protobuf/proto
#RUN go install src-code
#RUN go install src-code/proto
#WORKDIR /go/src/
#RUN go build -o main .
#CMD ["/go/src/main", "1"]