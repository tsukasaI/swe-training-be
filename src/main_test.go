package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type requestPost struct {
	Comment string `json:"comment"`
}

// test Gin
func TestRequests(t *testing.T) {
	router := setUpRouter()

	t.Run("Execute home get request.", func(t *testing.T) {
		w := httptest.NewRecorder()
		// invalid request
		homeInvalidReq, _ := http.NewRequest(http.MethodGet, "/post/home", nil)
		router.ServeHTTP(w, homeInvalidReq)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected '%d' but got '%d'", http.StatusNotFound, w.Code)
		}

		// valid request
		homeValidReq, _ := http.NewRequest(http.MethodGet, "/post/home?userId=1", nil)
		w = httptest.NewRecorder()
		router.ServeHTTP(w, homeValidReq)

		if w.Code != http.StatusOK {
			t.Errorf("expected '%d' but got '%d'", http.StatusOK, w.Code)
		}
	})

	t.Run("Execute post post request", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := requestPost{
			Comment: "Good morning and have a nice day.",
		}
		// set payload
		payload, err := json.Marshal(data)
		if err != nil {
			t.Errorf("Error occurred when create post body: %q", err.Error())
		}
		postValidReq, err := http.NewRequest(http.MethodPost, "/post", bytes.NewBuffer([]byte(payload)))
		if err != nil {
			t.Errorf("Error occurred when execute post request: %q", err.Error())
		}

		// set valid query param
		q := postValidReq.URL.Query()
		q.Add("userId", "1")
		postValidReq.URL.RawQuery = q.Encode()
		fmt.Printf("\n\n%v\n\n", postValidReq)

		router.ServeHTTP(w, postValidReq)
		if w.Code != http.StatusOK {
			t.Errorf("expected '%d' but got '%d'", http.StatusOK, w.Code)
		}

		// set valid query param
		q.Del("userId")
		q.Add("userId", "")
		postValidReq.URL.RawQuery = q.Encode()
		fmt.Printf("\n\n%v\n\n", postValidReq)

		router.ServeHTTP(w, postValidReq)
		if w.Code != http.StatusNotFound {
			t.Errorf("expected '%d' but got '%d'", http.StatusNotFound, w.Code)
		}

	})
}
