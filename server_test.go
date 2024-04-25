package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddData(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(addData))
	defer server.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", server.URL, bytes.NewBuffer([]byte(`{"id":"123", "ml":"100", "waterType":"still"}`)))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", resp.Status)
	}
}

func TestHandleRequest(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRequest)
	mux.HandleFunc("/add", addData)

	server := httptest.NewServer(mux)
	defer server.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", server.URL+"/add", bytes.NewBuffer([]byte(`{"id":"123", "ml":"100", "waterType":"still"}`)))
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("POST", server.URL, bytes.NewBuffer([]byte(`{"id":"123"}`)))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", resp.Status)
	}
}

func TestAddDataManually(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(addDataManually))
	defer server.Close()

	resp, err := http.Get(server.URL + "?id=456&ml=200&waterType=sprudel")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", resp.Status)
	}
}
