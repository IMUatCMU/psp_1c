package main

type Stack interface {
	Push(item interface{}) error
	Pop() (interface{}, error)
	ToList() List
}

func ReadStack(filePath string) (Stack, error) {
	if l, err := ReadList(filePath); err != nil {
		return nil, err
	} else {
		return l.(*list).ToStack(), nil
	}
}

type stack struct {
	list *list
}

func (s *stack) Push(item interface{}) error {
	return s.list.Add(0, item)
}

func (s *stack) Pop() (interface{}, error) {
	lastIndex := len(s.list.data) - 1
	if lastIndex < 0 {
		return nil, errIndexOutOfBound(lastIndex)
	}

	e := s.list.data[lastIndex]
	err := s.list.Remove(lastIndex)
	return e, err
}

func (s *stack) ToList() List {
	return s.list
}
