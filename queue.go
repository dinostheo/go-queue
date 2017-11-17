package queue

import "container/list"

// Queue is a queue structure implemented on top of a linked list
type Queue struct {
	storage *list.List
}

// New returns an initialized queue.
func New() *Queue {
	return new(Queue).Init()
}

// Init initializes our queue through its internal storage (linked list)
func (q *Queue) Init() *Queue {
	q.storage.Init()
	return q
}

// Len returns the length of the queue
func (q *Queue) Len() int {
	return q.storage.Len()
}

// Enqueue adds a value in the end of the queue
func (q *Queue) Enqueue(v interface{}) {
	q.storage.PushBack(v)
}

// Dequeue removes and returns the first value of the queue
func (q *Queue) Dequeue() interface{} {
	el := q.storage.Front()
	return q.storage.Remove(el)
}
