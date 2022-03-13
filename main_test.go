package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var scenarios = map[int]State{
	0: {
		known: map[int]byte{
			1: 'o',
			3: 'e',
		},
		misplaced: map[intChar]bool{
			{1, 'e'}: true,
			{2, 'o'}: true,
			{2, 'w'}: true,
			{3, 'n'}: true,
		},
		wrong: map[byte]bool{},
	},
}

func TestSolution(t *testing.T) {
	s := scenarios[0]
	r := Solution(s)
	assert.Greater(t, len(r), 0)
	log.Println("Matches:", len(r))
	for _, w := range r {
		log.Println(w)
	}
}

func TestBuildDict(t *testing.T) {
	dict, err := buildDict("dictionary.csv")
	assert.NoError(t, err)
	assert.True(t, len(dict) > 0)
	log.Println("Words:", len(dict))
}

func TestQuordle(t *testing.T) {
	filename := "test_quordle.html"
	f, err := os.Open(filename)
	assert.NoError(t, err)
	states := parseQuordle(f)
	results := QuordleSolutions(states)
	for i, result := range results {
		fmt.Println("Results for", i)
		for _, r := range result {
			fmt.Println("  ", r)
		}
	}
}
