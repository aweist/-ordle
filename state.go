package main

import "fmt"

type State struct {
	known     map[int]byte
	misplaced map[intChar]bool
	wrong     map[byte]bool
}

type intChar struct {
	index int
	char  byte
}

func NewState() State {
	return State{
		known:     map[int]byte{},
		misplaced: map[intChar]bool{},
		wrong:     map[byte]bool{},
	}
}

func (s *State) Known(index int, char byte) {
	s.known[index] = char
}

func (s *State) IsKnown(index int) (byte, bool) {
	char, ok := s.known[index]
	return char, ok
}

func (s *State) Misplaced(index int, char byte) {
	s.misplaced[intChar{index, char}] = true
}

func (s *State) IsMisplaced(index int, char byte) bool {
	return s.misplaced[intChar{index, char}]
}

func (s *State) Wrong(char byte) {
	s.wrong[char] = true
}

func (s *State) IsWrong(char byte) bool {
	return s.wrong[char]
}

func (s *State) AllMisplaced() map[byte]bool {
	allMisplaced := map[byte]bool{}
	for intChar := range s.misplaced {
		allMisplaced[intChar.char] = true
	}
	return allMisplaced
}

func (s *State) Print() {
	fmt.Println("known:")
	for i, c := range s.known {
		fmt.Println("  ", i, string(c))
	}
	fmt.Println("misplaced:")
	for k := range s.misplaced {
		fmt.Println("  ", k.index, string(k.char))
	}
	fmt.Println("wrong:")
	for c := range s.wrong {
		fmt.Println("  ", string(c))
	}

}
