package database

import (
	"testing"
)

func TestConnectDb(t *testing.T) {
	if _, err := ConnectDb(); err != nil {
		t.Errorf("Error opening database")
	}
}
