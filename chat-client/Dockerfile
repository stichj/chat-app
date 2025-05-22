# chat-client Dockerfile

FROM golang:1.24-alpine3.21

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o chat-client

CMD [ "./chat-client" ]