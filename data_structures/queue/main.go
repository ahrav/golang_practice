package main

// Enqueue
type Queue struct {
	items []int
}

// Enqueue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue
func (q *Queue) Dequeue() int {
	val := q.items[0]
	q.items = q.items[1:]
	return val
}

func main() {

}
