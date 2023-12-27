From golang:1.18

WORKDIR /go/src/app

COPY . .
EXPOSE $GO_DOCKER_PORT
RUN go build -o main main.go

CMD ["./main"]
