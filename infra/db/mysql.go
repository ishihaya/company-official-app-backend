package db

import (
	"database/sql"
	"log"

	// to connect mysql db
	_ "github.com/go-sql-driver/mysql"
	"github.com/ishihaya/company-official-app-backend/config"
)

type Conn struct {
	*sql.DB
}

func New() *Conn {
	dsn := config.DSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %+v", err)
	}
	conn := &Conn{db}
	return conn
}
