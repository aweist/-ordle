package parse

import (
	"fmt"
	"os"
	"testing"

	"github.com/aweist/ordle/html"
	"github.com/stretchr/testify/assert"
)

func Test_getDordleGameBoards(t *testing.T) {
	f, err := os.Open("test_dordle.html")
	assert.NoError(t, err)

	root, err := html.Parse(f)
	assert.NoError(t, err)

	boards := getDordleGameBoards(root)

	assert.Equal(t, 2, len(boards))
}

func TestParseDordle(t *testing.T) {
	f, err := os.Open("test_dordle.html")
	assert.NoError(t, err)

	states := ParseDordle(f)
	assert.Equal(t, 2, len(states))

	assert.Equal(t, 5, len(states[0].known))
	assert.Equal(t, 2, len(states[0].misplaced))

	fmt.Println("known", states[1].known)
	fmt.Println("misplaced", states[1].misplaced)
	assert.Equal(t, 2, len(states[1].known))
	assert.Equal(t, 2, len(states[1].misplaced))

}
