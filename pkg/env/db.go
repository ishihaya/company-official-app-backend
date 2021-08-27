package env

import (
	"fmt"
	"os"
)

func DSN() string {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	host := os.Getenv("DB_HOST")
	if host != "" {
		// tcp
		port := os.Getenv("DB_PORT")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", user, password, host, port, dbname)
		return dsn
	}

	// unix
	dbSocketPath := os.Getenv("DB_SOCKET_PATH")
	dsn := fmt.Sprintf("%s:%s@unix(%s)/%s?parseTime=true", user, password, dbSocketPath, dbname)
	return dsn
}
