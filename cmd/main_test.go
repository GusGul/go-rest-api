package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func TestServerStartsSuccessfully(t *testing.T) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}

func TestServerFailsToStart(t *testing.T) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	err := http.ListenAndServe("invalid:address", router)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
