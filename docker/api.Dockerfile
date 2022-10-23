FROM golang:1.19.1 AS base
FROM base as dev
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
WORKDIR /opt/app/api

CMD ["air", "-c", ".air.http.toml"]

FROM base as built

WORKDIR /go/app/api
COPY . .

ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN go build -o /tmp/api-server ./*.go

FROM busybox

COPY --from=built /tmp/api-server /usr/bin/api-server
CMD ["api-server", "start"]

#COPY ../go.mod ./
#COPY ../go.sum ./
#
#RUN go mod download
#
#COPY .. .
#
#RUN go build -o ./clean-architecture
#
#WORKDIR /app
#
#EXPOSE 8080
#
#ENTRYPOINT ["./clean-architecture", "http"]