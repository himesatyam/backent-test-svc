FROM golang:alpine

WORKDIR /app

COPY . .


RUN go mod download

EXPOSE 8080

CMD go run ./src/main.go