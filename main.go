package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	defaultFileName = "psp4c.txt"
)

func ReadList(filePath string) (List, error) {
	l := &list{}

	if file, err := os.Open(filePath); err != nil {
		return nil, err
	} else {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		i := 0

		for scanner.Scan() {
			if i == 0 {
				if num, err := strconv.Atoi(strings.TrimSpace(scanner.Text())); err != nil {
					return nil, err
				} else {
					l.dt = dataType(num)
				}
			} else {
				switch l.dt {
				case intType:
					if num, err := strconv.Atoi(strings.TrimSpace(scanner.Text())); err != nil {
						return nil, err
					} else {
						l.Add(i-1, num)
					}
				case stringType:
					l.Add(i-1, strings.TrimSpace(scanner.Text()))
				default:
					return nil, errInvalidType("string or number")
				}
			}
			i++
		}
	}

	return l, nil
}

type List interface {
	Stat

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

	// Sort the list
	Sort(ascending bool) error

	// decides the order of two elements
	lessThan(a, b interface{}) bool

	// Split the list at given index, the item at the given index
	// belongs to the second list
	Split(index int) (first List, second List, err error)

	// Write to file
	WriteToFile(fileName string) error

	// Merge with another list
	Merge(another List) (List, error)

	// Check if the list is sorted (ascending or descending)
	IsSorted(expectAsc bool) (bool, error)

	// Convert to Stack
	ToStack() Stack

	// Convert to Queue
	ToQueue() Queue
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

	maxStringLength int = 250
)

var (
	errIndexOutOfBound = func(i int) error { return fmt.Errorf("index out of bound: %d", i) }
	errLengthExceeded  = func(maxLen int) error { return fmt.Errorf("length exceeded: max=%d", maxLen) }
	errInvalidType     = func(exp string) error { return fmt.Errorf("invalid argument type, expected %s", exp) }
	errUnsortable      = errors.New("cannot perform sort on containing data type")
)

type list struct {
	data []interface{}
	dt   dataType
	stat Stat
}

func (l *list) Count() (int, error) {
	return l.stat.Count()
}

func (l *list) Average() (float64, error) {
	return l.stat.Average()
}

func (l *list) Max() (int, error) {
	return l.stat.Max()
}

func (l *list) Min() (int, error) {
	return l.stat.Min()
}

func (l *list) Mean() (float64, error) {
	return l.stat.Mean()
}

func (l *list) Std() (float64, error) {
	return l.stat.Std()
}

func (l *list) typeCheck(arg interface{}) error {
	switch l.dt {
	case intType:
		switch arg.(type) {
		case int, int16, int32, int64:
		case float32, float64:
		default:
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
	} else if err := l.typeCheck(entry); err != nil {
		return err
	}
	if stringType == l.dt {
		if err := l.lengthCheck(entry.(string), maxStringLength); err != nil {
			return err
		}
	}
	switch l.dt {
	case intType:
		var newElem int
		switch entry.(type) {
		case int, int16, int32, int64:
			newElem = entry.(int)
		case float32:
			newElem = int(entry.(float32))
		case float64:
			newElem = int(entry.(float64))
		default:
			return errInvalidType("int")
		}
		l.data = append(l.data[:index], append([]interface{}{newElem}, l.data[index:]...)...)
	default:
		l.data = append(l.data[:index], append([]interface{}{entry}, l.data[index:]...)...)
	}
	return nil
}

func (l *list) Sort(ascending bool) error {
	if l.dt == unspecified {
		return errUnsortable
	} else if len(l.data) == 0 {
		return ErrEmpty
	}
	s := ByOrder{
		Data:     l.data,
		Asc:      ascending,
		LessThan: l.lessThan,
	}
	sort.Sort(s)
	l.data = s.Data
	return nil
}

func (l *list) lessThan(a, b interface{}) bool {
	switch l.dt {
	case intType:
		return a.(int) < b.(int)
	case stringType:
		return a.(string) < b.(string)
	default:
		panic(errInvalidType("string or integer"))
	}
}

func (l *list) greaterThan(a, b interface{}) bool {
	switch l.dt {
	case intType:
		return a.(int) > b.(int)
	case stringType:
		return a.(string) > b.(string)
	default:
		panic(errInvalidType("string or integer"))
	}
}

func (l *list) Split(index int) (first List, second List, err error) {
	if index < 0 || index > len(l.data)-1 {
		return nil, nil, errIndexOutOfBound(index)
	} else if index == 0 {
		return &list{data: make([]interface{}, 0), dt: l.dt}, l, nil
	} else if index == len(l.data)-1 {
		return l, &list{data: make([]interface{}, 0), dt: l.dt}, nil
	} else {
		a, b := l.data[:index], l.data[index:]
		return &list{data: a, dt: l.dt}, &list{data: b, dt: l.dt}, nil
	}
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

func (l *list) WriteToFile(fileName string) error {
	if len(fileName) == 0 {
		fileName = "/tmp/" + defaultFileName
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		if _, err := os.Create(fileName); err != nil {
			return err
		}
	} else {
		log.Println("file already exists, will overwrite.")
	}

	return ioutil.WriteFile(fileName, l.fileContent(), 0644)
}

func (l *list) Merge(another List) (l1 List, err error) {
	l0 := another.(*list)
	if l.dt != l0.dt {
		return nil, errInvalidType("same: string or number")
	}

	switch l.dt {
	case stringType:
		l1 = NewStringList()
	case intType:
		l1 = NewNumList()
	}

	for _, e := range l.data {
		l1.(*list).data = append(l1.(*list).data, e)
	}

	for _, e := range l0.data {
		l1.(*list).data = append(l1.(*list).data, e)
	}

	return l1, err
}

func (l *list) IsSorted(expectAsc bool) (bool, error) {
	switch l.dt {
	case stringType:
	case intType:
	default:
		return false, errInvalidType("string or number")
	}

	for i := range l.data {
		if i > 0 {
			if expectAsc {
				if l.greaterThan(l.data[i-1], l.data[i]) {
					return false, nil
				}
			} else {
				if l.lessThan(l.data[i-1], l.data[i]) {
					return false, nil
				}
			}
		}
	}

	return true, nil
}

func (l *list) ToStack() Stack {
	return &stack{list: l}
}

func (l *list) ToQueue() Queue {
	return &queue{list: l}
}

func (l *list) fileContent() []byte {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("%d\n", l.dt))
	for _, e := range l.data {
		b.WriteString(fmt.Sprintf("%v\n", e))
	}
	return b.Bytes()
}

type ByOrder struct {
	Data     []interface{}
	Asc      bool
	LessThan func(a, b interface{}) bool
}

func (s ByOrder) Len() int {
	return len(s.Data)
}

func (s ByOrder) Swap(i, j int) {
	s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
}

func (s ByOrder) Less(i, j int) bool {
	if s.Asc {
		return s.LessThan(s.Data[i], s.Data[j])
	} else {
		return !s.LessThan(s.Data[i], s.Data[j])
	}
}

func NewList() List {
	return &list{data: make([]interface{}, 0), dt: unspecified, stat: noStat{}}
}

func NewNumList() NumList {
	l := &list{data: make([]interface{}, 0), dt: intType}
	l.stat = numberStat{l: l}
	return l
}

func NewStringList() StringList {
	l := &list{data: make([]interface{}, 0), dt: stringType}
	l.stat = stringStat{l: l}
	return l
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
