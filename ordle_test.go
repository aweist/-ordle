package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/aweist/ordle/parse"
	"github.com/stretchr/testify/assert"
)

func generateTestState() parse.State {
	state := parse.NewState()
	state.Known(1, 'o')
	state.Known(3, 'e')

	state.Misplaced(1, 'e')
	state.Misplaced(2, 'o')
	state.Misplaced(2, 'w')
	state.Misplaced(3, 'n')
	return state
}

func TestSolution(t *testing.T) {

	s := generateTestState()
	r := Solution(s)
	assert.Greater(t, len(r), 0)
	log.Println("Matches:", len(r))
	for _, w := range r {
		log.Println(w)
	}
}

func TestQuordle(t *testing.T) {
	filename := "parse/test_quordle.html"
	f, err := os.Open(filename)
	assert.NoError(t, err)
	states := parse.ParseQuordle(f)
	results := QuordleSolutions(states)
	for i, result := range results {
		fmt.Println("Results for", i)
		for _, r := range result {
			fmt.Println("  ", r)
		}
	}
}
