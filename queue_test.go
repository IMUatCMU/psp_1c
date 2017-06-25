package main

import "testing"

func TestQueueOperations(t *testing.T) {
	for _, test := range []struct {
		createList  func() List
		assertQueue func(Queue)
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
			assertQueue: func(queue Queue) {
				var (
					p   interface{}
					err error
				)

				p, err = queue.Dequeue()
				if err != nil || p != 1 {
					t.Log("test failed")
					t.FailNow()
				}

				p, err = queue.Dequeue()
				if err != nil || p != 2 {
					t.Log("test failed")
					t.FailNow()
				}

				p, err = queue.Dequeue()
				if err != nil || p != 3 {
					t.Log("test failed")
					t.FailNow()
				}

				p, err = queue.Dequeue()
				if err == nil {
					t.Log("test failed")
					t.FailNow()
				}
			},
			newElements: []interface{}{4, 5, 6},
			assertList: func(l List) {
				if l.(*list).data[0] != 4 {
					t.Log("test failed")
					t.FailNow()
				}
				if l.(*list).data[1] != 5 {
					t.Log("test failed")
					t.FailNow()
				}
				if l.(*list).data[2] != 6 {
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
			assertQueue: func(queue Queue) {
				var (
					p   interface{}
					err error
				)
				p, err = queue.Dequeue()
				if err != nil || p != "a" {
					t.Log("test failed")
					t.FailNow()
				}

				p, err = queue.Dequeue()
				if err != nil || p != "b" {
					t.Log("test failed")
					t.FailNow()
				}

				p, err = queue.Dequeue()
				if err != nil || p != "c" {
					t.Log("test failed")
					t.FailNow()
				}

				p, err = queue.Dequeue()
				if err == nil {
					t.Log("test failed")
					t.FailNow()
				}
			},
			newElements: []interface{}{"x", "y", "z"},
			assertList: func(l List) {
				if l.(*list).data[0] != "x" {
					t.Log("test failed")
					t.FailNow()
				}
				if l.(*list).data[1] != "y" {
					t.Log("test failed")
					t.FailNow()
				}
				if l.(*list).data[2] != "z" {
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
			assertQueue: func(queue Queue) {},
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
		q := l.ToQueue()
		test.assertQueue(q)

		for _, e := range test.newElements {
			q.Enqueue(e)
		}
		test.assertList(q.ToList())
	}
}

func TestReadQueue(t *testing.T) {
	for _, test := range []struct {
		getFileLocation func() string
		assertQueue     func(q Queue, err error)
	}{
		// 2.1 proper list content read as stack
		{
			getFileLocation: func() string {
				fileName := "/tmp/D899F003-9697-49AE-8925-27220C473279.txt"
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.WriteToFile(fileName)
				return fileName
			},
			assertQueue: func(q Queue, err error) {
				if err != nil {
					t.Log("test failed")
					t.FailNow()

					var (
						p   interface{}
						err error
					)
					p, err = q.Dequeue()
					if err != nil || p != 1 {
						t.Log("test failed")
						t.FailNow()
					}

					p, err = q.Dequeue()
					if err != nil || p != 2 {
						t.Log("test failed")
						t.FailNow()
					}

					p, err = q.Dequeue()
					if err != nil || p != 3 {
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
			assertQueue: func(q Queue, err error) {
				if err == nil {
					t.Log("test failed")
					t.FailNow()
				}
			},
		},
	} {
		s := test.getFileLocation()
		test.assertQueue(ReadQueue(s))
	}
}
