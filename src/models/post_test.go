package models

import (
	"forbizbe/src/database"
	"testing"
)

func TestFindPost(t *testing.T) {
	db, err := database.ConnectDb()
	if err != nil {
		t.Errorf("Error Opening DB %s", err)
	}
	user, err := FindUser(db, "1")
	if err != nil {
		t.Errorf("Error Fetching User %s", err)
	}

	_, err = GetHomePosts(db, user)
	if err != nil {
		t.Errorf("Error fetching posts %s", err)
	}
}
