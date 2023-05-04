FROM golang:1.12.5-alpine

RUN mkdir -p /go/src/provisioned-api

WORKDIR /go/src/provisioned-api
COPY go.mod  ./
RUN go mod download && go mod verify
COPY . /go/src/provisioned-api
RUN go install  provisioned-api


CMD ["/go/bin/provisioned-api"]
EXPOSE 8080