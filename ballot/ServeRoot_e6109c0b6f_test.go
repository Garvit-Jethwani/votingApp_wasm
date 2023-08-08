package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
)

type Response struct {
	Results    []CandidateVotes `json:"results"`
	TotalVotes int              `json:"total_votes"`
}

type CandidateVotes struct {
	CandidateID string `json:"candidate_id"`
	Votes       int    `json:"votes"`
}

type Vote struct {
	VoterID     string `json:"voter_id"`
	CandidateID string `json:"candidate_id"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodGet:
		defer r.Body.Close()
		log.Println("result request received")
		res := Response{}
		votes := getCandidatesVote()
		for candidateID, votes := range votes {
			res.Results = append(res.Results, CandidateVotes{candidateID, votes})
			res.TotalVotes += votes
		}

		sort.Slice(res.Results, func(i, j int) bool {
			return res.Results[i].Votes > res.Results[j].Votes
		})

		log.Printf("result data: %+v", res)

		out, err := json.Marshal(res)
		if err != nil {
			log.Println("error marshaling response to result request. error: ", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)

	case http.MethodPost:
		log.Println("vote received")
		vote := Vote{}
		status := Status{}
		defer r.Body.Close()

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&vote)
		if err != nil {
			log.Printf("error parsing vote data. error: %v\n", err)
			status.Code = http.StatusBadRequest
			status.Message = "Vote is not valid. Vote can not be saved"
			writeVoterResponse(w, status)
			return
		}
		log.Printf("Voting done by voter: %s to candidate: %s\n", vote.VoterID, vote.CandidateID)
		err = saveVote(vote)
		if err != nil {
			log.Println(err)
			status.Code = http.StatusBadRequest
			status.Message = "Vote is not valid. Vote can not be saved"
			writeVoterResponse(w, status)
			return
		}
		status.Code = http.StatusCreated
		status.Message = "Vote saved sucessfully"
		writeVoterResponse(w, status)
		return

	default:
		status := Status{}
		status.Code = http.StatusMethodNotAllowed
		status.Message = "Bad Request. Vote can not be saved"
		writeVoterResponse(w, status)
		return
	}
}

func writeVoterResponse(w http.ResponseWriter, status Status) {
	response, err := json.Marshal(status)
	if err != nil {
		log.Println("error marshaling response to voter. error: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func getCandidatesVote() map[string]int {
	// TODO: Implement getCandidatesVote() function
	return nil
}

func saveVote(vote Vote) error {
	// TODO: Implement saveVote() function
	return nil
}

func TestServeRoot_e6109c0b6f(t *testing.T) {
	// Test case 1: GET request
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(serveRoot)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// TODO: Check the response body for correctness

	// Test case 2: POST request
	vote := Vote{
		VoterID:     "12345",
		CandidateID: "A",
	}

	voteJSON, err := json.Marshal(vote)
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(voteJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// TODO: Check the response body for correctness
}
