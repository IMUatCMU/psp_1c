package main

import "testing"

func TestStackOperations(t *testing.T) {
	for _, test := range []struct {
		createList  func() List
		assertStack func(Stack)
		newElements []interface{}
		assertList  func(List)
	}{
		// 1.1 number list and pop after empty
		{
			createList: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				return l
			},
			assertStack: func(stack Stack) {
				var (
					p   interface{}
					err error
				)

				p, err = stack.Pop()
				if err != nil || p != 3 {
					t.Log("test failed")
					t.FailNow()
				}

				p, err = stack.Pop()
				if err != nil || p != 2 {
					t.Log("test failed")
					t.FailNow()
				}

				p, err = stack.Pop()
				if err != nil || p != 1 {
					t.Log("test failed")
					t.FailNow()
				}

				p, err = stack.Pop()
				if err == nil {
					t.Log("test failed")
					t.FailNow()
				}
			},
			newElements: []interface{}{4, 5, 6},
			assertList: func(l List) {
				if l.(*list).data[0] != 6 {
					t.Log("test failed")
					t.FailNow()
				}
				if l.(*list).data[1] != 5 {
					t.Log("test failed")
					t.FailNow()
				}
				if l.(*list).data[2] != 4 {
					t.Log("test failed")
					t.FailNow()
				}
			},
		},
		// 1.2 string list
		{
			createList: func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "b")
				l.Add(2, "c")
				return l
			},
			assertStack: func(stack Stack) {
				var (
					p   interface{}
					err error
				)
				p, err = stack.Pop()
				if err != nil || p != "c" {
					t.Log("test failed")
					t.FailNow()
				}

				p, err = stack.Pop()
				if err != nil || p != "b" {
					t.Log("test failed")
					t.FailNow()
				}

				p, err = stack.Pop()
				if err != nil || p != "a" {
					t.Log("test failed")
					t.FailNow()
				}

				p, err = stack.Pop()
				if err == nil {
					t.Log("test failed")
					t.FailNow()
				}
			},
			newElements: []interface{}{"x", "y", "z"},
			assertList: func(l List) {
				if l.(*list).data[0] != "z" {
					t.Log("test failed")
					t.FailNow()
				}
				if l.(*list).data[1] != "y" {
					t.Log("test failed")
					t.FailNow()
				}
				if l.(*list).data[2] != "x" {
					t.Log("test failed")
					t.FailNow()
				}
			},
		},
		// 1.3 push after full
		{
			createList: func() List {
				l := NewNumList()
				return l
			},
			assertStack: func(stack Stack) {},
			newElements: []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
			assertList: func(l List) {
				if len(l.(*list).data) != 20 {
					t.Log("test failed")
					t.FailNow()
				}
			},
		},
	} {
		l := test.createList()
		s := l.ToStack()
		test.assertStack(s)

		for _, e := range test.newElements {
			s.Push(e)
		}
		test.assertList(s.ToList())
	}
}

func TestReadStack(t *testing.T) {
	for _, test := range []struct {
		getFileLocation func() string
		assertStack     func(s Stack, err error)
	}{
		// 2.1 proper list content read as stack
		{
			getFileLocation: func() string {
				fileName := "/tmp/22F8830D-AD8A-4511-B60B-F7943AFEEDCD.txt"
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.WriteToFile(fileName)
				return fileName
			},
			assertStack: func(s Stack, err error) {
				if err != nil {
					t.Log("test failed")
					t.FailNow()

					var (
						p   interface{}
						err error
					)
					p, err = s.Pop()
					if err != nil || p != 3 {
						t.Log("test failed")
						t.FailNow()
					}

					p, err = s.Pop()
					if err != nil || p != 2 {
						t.Log("test failed")
						t.FailNow()
					}

					p, err = s.Pop()
					if err != nil || p != 1 {
						t.Log("test failed")
						t.FailNow()
					}
				}
			},
		},
		// 2.2 try to read an unexisting file
		{
			getFileLocation: func() string {
				return "/tmp/not_existing.txt"
			},
			assertStack: func(s Stack, err error) {
				if err == nil {
					t.Log("test failed")
					t.FailNow()
				}
			},
		},
	} {
		s := test.getFileLocation()
		test.assertStack(ReadStack(s))
	}
}
