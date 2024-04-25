package collections

// You can create custom collections by defining your own struct types and implementing methods for them. Let’s create a custom collection: a Queue. The Queue collection follows the First-In-First-Out (FIFO) principle.

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	if len(q.items) < 1 {
		return -1 // returns -1 if the queue is empty
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}