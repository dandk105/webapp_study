package main

import (
	"log"
	"net/http/httptest"
	"testing"
)

func TestGetUserslistHandler(t *testing.T) {
	// serverが動いていないので、リクエストを送っても当然空白のレスポンスが返ってくるだけ
	req := httptest.NewRequest("GET", "/api/users", nil)
	w := httptest.NewRecorder()

	log.Print(req)
	log.Print(w)

}
