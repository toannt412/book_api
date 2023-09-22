package test_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAdmin(t *testing.T) {
	var jsonStr = []byte(`{"username":"viph","password":"123456","email":"viph@gmail.com","role":"admin"}`)

	req, err := http.NewRequest("POST", "/admin", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"username":"viph","password":"123456","email":"viph@gmail.com","role":"admin""}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetAAdmin(t *testing.T) {

}

func TestEditAAdmin(t *testing.T) {

}

func TestDeleteAAdmin(t *testing.T) {

}

func TestLoginAccountAdmin(t *testing.T) {

}
