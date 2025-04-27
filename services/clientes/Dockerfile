FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/services/clientes

RUN go build -o main .

CMD ["./main"]
