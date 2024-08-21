package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
}

func NewDB() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Connect() (*gorm.DB, error) {
	// dsn := "postgres://zakiarsyad:postgres@localhost:5432/neurons" //local DB
	// dsn := "postgres://postgres:thMprixGIqxU01R@localhost:15432/back_end_wild_grass_8451" //Proxy to Fly.io
	dsn := os.Getenv("DATABASE_URI") //Fly.io

	dsn += "?client_encoding=UTF8"

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}
