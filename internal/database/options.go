package database

import (
	"os"
)

type PostgreSQLOptions struct {
	dsn string
}

func NewPostgreSQLOptions(dsn string) *PostgreSQLOptions {
	return &PostgreSQLOptions{
		dsn: dsn,
	}
}

func DefaultPostgreSQLOptions() *PostgreSQLOptions {
	// Obter a string de conexão do ambiente
	dsn := os.Getenv("POSTGRESQL_CONNECTION_STRING")
	return NewPostgreSQLOptions(dsn)
}
