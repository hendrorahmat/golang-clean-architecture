FROM golang:1.19.1 AS base
FROM base as dev
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
WORKDIR /opt/app/api

CMD ["air", "-c", ".air.http.toml"]

FROM base as built

WORKDIR /app/api
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN go get -d -v ./...
RUN go build -o api-server ./*.go

FROM busybox
WORKDIR /app
EXPOSE ${HTTP_PORT}
COPY --from=built /app/api/ /app/
CMD ["./api-server", "http"]