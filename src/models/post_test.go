package models

import (
	"forbizbe/src/database"
	"testing"
)

func TestFindPostForHome(t *testing.T) {
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

func TestCreatePost(t *testing.T) {
	db, err := database.ConnectDb()
	if err != nil {
		t.Errorf("Error Opening DB %s", err)
	}

	const insertComment = "test"
	post, err := CreatePost(db, "1", insertComment)
	if err != nil {
		t.Errorf("Error Creating post %q", err)
	}
	if post.Comment != insertComment {
		t.Errorf("Comment is not inserted collectly %q", err)
	}
}
