package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// test Gin
func TestRoute(t *testing.T) {
	router := setUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/home", nil)
	router.ServeHTTP(w, req)

	expected := http.StatusBadRequest
	if w.Code != expected {
		t.Errorf("expected '%d' but got '%d'", expected, w.Code)
	}
}

// getFollowIds(user User)
// func testGetFollowIds() {}

// getHomeData(db *gorm.DB, user User)
