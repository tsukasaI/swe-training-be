package main

import (
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

// getFollowIds(user User)
// func testGetFollowIds() {}

// getHomeData(db *gorm.DB, user User)
