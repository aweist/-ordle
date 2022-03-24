package parse

import (
	"io"
	"log"

	"github.com/aweist/ordle/html"
)

func ParseWordle(r io.Reader) []State {
	states := []State{}

	root, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	boards := GetQuordleGameBoards(root)

	for i := range boards {
		board := boards[i]
		state := NewState()
		rows := html.FindNodeByAttr(board, "role", "row")
		for j := range rows {
			row := rows[j]
			cells := html.FindNodeByAttr(row, "role", "cell")
			for k := range cells {
				cell := cells[k]
				content := html.FindNodeByAttr(cell, "class", "quordle-box-content")
				char := html.NodeValue(content[0])
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
