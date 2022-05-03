FROM golang:1.18-alpine as prod

WORKDIR /app/car-sales-api

RUN apk add --no-cache git build-base

COPY . .

## download deps and build golang bin
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o bin/api main.go

## run binary
CMD ["./bin/api"]
