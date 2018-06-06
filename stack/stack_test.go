package stack_test

import (
	"testing"

	"github.com/ezeev/go-examples/stack"
)

func TestStack(t *testing.T) {

	s := stack.Stack{}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	t.Log(s)

	t.Logf("removing %d", s.Pop())
	t.Log(s)

	t.Logf("removing %d", s.Pop())
	t.Log(s)

	t.Logf("removing %d", s.Pop())
	t.Log(s)

	s.Push(6)
	s.Push(7)
	s.Push(8)

	t.Log(s)

	t.Logf("removing %d", s.Pop())
	t.Log(s)

	t.Logf("removing %d", s.Pop())
	t.Log(s)

	t.Logf("removing %d", s.Pop())
	t.Log(s)

}
