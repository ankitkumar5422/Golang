package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandler(t *testing.T) {
	expected := "Hello Ani"
    req := httptest.NewRequest(http.MethodGet, "/Ankit?name=Ani", nil)
	w := httptest.NewRecorder()
	RequestHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expected {
		t.Errorf("Expected %s but got %s", expected, string(data))
	}
}

func TestRequestHandlerBadRequest(t *testing.T) {
	expected := "Serve is pass"
	req := httptest.NewRequest(http.MethodGet, "/Ankit", nil)
	w := httptest.NewRecorder()
	RequestHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if string(data) != expected {
		t.Errorf("Expected %s but got %s", expected, string(data))
	}
}

func TestRequestHandlerInvalidRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/Ankit", nil)
	w := httptest.NewRecorder()
	RequestHandler(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d but got %d", http.StatusBadRequest, w.Code)
	}
}
