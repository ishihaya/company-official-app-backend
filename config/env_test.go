package config

import (
	"fmt"
	"os"
	"testing"
)

func TestPORT(t *testing.T) {
	dbname := os.Getenv("MYSQL_DATABASE")
	fmt.Println("--------dbname-----------", dbname)
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PORT(); got != tt.want {
				t.Errorf("PORT() = %v, want %v", got, tt.want)
			}
		})
	}
}
