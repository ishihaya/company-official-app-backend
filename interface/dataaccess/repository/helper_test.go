package repository

import (
	"fmt"
	"testing"

	"github.com/ishihaya/company-official-app-backend/pkg/db"
)

func CleanUpRepositoryTest(tb testing.TB, conn *db.Conn, tables []string) {
	tb.Helper()
	conn.MustExec("SET FOREIGN_KEY_CHECKS = 0")
	for _, table := range tables {
		// テーブル名にプレースホルダは使えないのでSprintfで整形
		conn.MustExec(fmt.Sprintf("TRUNCATE table %s", table))
	}
	conn.MustExec("SET FOREIGN_KEY_CHECKS = 1")

	defer conn.Close()
}
