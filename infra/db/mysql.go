package db

import (
	"log"

	// to connect mysql db
	_ "github.com/go-sql-driver/mysql"
	"github.com/ishihaya/company-official-app-backend/config"
	"github.com/jmoiron/sqlx"
)

type Conn struct {
	*sqlx.DB
}

func New() *Conn {
	dsn := config.DSN()
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %+v", err)
	}
	conn := &Conn{db}
	return conn
}
