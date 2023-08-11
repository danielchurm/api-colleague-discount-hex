package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	log "github.com/JSainsburyPLC/go-logrus-wrapper"

	. "github.com/JSainsburyPLC/smartshop-api-go-template/config"
)

func NewConnection(config DatabaseConfig) *sql.DB {
	connectionString := fmt.Sprintf("%s://%s:%s@%s:%d/%s?connect_timeout=%d&sslmode=%s", config.Type, config.User, config.Password, config.Host, config.Port, config.Name, config.TimeoutSeconds, config.SSLMode)

	db, err := sql.Open(config.Type, connectionString)
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	return db
}
