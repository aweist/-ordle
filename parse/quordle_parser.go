package parse

import (
	"io"
	"log"
	"strings"

	"github.com/aweist/ordle/html"
	net_html "golang.org/x/net/html"
)

func ParseQuordle(r io.Reader) []State {
	states := []State{}

	root, err := net_html.Parse(r)
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

func GetQuordleGameBoards(root *net_html.Node) []*net_html.Node {
	boards := html.FindNodeByAttr(root, "aria-label", "Game Boards")
	return html.FindNodeByAttr(boards[0], "role", "table")
}

func readCell(cell *net_html.Node) CellResult {
	class := html.GetAttr(cell, "class")
	if strings.Contains(class, "bg-box-correct") {
		return Correct
	}
	if strings.Contains(class, "bg-box-diff") {
		return Misplaced
	}
	return Wrong
}
