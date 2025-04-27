package db

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

var PostgresDB *sql.DB

func ConnectPostgres(dsn string) {
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to PostgreSQL: %v", err)
    }

    // Test connection
    if err := db.Ping(); err != nil {
        log.Fatalf("Failed to ping PostgreSQL: %v", err)
    }

    log.Println("Connected to PostgreSQL")
    PostgresDB = db
}