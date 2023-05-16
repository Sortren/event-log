package config

import (
	"fmt"
	"os"
)

type Postgres struct {
	postgresDb       string
	postgresUser     string
	postgresPassword string
	postgresPort     string
	postgresHost     string
	sslMode          string
	debugMode        string
}

func NewPostgres() Postgres {
	return Postgres{
		postgresDb:       os.Getenv("POSTGRES_DB"),
		postgresUser:     os.Getenv("POSTGRES_USER"),
		postgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		postgresPort:     os.Getenv("POSTGRES_PORT"),
		postgresHost:     os.Getenv("POSTGRES_HOST"),
		sslMode:          os.Getenv("SSL_MODE"),
		debugMode:        os.Getenv("DEBUG_MODE"),
	}
}

func (p Postgres) Dsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		p.postgresUser,
		p.postgresPassword,
		p.postgresHost,
		p.postgresPort,
		p.postgresDb,
		p.sslMode,
	)
}

func (p Postgres) DebugMode() bool {
	if p.debugMode == "1" {
		return true
	}
	return false
}
