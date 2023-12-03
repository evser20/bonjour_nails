package pg

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Address  string
	Port     uint
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type DB struct {
	*sql.DB
}

func NewDB(c DBConfig) (*DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Address, c.Port, c.User, c.Password, c.DBName, "disable",
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	pgDB := DB{DB: db}

	return &pgDB, nil
}
