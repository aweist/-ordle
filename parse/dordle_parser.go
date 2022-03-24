package parse

import (
	"io"
	"log"

	"github.com/aweist/ordle/html"
)

const (
	DordleBlack  string = "var(--bgc)"
	DordleGreen  string = "var(--okc)"
	DordleYellow string = "" // yellow uses an image instead of color
)

func ParseDordle(r io.Reader) []State {
	states := []State{}

	root, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	boards := getDordleGameBoards(root)

	for _, board := range boards {
		state := NewState()
		cells := html.FindNodeByAttr(board, "class", "box button")
		for i, c := range cells {
			index := i % 5
			char := html.NodeValue(c)
			if char != ' ' {
				cellResult := readDordleCell(c)
				switch cellResult {
				case Correct:
					state.Known(index, char)
				case Misplaced:
					state.Misplaced(index, char)
				case Wrong:
					state.Wrong(char)
				default:
				}
			}
		}
		states = append(states, state)
	}

	return states
}

func readDordleCell(node *html.Node) CellResult {
	style := html.GetAttr(node, "style")
	styleMap := styleMap(style)
	bg := styleMap["background-color"]
	switch bg {
	case DordleBlack:
		return Wrong
	case DordleGreen:
		return Correct
	case DordleYellow:
		return Misplaced
	default:
		return Empty
	}
}

func getDordleGameBoards(root *html.Node) []*html.Node {
	game := html.FindNodeByAttr(root, "id", "game")
	boards := html.FindNodeByAttr(game[0], "class", "table_guesses")
	result := []*html.Node{}
	// Need to drop the keyboard from the nodes, as it also has the table_guesses class
	for _, b := range boards {
		id := html.GetAttr(b, "id")
		if id != "keyboard" {
			result = append(result, b)
		}
	}
	return result
}
