package main

import (
	"fmt"
	"errors"
)

type List interface {
	// print out the list and order of the items
	Print()

	// remove an item
	//
	// returns error if index is out of bounds
	Remove(index int) error

	// add an entry after an item identified by the index
	//
	// returns error if index is out of bounds
	Add(index int, entry interface{}) error
}

type list struct {
	data 	[]interface{}
}

func (l *list) Print() {
	for i, e := range l.data {
		fmt.Printf("%d: %v\n", i, e)
	}
}

func (l *list) Remove(index int) error {
	if index < 0 || index > len(l.data) - 1 {
		return errors.New("index out of bound")
	}
	l.data = append(l.data[:index], l.data[index+1:]...)
	return nil
}

func (l *list) Add(index int, entry interface{}) error {
	if index < 0 || index > len(l.data) {
		return errors.New("index out of bound")
	} else if len(l.data) == 20 {
		return errors.New("cannot store any more items")
	}
	l.data = append(l.data[:index], append([]interface{}{entry}, l.data[index:]...)...)
	return nil
}

func NewList() List {
	return &list{data:make([]interface{}, 0, 20)}
}

func main()  {
	errCheck := func(err error) {
		if err != nil {
			panic(err)
		}
	}
	l := NewList()
	errCheck(l.Add(0, "a"))
	errCheck(l.Add(0, "b"))
	errCheck(l.Add(2, "c"))
	l.Print()
}