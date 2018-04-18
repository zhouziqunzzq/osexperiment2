package main

type Queue struct {
	Queue []interface{}
}

func NewQueue() *Queue {
	q := &Queue{}
	q.Queue = make([]interface{}, 0)
	return q
}

func (q *Queue) Push(item interface{}) {
	q.Queue = append(q.Queue, item)
}

func (q *Queue) Top() interface{} {
	return q.Queue[0]
}

func (q *Queue) Pop() interface{} {
	t := q.Top()
	q.Queue = q.Queue[1:]
	return t
}

func (q *Queue) IsEmpty() bool {
	return len(q.Queue) == 0
}

func (q *Queue) Count() int {
	return len(q.Queue)
}
