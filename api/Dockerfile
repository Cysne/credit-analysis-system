FROM golang:1.22-alpine

# 1. Primeira pasta de trabalho
WORKDIR /app

# 2. Copia tudo
COPY . .

# 3. Baixar dependências (o go.mod está na raiz)
RUN go mod download

# 4. Agora, muda pra pasta API
WORKDIR /app/api

# 5. Builda dentro da pasta correta
RUN go build -o main .

# 6. Rodar
CMD ["./main"]