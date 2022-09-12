FROM golang:1.19.1 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./clean-architecture

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["./clean-architecture", "http"]