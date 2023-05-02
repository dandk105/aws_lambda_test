FROM golang:1.20.0-buster AS builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go ./
RUN go build -v -o /usr/bin/app

CMD [ "/usr/bin/app" ]