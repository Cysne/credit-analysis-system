package main

import (
    "log"
    "os"

    "credit-analysis-system/internal/db"
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Ler variáveis de ambiente
    mongoURI := os.Getenv("MONGO_URI")       // URI do MongoDB
    redisAddr := os.Getenv("REDIS_ADDR")     // Endereço do Redis
    postgresDSN := os.Getenv("POSTGRES_DSN") // DSN do PostgreSQL

    // Conectar aos bancos de dados
    db.ConnectMongoDB(mongoURI)
    db.ConnectRedis(redisAddr, "", 0)
    db.ConnectPostgres(postgresDSN)

    app := fiber.New()

    app.Get("/health", func(c *fiber.Ctx) error {
        return c.SendString("service is healthy")
    })

    log.Fatal(app.Listen(":8082"))
}