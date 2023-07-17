FROM golang:alpine3.18

WORKDIR /app

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 8000

