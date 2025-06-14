FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
COPY .env .env

RUN go mod download

COPY . .

RUN go build -o main ./src/cmd

EXPOSE 8080

CMD ["./main"] 