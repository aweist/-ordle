package html

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestFindNodeByAttr(t *testing.T) {
	f, err := os.Open("test_input.html")
	assert.NoError(t, err)

	h, err := html.Parse(f)
	assert.NoError(t, err)
	type args struct {
		root  *html.Node
		key   string
		value string
	}
	tests := []struct {
		name       string
		args       args
		numResults int
	}{
		{
			name: "base",
			args: args{
				root:  h,
				key:   "id",
				value: "box3,2,1",
			},
			numResults: 1,
		},
		{
			name: "base",
			args: args{
				root:  h,
				key:   "class",
				value: "box button",
			},
			numResults: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResults := FindNodeByAttr(tt.args.root, tt.args.key, tt.args.value); !reflect.DeepEqual(len(gotResults), tt.numResults) {
				t.Errorf("FindNodeByAttr() = %v, want %v", gotResults, tt.numResults)
			}
		})
	}
}

func TestGetAttr(t *testing.T) {
	s := `<p foo="bar"></p>`
	root, err := html.Parse(strings.NewReader(s))
	assert.NoError(t, err)
	node := FindNodeByAttr(root, "foo", "bar")

	type args struct {
		node *html.Node
		key  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				node: node[0],
				key:  "foo",
			},
			want: "bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAttr(tt.args.node, tt.args.key); got != tt.want {
				t.Errorf("GetAttr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cellValue(t *testing.T) {
	s := `<p foo="bar">TOO</p>`
	root, err := html.Parse(strings.NewReader(s))
	assert.NoError(t, err)
	node := FindNodeByAttr(root, "foo", "bar")

	type args struct {
		cell *html.Node
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "base",
			args: args{
				cell: node[0],
			},
			want: 't',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nodeValue(tt.args.cell); got != tt.want {
				t.Errorf("cellValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
