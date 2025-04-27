package db

import (
    "context"
    "log"

    "github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis(addr, password string, db int) {
    client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       db,
    })

    // Criar um contexto manualmente
    ctx := context.Background()

    // Testar a conexão usando o contexto criado
    _, err := client.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }

    log.Println("Connected to Redis")
    RedisClient = client
}
