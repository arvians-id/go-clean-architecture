FROM golang:1.20-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY /cmd/server/main.go .
COPY . .

EXPOSE 8080

RUN go mod tidy

RUN go build -o main

CMD ./main