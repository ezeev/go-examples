package arraychunk_test

import (
	"testing"

	"github.com/ezeev/go-examples/arraychunk"
)

func TestChunk(t *testing.T) {
	in := []int{3, 5, 1, 12, 45, 65, 23, 65, 9, 100}
	chunks := arraychunk.IntChunk(in, 3)
	t.Logf("intput: %d", in)
	t.Logf("output: %d", chunks)

	in = []int{1, 2, 3, 4, 5, 6}
	chunks = arraychunk.IntChunk(in, 4)
	t.Log(chunks)

}
