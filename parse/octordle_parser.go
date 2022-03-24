package parse

import (
	"io"
	"log"
	"strings"

	"github.com/aweist/ordle/html"
)

const (
	Black  string = "rgb(24,26,27)"
	Green  string = "rgb(0,204,136)"
	Yellow string = "rgb(255,204,0)"
)

func ParseOctordle(r io.Reader) []State {
	states := []State{}

	root, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	boards := GetOctordleGameBoards(root)

	for _, board := range boards {
		state := NewState()
		cells := html.FindNodeByAttr(board, "class", "box button")
		for i, c := range cells {
			index := i % 5
			char := html.NodeValue(c)
			if char != ' ' {
				cellResult := readOctordleCell(c)
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

func GetOctordleGameBoards(root *html.Node) []*html.Node {
	game := html.FindNodeByAttr(root, "id", "game")
	div := html.FindNodeByAttr(game[0], "id", "normal-container")
	return html.FindNodeByAttr(div[0], "class", "table_guesses")
}

func readOctordleCell(node *html.Node) CellResult {
	style := html.GetAttr(node, "style")
	styleMap := styleMap(style)
	bg := styleMap["background-color"]
	switch bg {
	case Black:
		return Wrong
	case Green:
		return Correct
	case Yellow:
		return Misplaced
	default:
		return Empty
	}
}

func styleMap(style string) map[string]string {
	style = strings.ReplaceAll(style, " ", "")
	m := map[string]string{}
	pairs := strings.Split(style, ";")
	for _, p := range pairs {
		if len(p) > 0 {
			kv := strings.Split(p, ":")
			if len(kv) == 2 {
				m[kv[0]] = kv[1]
			}
		}
	}
	return m
}
