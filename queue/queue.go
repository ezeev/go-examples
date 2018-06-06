package queue

// FIFO
type Queue struct {
	items []int
}

func (q *Queue) Add(n int) {
	q.items = append(q.items, n)
}

func (q *Queue) Remove() int {
	n := q.items[0]
	q.items = q.items[1:]
	return n
}

func (q *Queue) Items() []int {
	return q.items
}

func (q *Queue) Peek() int {
	return q.items[0]
}
