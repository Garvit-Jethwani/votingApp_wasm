package main

import (
	"testing"
)

type Vote struct {
	CandidateID int
}

func saveVote(vote Vote) error {
	// Implementation of saveVote function
	return nil
}

func TestSaveVote_Success(t *testing.T) {
	vote := Vote{
		CandidateID: 1,
	}

	err := saveVote(vote)

	if err != nil {
		t.Error("Expected no error, got", err)
	}
}

func TestSaveVote_CandidateNotFound(t *testing.T) {
	vote := Vote{
		CandidateID: 10,
	}

	err := saveVote(vote)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}
