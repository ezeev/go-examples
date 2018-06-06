package queue_test

import (
	"testing"

	"github.com/ezeev/go-examples/queue"
)

func TestQueue(t *testing.T) {
	q := queue.Queue{}
	q.Add(1)
	q.Add(2)
	q.Add(3)
	q.Add(4)
	t.Log(q.Items())

	t.Logf("removing %d\n", q.Remove())
	t.Log(q.Items())
	q.Add(5)
	t.Logf("removing %d\n", q.Remove())
	q.Add(6)
	t.Log(q.Items())

	t.Logf("the next item to be removed will be: %d\n", q.Peek())
	t.Logf("removing %d\n", q.Remove())
	t.Log(q.Items())

}
