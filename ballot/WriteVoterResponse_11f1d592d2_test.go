package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Status struct {
	Message string `json:"message"`
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(status)
	if err != nil {
		log.Println("error marshaling response to vote request. error: ", err)
	}
	w.Write(resp)
}

func TestWriteVoterResponse_Success(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	status := Status{Message: "Success"}
	writeVoterResponse(recorder, status)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	expectedBody := `{"message":"Success"}`
	if recorder.Body.String() != expectedBody {
		t.Errorf("Expected response body %s, but got %s", expectedBody, recorder.Body.String())
	}
}

func TestWriteVoterResponse_Error(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	status := Status{Message: "Error"}
	writeVoterResponse(recorder, status)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	expectedBody := `{"message":"Error"}`
	if recorder.Body.String() != expectedBody {
		t.Errorf("Expected response body %s, but got %s", expectedBody, recorder.Body.String())
	}
}
