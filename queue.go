package main

import "github.com/pkg/errors"

type Queue interface {
	Enqueue(item interface{}) error
	Dequeue() (interface{}, error)
	ToList() List
}

func ReadQueue(filePath string) (Queue, error) {
	if l, err := ReadList(filePath); err != nil {
		return nil, err
	} else {
		return l.(*list).ToQueue(), nil
	}
}

type queue struct {
	list *list
}

func (q *queue) Enqueue(item interface{}) error {
	return q.list.Add(len(q.list.data), item)
}

func (q *queue) Dequeue() (interface{}, error) {
	if len(q.list.data) <= 0 {
		return nil, errors.New("queue is empty")
	}

	e := q.list.data[0]
	err := q.list.Remove(0)
	return e, err
}

func (q *queue) ToList() List {
	return q.list
}
