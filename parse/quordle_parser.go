package parse

import (
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type CellResult string

const (
	Correct   CellResult = "correct"
	Misplaced CellResult = "Misplaced"
	Wrong     CellResult = "wrong"
	Empty     CellResult = ""
)

func ParseQuordle(r io.Reader) []State {
	states := []State{}

	root, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	boards := GetGameBoards(root)

	for i := range boards {
		board := boards[i]
		state := NewState()
		rows := FindNodeByAttr(board, "role", "row")
		for j := range rows {
			row := rows[j]
			cells := FindNodeByAttr(row, "role", "cell")
			for k := range cells {
				cell := cells[k]
				content := FindNodeByAttr(cell, "class", "quordle-box-content")
				char := nodeValue(content[0])
				cellResult := readCell(cell)
				if char != ' ' {
					switch cellResult {
					case Correct:
						state.Known(k, char)
					case Misplaced:
						state.Misplaced(k, char)
					case Wrong:
						state.Wrong(char)
					default:
					}
				}
			}
		}
		states = append(states, state)
	}

	return states
}

func GetGameBoards(root *html.Node) []*html.Node {
	boards := FindNodeByAttr(root, "aria-label", "Game Boards")
	return FindNodeByAttr(boards[0], "role", "table")
}

func readCell(cell *html.Node) CellResult {
	class := GetAttr(cell, "class")
	if strings.Contains(class, "bg-box-correct") {
		return Correct
	}
	if strings.Contains(class, "bg-box-diff") {
		return Misplaced
	}
	return Wrong
}
