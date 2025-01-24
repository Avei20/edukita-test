FROM golang:1.22.0 AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

# RUN go install github.com/swaggo/swag/cmd/swag@latest

# RUN swag init

RUN go build -o main .

CMD ["./main"]
