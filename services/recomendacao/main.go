package main

import (
    "log"

    "credit-analysis-system/internal/db"
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Connect to databases
    db.ConnectMongoDB("mongodb://mongodb:27017")
    db.ConnectRedis("redis:6379", "", 0)
    db.ConnectPostgres("postgres://user:password@postgres:5432/credit_analysis?sslmode=disable")

    app := fiber.New()

    app.Get("/health", func(c *fiber.Ctx) error {
        return c.SendString("service is healthy")
    })

    log.Fatal(app.Listen(":8083"))
}