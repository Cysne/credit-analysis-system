package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
    app := fiber.New()

    app.Use("/clientes", proxy.Forward("http://clientes:8082"))
    app.Use("/score", proxy.Forward("http://score:8084"))
    app.Use("/analise", proxy.Forward("http://analise:8081"))
    app.Use("/recomendacao", proxy.Forward("http://recomendacao:8083"))

    app.Listen(":8080")
}
