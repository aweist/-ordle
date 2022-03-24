package parse

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	html "golang.org/x/net/html"
)

func TestGetOctordleGameBoards(t *testing.T) {
	f, err := os.Open("test_octordle.html")
	assert.NoError(t, err)
	root, err := html.Parse(f)
	assert.NoError(t, err)

	boards := GetOctordleGameBoards(root)
	assert.Equal(t, 8, len(boards))
}

func Test_styleMap(t *testing.T) {
	tests := []struct {
		name  string
		style string
		want  map[string]string
	}{
		{
			name:  "base",
			style: "color: black; background-color: rgb(255, 204, 0);",
			want: map[string]string{
				"color":            "black",
				"background-color": "rgb(255,204,0)",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := styleMap(tt.style); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("styleMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseOctordle(t *testing.T) {
	type want struct {
		misplaced int
		known     int
	}

	filename := "test_octordle.html"
	wants := []want{
		{
			misplaced: 3,
		},
		{
			misplaced: 3,
		},
		{
			misplaced: 5,
		},
		{
			misplaced: 3,
			known:     1,
		},
		{
			misplaced: 2,
			known:     2,
		},
		{
			misplaced: 1,
			known:     2,
		},
		{
			misplaced: 4,
			known:     1,
		},
		{
			misplaced: 4,
		},
	}

	r, err := os.Open(filename)
	assert.NoError(t, err)

	states := ParseOctordle(r)
	assert.Equal(t, len(wants), len(states))
	for i, s := range states {
		assert.Equal(t, wants[i].misplaced, len(s.misplaced), "State %d", i)
		assert.Equal(t, wants[i].known, len(s.known), "State %d", i)
	}
}
