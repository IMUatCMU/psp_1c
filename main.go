package main

import (
	"fmt"
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

type NumList interface {
	List
}

type StringList interface {
	List

	// update an entry identified by the index
	//
	// returns error if index is out of bounds
	Update(index int, entry string) error
}

type dataType int

const (
	unspecified dataType = 0
	intType     dataType = 1
	stringType  dataType = 2

	maxListCapacity int = 20
	maxStringLength int = 250
)

var (
	errIndexOutOfBound = func(i int) error { return fmt.Errorf("index out of bound: %d", i) }
	errCapacityReached = func(cap int) error { return fmt.Errorf("capacity reached: cap=%d", cap) }
	errLengthExceeded  = func(maxLen int) error { return fmt.Errorf("length exceeded: max=%d", maxLen) }
	errInvalidType     = func(exp string) error { return fmt.Errorf("invalid argument type, expected %s", exp) }
)

type list struct {
	data []interface{}
	dt   dataType
}

func (l *list) typeCheck(arg interface{}) error {
	switch l.dt {
	case intType:
		if _, ok := arg.(int); !ok {
			return errInvalidType("int")
		}
	case stringType:
		if _, ok := arg.(string); !ok {
			return errInvalidType("string")
		}
	}
	return nil
}

func (l *list) lengthCheck(arg string, maxLen int) error {
	if len(arg) > maxLen {
		return errLengthExceeded(maxStringLength)
	}
	return nil
}

func (l *list) Print() {
	for i, e := range l.data {
		fmt.Printf("%d: %v\n", i, e)
	}
}

func (l *list) Remove(index int) error {
	if index < 0 || index > len(l.data)-1 {
		return errIndexOutOfBound(index)
	}
	l.data = append(l.data[:index], l.data[index+1:]...)
	return nil
}

func (l *list) Add(index int, entry interface{}) error {
	if index < 0 || index > len(l.data) {
		return errIndexOutOfBound(index)
	} else if len(l.data) == 20 {
		return errCapacityReached(maxListCapacity)
	} else if err := l.typeCheck(entry); err != nil {
		return err
	}
	if stringType == l.dt {
		if err := l.lengthCheck(entry.(string), maxStringLength); err != nil {
			return err
		}
	}
	l.data = append(l.data[:index], append([]interface{}{entry}, l.data[index:]...)...)
	return nil
}

func (l *list) Update(index int, entry string) error {
	if index < 0 || index > len(l.data)-1 {
		return errIndexOutOfBound(index)
	} else if err := l.lengthCheck(entry, maxStringLength); err != nil {
		return err
	}
	l.data[index] = entry
	return nil
}

func NewList() List {
	return &list{data: make([]interface{}, 0, maxListCapacity), dt: unspecified}
}

func NewNumList() NumList {
	return &list{data: make([]interface{}, 0, maxListCapacity), dt: intType}
}

func NewStringList() StringList {
	return &list{data: make([]interface{}, 0, maxListCapacity), dt: stringType}
}

func main() {
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
