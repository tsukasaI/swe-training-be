package models

import (
	"forbizbe/src/database"
	"testing"
)

// ホントはmock入れたいけど時間の都合上諦め
// 30 minくらいやって"github.com/DATA-DOG/go-sqlmock"使うと良さそうなのだけわかった
func TestFindUser(t *testing.T) {
	db, err := database.ConnectDb()
	if err != nil {
		t.Errorf("Error Opening DB %s", err)
	}
	_, err = FindUser(db, "1")

	if err != nil {
		t.Errorf("Error Fetching User %s", err)
	}

	_, err = FindUser(db, "100000")
	if err == nil {
		t.Error("Expected error when fetching not exist user but not caught one")
	}

	_, err = FindUser(db, "abcde")
	if err == nil {
		t.Error("Expected error when fetching not exist user but not caught one")
	}
}

func TestGetFollowIds(t *testing.T) {
	db, err := database.ConnectDb()
	if err != nil {
		t.Errorf("Error Opening DB %s", err)
	}
	user, err := FindUser(db, "1")
	if err != nil {
		t.Errorf("Error Fetching User %s", err)
	}

	result := getFollowIds(user)
	if len(result) == 0 {
		t.Errorf("expected not 0 follows, but got 0")
	}
}
