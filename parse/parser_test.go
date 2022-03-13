package parse

import (
	"bytes"
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
	body := node.FirstChild.LastChild
	table := children(body, "table")[0]
	tbody := children(table, "tbody")[0]
	row := children(tbody, "tr")[0]
	cell := children(row, "td")[0]
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

var input string = `
<table><tbody>
	<tr>
		<td class="box button" width="20%" id="box3,1,1" style="color: white; background-color: rgb(24, 26, 27);">CELL CONTENTS</td>
	</tr>
</tbody></table>`

var testTable string = `
<table><tbody><tr>
            <td class="box button" width="20%" id="box3,1,1" style="color: white; background-color: rgb(24, 26, 27);">T</td>
            <td class="box button" width="20%" id="box3,1,2" style="color: black; background-color: rgb(255, 204, 0);">E</td>
            <td class="box button" width="20%" id="box3,1,3" style="color: white; background-color: rgb(24, 26, 27);">A</td>
            <td class="box button" width="20%" id="box3,1,4" style="color: white; background-color: rgb(24, 26, 27);">R</td>
            <td class="box button" width="20%" id="box3,1,5" style="color: white; background-color: rgb(24, 26, 27);">S</td>
        </tr>
</tbody></table>`

func TestParseDoc(t *testing.T) {
	r := strings.NewReader(testTable)
	root, err := html.Parse(r)
	assert.NoError(t, err)
	data := parseDoc(root)
	assert.NotEmpty(t, data)
	assert.Equal(t, 5, len(data[0]))
	expected := []cellData{
		{
			Letter: 't',
			Color:  "rgb(24,26,27)",
		},
		{
			Letter: 'e',
			Color:  "rgb(255,204,0)",
		},
		{
			Letter: 'a',
			Color:  "rgb(24,26,27)",
		},
		{
			Letter: 'r',
			Color:  "rgb(24,26,27)",
		},
		{
			Letter: 's',
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
