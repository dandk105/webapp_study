package main

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestCreateUserDataHandler(t *testing.T) {
	/* Create a new request with a JSON body
	jsonStr := []byte(`{"name":"John Doe","birth_day":"2000-01-01"}`)
	req, err := http.NewRequest("POST", "/api/createuserdata", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function with the recorder and request
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		createUserDataHandler(w, r, db)
	})
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := `User created successfully`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	*/
}
