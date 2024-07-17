FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

COPY . ./

RUN go build -o main

EXPOSE 9090

CMD ["./main"]