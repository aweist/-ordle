package parse

import (
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func parseInput(filename string) (
	known map[int]byte,
	misplaced map[int]map[byte]bool,
	wrong map[byte]bool,
) {
	known = map[int]byte{}
	misplaced = map[int]map[byte]bool{}
	wrong = map[byte]bool{}

	r, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	node, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	data := parseDoc(node)
	for i := range data {
		for j := range data[i] {
			color, letter := data[i][j].Color, data[i][j].Letter
			switch color {
			case Black:
				wrong[letter] = true
			case Yellow:
				if misplaced[j] == nil {
					misplaced[j] = map[byte]bool{}
				}
				misplaced[j][letter] = true
			case Green:
				known[j] = letter
			}
		}
	}

	// Must remove all "known" from "wrong", since letters could appear more than once
	for _, c := range known {
		delete(wrong, c)
	}

	return
}

const (
	Black  string = "rgb(24,26,27)"
	Green  string = "rgb(0,204,136)"
	Yellow string = "rgb(255,204,0)"
)

func getBackground(cell *html.Node) string {
	attrs := cell.Attr
	for _, attr := range attrs {
		if attr.Key == "style" {
			style := attr.Val
			return bgFromStyle(style)
		}
	}
	return ""
}

func bgFromStyle(style string) string {
	style = strings.ReplaceAll(style, " ", "")
	tuples := strings.Split(style, ";")
	for _, t := range tuples {
		kv := strings.Split(t, ":")
		if len(kv) == 2 {
			k, v := kv[0], kv[1]
			if k == "background-color" {
				return v
			}
		}
	}
	return ""
}

type cellData struct {
	Letter byte
	Color  string
}

func parseDoc(root *html.Node) [][]cellData {
	result := [][]cellData{}
	// First child should be html
	html := children(root, "html")[0]

	// Go to body
	body := children(html, "body")[0]

	table := children(body, "table")[0]

	tbody := children(table, "tbody")[0]

	rows := children(tbody, "tr")

	for _, r := range rows {
		newRow := []cellData{}
		cells := children(r, "td")
		for _, c := range cells {
			if c.FirstChild != nil {
				cellData := cellData{
					Color:  getBackground(c),
					Letter: c.FirstChild.Data[0] + 'a' - 'A',
				}
				newRow = append(newRow, cellData)
			}
		}
		result = append(result, newRow)
	}

	return result
}

func children(node *html.Node, data string) []*html.Node {
	children := []*html.Node{}
	for node = node.FirstChild; node != nil; node = node.NextSibling {
		if node.Data == data {
			children = append(children, node)
		}
	}
	return children
}
