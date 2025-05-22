# chat-server Dockerfile

FROM golang:1.24-alpine3.21

WORKDIR /app

COPY go.mod ./

RUN go mod download 

COPY . .

RUN go build -o chat-server

EXPOSE 9000

CMD ["./chat-server"]