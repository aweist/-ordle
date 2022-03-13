package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestGetStyle(t *testing.T) {
	s := input
	r := strings.NewReader(s)
	node, err := html.Parse(r)
	assert.NoError(t, err)
	buf := new(bytes.Buffer)
	err = html.Render(buf, node)
	assert.NoError(t, err)
	// fmt.Println(buf.String())
	body := node.FirstChild.LastChild
	table := children(body, "table")[0]
	tbody := children(table, "tbody")[0]
	row := children(tbody, "tr")[0]
	cell := children(row, "td")[0]
	// fmt.Println("Cell:", cell)
	inner := cell.FirstChild
	assert.Equal(t, "CELL CONTENTS", inner.Data)
}

func Test_bgFromStyle(t *testing.T) {
	tests := []struct {
		name  string
		style string
		want  string
	}{
		{
			name:  "black",
			style: "color: white; background-color: rgb(24, 26, 27);",
			want:  Black,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bgFromStyle(tt.style); got != tt.want {
				t.Errorf("bgFromStyle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoo(t *testing.T) {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	err = html.Render(buf, doc)
	assert.NoError(t, err)
	fmt.Println(buf.String())
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

var input string = `
<table><tbody>
	<tr>
		<td class="box button" width="20%" id="box3,1,1" style="color: white; background-color: rgb(24, 26, 27);">CELL CONTENTS</td>
	</tr>
</tbody></table>`

func TestParse(t *testing.T) {
	s := input
	// s := `<td class="box button" width="20%" id="box3,1,1" style="color: white; background-color: rgb(24, 26, 27);">T</td>`
	// s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	r := strings.NewReader(s)
	node, err := html.Parse(r)
	assert.NoError(t, err)
	node = node.FirstChild.LastChild.FirstChild.FirstChild
	fmt.Println("Node:", node.Type)
	buf := new(bytes.Buffer)
	err = html.Render(buf, node)
	assert.NoError(t, err)
	fmt.Println(buf.String())
}

func TestParseDoc(t *testing.T) {
	r := strings.NewReader(testTable)
	root, err := html.Parse(r)
	assert.NoError(t, err)
	data := parseDoc(root)
	assert.NotEmpty(t, data)
	assert.Equal(t, 5, len(data[0]))
	expected := []cellData{
		{
			Letter: 'T',
			Color:  "rgb(24,26,27)",
		},
		{
			Letter: 'E',
			Color:  "rgb(255,204,0)",
		},
		{
			Letter: 'A',
			Color:  "rgb(24,26,27)",
		},
		{
			Letter: 'R',
			Color:  "rgb(24,26,27)",
		},
		{
			Letter: 'S',
			Color:  "rgb(24,26,27)",
		},
	}
	assert.Equal(t, data[0], expected)
}

func TestParseInput(t *testing.T) {
	known, misplaced, wrong := parseInput("test_input.html")
	assert.Equal(t, 1, len(known))
	assert.Equal(t, 1, len(misplaced))
	assert.Equal(t, 3, len(wrong))
}
