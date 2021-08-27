package db

import (
	"sync"

	// to connect mysql db
	_ "github.com/go-sql-driver/mysql"
	"github.com/ishihaya/company-official-app-backend/pkg/env"
	"github.com/jmoiron/sqlx"
)

type Conn struct {
	*sqlx.DB
}

var sharedInstance *Conn
var once sync.Once

func GetInstance() *Conn {
	once.Do(func() {
		sharedInstance = newInstance()
	})
	return sharedInstance
}

func newInstance() *Conn {
	dsn := env.DSN()
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	conn := &Conn{db}
	return conn
}
