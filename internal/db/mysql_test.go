package db

import (
	"testing"
)

func TestConnectDatabase(t *testing.T) {
	err := InitDatabase("gopher:Gopher@tcp(localhost:3306)/golang")
	if err != nil {
		return
	}
	if DB == nil {
		t.Fatal("Expected db to be initialized, but it is nil")
	}
}
