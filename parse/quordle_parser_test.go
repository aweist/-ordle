package parse

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestGetGameBoards(t *testing.T) {
	f, err := os.Open("test_quordle.html")
	assert.NoError(t, err)

	h, err := html.Parse(f)
	assert.NoError(t, err)

	type args struct {
		root *html.Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base",
			args: args{
				root: h,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetQuordleGameBoards(tt.args.root); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("GetGameBoards() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func Test_ParseQuordle(t *testing.T) {

	f, err := os.Open("test_quordle.html")
	assert.NoError(t, err)

	states := ParseQuordle(f)

	expectations := []struct {
		Known     int
		Misplaced int
	}{
		{
			Known:     5,
			Misplaced: 5,
		},
		{
			Known:     5,
			Misplaced: 7,
		},
		{
			Known:     5,
			Misplaced: 3,
		},
		{
			Known:     1,
			Misplaced: 6,
		},
	}

	for i, s := range states {
		e := expectations[i]
		assert.Equal(t, e.Known, len(s.known), fmt.Sprintf("Incorrect known for state %v", i))
		assert.Equal(t, e.Misplaced, len(s.misplaced), fmt.Sprintf("Incorrect mispalced for state %v", i))
	}

}
