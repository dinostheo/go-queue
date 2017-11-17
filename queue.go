package queue

import "container/list"

type Queue struct {
	storage *list.List
}

// New returns an initialized queue.
func New() *Queue {
	return new(Queue).Init()
}

func (q *Queue) Init() *Queue {
	q.storage.Init()
	return q
}

func (q *Queue) Len() int {
	return q.storage.Len()
}

func (q *Queue) Enqueue(v interface{}) {
	q.storage.PushBack(v)
}

func (q *Queue) Dequeue() interface{} {
	el := q.storage.Front()
	return q.storage.Remove(el)
}
