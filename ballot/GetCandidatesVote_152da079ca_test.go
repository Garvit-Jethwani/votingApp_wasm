package main

import (
	"sync"
	"testing"
)

var (
	once                sync.Once
	candidateVotesStore map[string]int
)

func getCandidatesVote() map[string]int {
	once.Do(func() {
		candidateVotesStore = make(map[string]int)
	})
	return candidateVotesStore
}

func TestGetCandidatesVote_Success(t *testing.T) {
	votes := getCandidatesVote()

	if len(votes) != 0 {
		t.Error("Expected an empty map, but got a non-empty map")
	}
}

func TestGetCandidatesVote_MultipleCalls(t *testing.T) {
	votes1 := getCandidatesVote()
	votes2 := getCandidatesVote()

	if &votes1 != &votes2 {
		t.Error("Expected both calls to getCandidatesVote to return the same map, but got different maps")
	}
}
