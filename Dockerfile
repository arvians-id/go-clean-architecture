FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go build -o clean-architecture

EXPOSE 8080

CMD ./clean-architecture