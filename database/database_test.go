package database

import (
	"testing"
)

func TestGetConn(t *testing.T) {
	db := GetConn()
	db.Close()
}
