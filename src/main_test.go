package main

import (
	"bytes"
	"encoding/json"
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
		const postComment = "Good morning and have a nice day."
		data := requestPost{postComment}
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

		router.ServeHTTP(w, postValidReq)
		if w.Code != http.StatusOK {
			t.Errorf("expected '%d' but got '%d'", http.StatusOK, w.Code)
		}
		// ここ本当はresponse bodyの中身チェックしたいけど、すぐできないから一旦諦め

		// set invalid query param
		postInvalidQueryReq, _ := http.NewRequest(http.MethodPost, "/post", bytes.NewBuffer([]byte(payload)))
		w = httptest.NewRecorder()
		q = postInvalidQueryReq.URL.Query()
		q.Del("userId")
		q.Add("userId", "")
		postInvalidQueryReq.URL.RawQuery = q.Encode()

		router.ServeHTTP(w, postInvalidQueryReq)
		if w.Code != http.StatusNotFound {
			t.Errorf("expected '%d' but got '%d'", http.StatusNotFound, w.Code)
		}

		// without payload
		var invalidRequestPost requestPost
		invalidPayload, err := json.Marshal(invalidRequestPost)
		if err != nil {
			t.Error("Error occurreded when invalid payload preparation")
		}

		postInvalidPayloadReq, _ := http.NewRequest(http.MethodPost, "/post", bytes.NewBuffer([]byte(invalidPayload)))
		w = httptest.NewRecorder()
		q = postInvalidPayloadReq.URL.Query()
		q.Del("userId")
		q.Add("userId", "1")
		postInvalidPayloadReq.URL.RawQuery = q.Encode()

		router.ServeHTTP(w, postInvalidPayloadReq)
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, w.Code)
		}

		// with over length payload
		var invalidOverLengthCommentRequestPost = requestPost{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		}
		invalidOverLengthCommentPayload, err := json.Marshal(invalidOverLengthCommentRequestPost)
		if err != nil {
			t.Error("Error occurreded when invalid payload preparation")
		}

		postInvalidOverLengthCommentPayloadReq, _ := http.NewRequest(http.MethodPost, "/post", bytes.NewBuffer([]byte(invalidOverLengthCommentPayload)))
		w = httptest.NewRecorder()
		q = postInvalidOverLengthCommentPayloadReq.URL.Query()
		q.Del("userId")
		q.Add("userId", "1")
		postInvalidOverLengthCommentPayloadReq.URL.RawQuery = q.Encode()

		router.ServeHTTP(w, postInvalidOverLengthCommentPayloadReq)
		if w.Code != http.StatusBadRequest {
			t.Errorf("expected '%d' but got '%d'", http.StatusBadRequest, w.Code)
		}

		// with not exist userId
		NotExistUserReq, _ := http.NewRequest(http.MethodPost, "/post", bytes.NewBuffer([]byte(payload)))
		w = httptest.NewRecorder()
		q = NotExistUserReq.URL.Query()
		q.Del("userId")
		q.Add("userId", "1000000")
		NotExistUserReq.URL.RawQuery = q.Encode()

		router.ServeHTTP(w, NotExistUserReq)
		if w.Code != http.StatusNotFound {
			t.Errorf("expected '%d' but got '%d'", http.StatusNotFound, w.Code)
		}
	})
}
