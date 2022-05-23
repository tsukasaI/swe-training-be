package main

import (
	"forbizbe/src/database"

	"net/http"
	"net/http/httptest"
	"testing"
)

// test Gin
func TestRequests(t *testing.T) {
	router := setUpRouter()

	w := httptest.NewRecorder()
	homeInvalidReq, _ := http.NewRequest("GET", "/home", nil)
	router.ServeHTTP(w, homeInvalidReq)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, w.Code)
	}

	homeValidReq, _ := http.NewRequest("GET", "/home?userId=1", nil)
	router.ServeHTTP(w, homeValidReq)
	w = httptest.NewRecorder()

	if w.Code != http.StatusOK {
		t.Errorf("expected '%d' but got '%d'", http.StatusOK, w.Code)
	}
}

// ホントはmock入れたいけど時間の都合上諦め
// 30 minくらいやって"github.com/DATA-DOG/go-sqlmock"使うと良さそうなのだけわかった
func TestFindUser(t *testing.T) {
	db, err := database.ConnectDb()
	if err != nil {
		t.Errorf("Error Opening DB %s", err)
	}
	_, err = findUser(db, "1")
	if err != nil {
		t.Errorf("Error Fetching User %s", err)
	}

	_, err = findUser(db, "100000")
	if err == nil {
		t.Error("Expected error when fetching not exist user but not caught one")
	}
}

func TestFindPost(t *testing.T) {
	db, err := database.ConnectDb()
	if err != nil {
		t.Errorf("Error Opening DB %s", err)
	}
	user, err := findUser(db, "1")
	if err != nil {
		t.Errorf("Error Fetching User %s", err)
	}

	_, err = getHomePosts(db, user)
	if err != nil {
		t.Errorf("Error fetching posts %s", err)
	}
}

func TestGetFollowIds(t *testing.T) {
	db, err := database.ConnectDb()
	if err != nil {
		t.Errorf("Error Opening DB %s", err)
	}
	user, err := findUser(db, "1")
	if err != nil {
		t.Errorf("Error Fetching User %s", err)
	}

	result := getFollowIds(user)
	if len(result) == 0 {
		t.Errorf("expected not 0 follows, but got 0")
	}
}
