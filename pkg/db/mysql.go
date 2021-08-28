package db

import (
	_ "github.com/go-sql-driver/mysql" // to connect mysql db
	"github.com/ishihaya/company-official-app-backend/pkg/env"
	"github.com/jmoiron/sqlx"
)

type Conn struct {
	*sqlx.DB
}

func New() *Conn {
	dsn := env.DSN()
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	conn := &Conn{db}
	return conn
}
