FROM golang:alpine

WORKDIR /app

COPY . .


RUN go mod download

EXPOSE 80

CMD go run ./src/main.go