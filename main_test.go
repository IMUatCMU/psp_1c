package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

const (
	fiftyChars = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
)

func TestNewList(t *testing.T) {
	l := NewList()
	if l == nil {
		t.Error("list is supposed to be non-nil")
	}

	if l0, ok := l.(*list); !ok {
		t.Error("list is supposed to be of type *list")
	} else {
		if len(l0.data) != 0 {
			t.Error("list is supposed to be empty from the start")
		}
		if cap(l0.data) != 20 {
			t.Error("list is supposed to be capped at 20")
		}
	}
}

func TestNewNumList(t *testing.T) {
	l := NewNumList()
	if l == nil {
		t.Error("list is supposed to be non-nil")
	}

	if l0, ok := l.(*list); !ok {
		t.Error("list is supposed to be of type *list")
	} else {
		if len(l0.data) != 0 {
			t.Error("list is supposed to be empty from the start")
		}
		if cap(l0.data) != 20 {
			t.Error("list is supposed to be capped at 20")
		}
	}
}

func TestNewStringList(t *testing.T) {
	l := NewStringList()
	if l == nil {
		t.Error("list is supposed to be non-nil")
	}

	if l0, ok := l.(*list); !ok {
		t.Error("list is supposed to be of type *list")
	} else {
		if len(l0.data) != 0 {
			t.Error("list is supposed to be empty from the start")
		}
		if cap(l0.data) != 20 {
			t.Error("list is supposed to be capped at 20")
		}
	}
}

func TestList_Add(t *testing.T) {
	l := NewList()

	err := l.Add(-1, "foo")
	if err == nil {
		t.Error("Add should have encountered index out of bounds")
	}

	err = l.Add(100, "foo")
	if err == nil {
		t.Error("Add should have encountered index out of bounds")
	}

	for i, e := range []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"11",
		"12",
		"13",
		"14",
		"15",
		"16",
		"17",
		"18",
		"19",
		"20",
	} {
		if err := l.Add(i, e); err != nil {
			t.Error("Add should have performed without error")
		}
	}

	err = l.Add(20, "foo")
	if err == nil {
		t.Error("Add should have failed after exceeding cap")
	}
}

func TestNumList_Add(t *testing.T) {
	l := NewNumList()

	err := l.Add(-1, 123)
	if err == nil {
		t.Error("Add should have encountered index out of bounds")
	}

	err = l.Add(100, 123)
	if err == nil {
		t.Error("Add should have encountered index out of bounds")
	}

	err = l.Add(0, "foo")
	if err == nil {
		t.Error("Add should have encountered invalid argument type")
	}

	for i, e := range []int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		11,
		12,
		13,
		14,
		15,
		16,
		17,
		18,
		19,
		20,
	} {
		if err := l.Add(i, e); err != nil {
			t.Error("Add should have performed without error")
		}
	}

	err = l.Add(20, 123)
	if err == nil {
		t.Error("Add should have failed after exceeding cap")
	}
}

func TestStringList_Add(t *testing.T) {
	l := NewStringList()

	err := l.Add(-1, "foo")
	if err == nil {
		t.Error("Add should have encountered index out of bounds")
	}

	err = l.Add(100, "foo")
	if err == nil {
		t.Error("Add should have encountered index out of bounds")
	}

	err = l.Add(0, 123)
	if err == nil {
		t.Error("Add should have encountered invalid argument type")
	}

	// 251 chars
	err = l.Add(0, fmt.Sprintf("%s%s%s%s%sx", fiftyChars, fiftyChars, fiftyChars, fiftyChars, fiftyChars))
	if err == nil {
		t.Error("Add should have encountered length exceeded error")
	}

	for i, e := range []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"11",
		"12",
		"13",
		"14",
		"15",
		"16",
		"17",
		"18",
		"19",
		"20",
	} {
		if err := l.Add(i, e); err != nil {
			t.Error("Add should have performed without error")
		}
	}

	err = l.Add(20, "foo")
	if err == nil {
		t.Error("Add should have failed after exceeding cap")
	}
}

func TestList_Remove(t *testing.T) {
	l := NewList()

	err := l.Remove(0)
	if err == nil {
		t.Error("Remove on empty list should have failed")
	}

	l.Add(0, "1")
	l.Add(1, "2")
	l.Add(2, "3")

	l.Remove(1)
	if len(l.(*list).data) != 2 {
		t.Error("Expected length to be 2")
	} else if l.(*list).data[0] != "1" {
		t.Error("Remove op resulted in incorrect array")
	} else if l.(*list).data[1] != "3" {
		t.Error("Remove op resulted in incorrect array")
	}

	l.Remove(1)
	if len(l.(*list).data) != 1 {
		t.Error("Expected length to be 1")
	} else if l.(*list).data[0] != "1" {
		t.Error("Remove op resulted in incorrect array")
	}

	l.Remove(0)
	if len(l.(*list).data) != 0 {
		t.Error("Expected length to be 0")
	}
}

func TestStringList_Update(t *testing.T) {
	l := NewStringList()

	l.Add(0, "foo")
	l.Update(0, "bar")

	if l.(*list).data[0] != "bar" {
		t.Error("Should have been updated to 'bar'")
	}

	err := l.Update(0, fmt.Sprintf("%s%s%s%s%sx", fiftyChars, fiftyChars, fiftyChars, fiftyChars, fiftyChars))
	if err == nil {
		t.Error("Add should have encountered length exceeded error")
	}
}

func TestList_Sort(t *testing.T) {
	for _, test := range []struct {
		listGen     func() List
		origData    []interface{}
		ascending   bool
		compareData []interface{}
	}{
		{
			listGen:     func() List { return NewNumList() },
			origData:    []interface{}{5, 3, 1, 2, 4},
			ascending:   true,
			compareData: []interface{}{1, 2, 3, 4, 5},
		},
		{
			listGen:     func() List { return NewNumList() },
			origData:    []interface{}{5, 3, 1, 2, 4},
			ascending:   false,
			compareData: []interface{}{5, 4, 3, 2, 1},
		},
		{
			listGen:     func() List { return NewStringList() },
			origData:    []interface{}{"a", "x", "r", "e", "z"},
			ascending:   true,
			compareData: []interface{}{"a", "e", "r", "x", "z"},
		},
		{
			listGen:     func() List { return NewStringList() },
			origData:    []interface{}{"a", "x", "r", "e", "z"},
			ascending:   false,
			compareData: []interface{}{"z", "x", "r", "e", "a"},
		},
		{
			listGen:     func() List { return NewList() },
			origData:    []interface{}{"a", "x", "r", "e", "z"},
			ascending:   false,
			compareData: []interface{}{"a", "x", "r", "e", "z"},
		},
	} {
		l := test.listGen()
		for i, e := range test.origData {
			l.Add(i, e)
		}
		l.Sort(test.ascending)
		for i, e := range test.compareData {
			if l.(*list).data[i] != e {
				t.Errorf("sort error, element at index %d should be %v", i, e)
			}
		}
		l.Print()
	}
}

func TestList_Split(t *testing.T) {
	for _, test := range []struct {
		listGen     func() List
		splitIndex  int
		expectError bool
		firstList   func() List
		secondList  func() List
	}{
		{
			listGen: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.Add(3, 4)
				l.Add(4, 5)
				l.Add(5, 6)
				return l
			},
			splitIndex:  3,
			expectError: false,
			firstList: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				return l
			},
			secondList: func() List {
				l := NewNumList()
				l.Add(0, 4)
				l.Add(1, 5)
				l.Add(2, 6)
				return l
			},
		},
		{
			listGen: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.Add(3, 4)
				l.Add(4, 5)
				l.Add(5, 6)
				return l
			},
			splitIndex:  0,
			expectError: false,
			firstList: func() List {
				l := NewNumList()
				return l
			},
			secondList: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.Add(3, 4)
				l.Add(4, 5)
				l.Add(5, 6)
				return l
			},
		},
		{
			listGen: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.Add(3, 4)
				l.Add(4, 5)
				l.Add(5, 6)
				return l
			},
			splitIndex:  5,
			expectError: false,
			firstList: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.Add(3, 4)
				l.Add(4, 5)
				l.Add(5, 6)
				return l
			},
			secondList: func() List {
				l := NewNumList()
				return l
			},
		},
		{
			listGen: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.Add(3, 4)
				l.Add(4, 5)
				l.Add(5, 6)
				return l
			},
			splitIndex:  -1,
			expectError: true,
			firstList: func() List {
				return nil
			},
			secondList: func() List {
				return nil
			},
		},
		{
			listGen: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.Add(3, 4)
				l.Add(4, 5)
				l.Add(5, 6)
				return l
			},
			splitIndex:  100,
			expectError: true,
			firstList: func() List {
				return nil
			},
			secondList: func() List {
				return nil
			},
		},
	} {
		l := test.listGen()
		a, b, err := l.Split(test.splitIndex)

		if test.expectError {
			if err == nil {
				t.Error("expects error but got no error")
			}
		} else {
			if err != nil {
				t.Error("expects no error but got", err)
			}

			if !reflect.DeepEqual(test.firstList().(*list).data, a.(*list).data) {
				t.Error("split error, the two lists are not equal", a.(*list).data)
			}
			if !reflect.DeepEqual(test.secondList().(*list).data, b.(*list).data) {
				t.Error("split error, the two lists are not equal", b.(*list).data)
			}

			a.Print()
			b.Print()
			fmt.Println()
		}
	}
}

func TestList_WriteToFile(t *testing.T) {
	for i, test := range []struct {
		fileName        string
		list            func() List
		expectedFile    string
		expectedContent string
	}{
		{
			"/tmp/52DB7899-E20B-4166-9E20-BA22A6250A1D.txt",
			func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "b")
				l.Add(2, "c")
				return l
			},
			"/tmp/52DB7899-E20B-4166-9E20-BA22A6250A1D.txt",
			"2\na\nb\nc\n",
		},
		{
			"/tmp/D36C27D7-EDD7-48EC-8421-25CB771CAFB0.txt",
			func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				return l
			},
			"/tmp/D36C27D7-EDD7-48EC-8421-25CB771CAFB0.txt",
			"1\n1\n2\n3\n",
		},
		{
			"/tmp/" + defaultFileName,
			func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "b")
				l.Add(2, "c")
				return l
			},
			"/tmp/" + defaultFileName,
			"2\na\nb\nc\n",
		},
		{
			"/tmp/" + defaultFileName,
			func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "b")
				l.Add(2, "c")
				return l
			},
			"/tmp/" + defaultFileName,
			"2\na\nb\nc\n",
		},
	} {
		t.Logf("testing case %d\n", i+1)

		l := test.list()
		err := l.WriteToFile(test.fileName)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if _, err := os.Stat(test.expectedFile); os.IsNotExist(err) {
			t.Errorf("Expected %s to exist", test.expectedFile)
			t.FailNow()
		}

		b, err := ioutil.ReadFile(test.expectedFile)
		if err != nil {
			t.Errorf("Expected reading %s is error free", test.expectedFile)
			t.FailNow()
		}

		if string(b) != test.expectedContent {
			t.Errorf("expected %s, got %s", test.expectedContent, string(b))
			t.FailNow()
		}
	}
}

func TestReadList(t *testing.T) {
	for _, test := range []struct {
		getList func() List
		path    string
		compare func(o, n List)
	}{
		{
			getList: func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "b")
				l.Add(2, "c")
				return l
			},
			path: "/tmp/DBC6BAF8-6150-4C9E-8CB2-FD4880AF2CD8.txt",
			compare: func(o, n List) {
				oc := string(o.(*list).fileContent())
				nc := string(n.(*list).fileContent())
				if oc != nc {
					t.Errorf("expected two lists to be the same, but o is %s, and n is %s", oc, nc)
					t.FailNow()
				} else {
					t.Log(oc)
					t.Log(nc)
				}
			},
		},
		{
			getList: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				return l
			},
			path: "/tmp/8A9CA546-FE47-47A2-8FAD-E0C9B0DDDC11.txt",
			compare: func(o, n List) {
				oc := string(o.(*list).fileContent())
				nc := string(n.(*list).fileContent())
				if oc != nc {
					t.Errorf("expected two lists to be the same, but o is %s, and n is %s", oc, nc)
					t.FailNow()
				} else {
					t.Log(oc)
					t.Log(nc)
				}
			},
		},
	} {
		o := test.getList()
		err := o.WriteToFile(test.path)
		if err != nil {
			t.Errorf("expected no error when writing file, but got %v", err)
			t.FailNow()
		}

		n, err := ReadList(test.path)
		if err != nil {
			t.Errorf("expected no error when reading file, but got %v", err)
			t.FailNow()
		}

		test.compare(o, n)
	}
}

func TestList_Merge(t *testing.T) {
	for _, test := range []struct {
		firstList  func() List
		secondList func() List
		assertion  func(List, error)
	}{
		// 2.1 string (10 element) with string (10 element) : expect string 20 element
		{
			firstList: func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "b")
				l.Add(2, "c")
				l.Add(3, "d")
				l.Add(4, "e")
				l.Add(5, "f")
				l.Add(6, "g")
				l.Add(7, "h")
				l.Add(8, "i")
				l.Add(9, "j")
				return l
			},
			secondList: func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "b")
				l.Add(2, "c")
				l.Add(3, "d")
				l.Add(4, "e")
				l.Add(5, "f")
				l.Add(6, "g")
				l.Add(7, "h")
				l.Add(8, "i")
				l.Add(9, "j")
				return l
			},
			assertion: func(l List, err error) {
				if err != nil {
					t.Errorf("expected error free, but got %v", err)
					t.FailNow()
				} else {
					lc := string(l.(*list).fileContent())
					if lc != "2\na\nb\nc\nd\ne\nf\ng\nh\ni\nj\na\nb\nc\nd\ne\nf\ng\nh\ni\nj\n" {
						t.Errorf("unexpected merged content: %s", lc)
						t.FailNow()
					}
				}
			},
		},
		// 2.2 number (5 element) with number (5 element): expect number 10 element
		{
			firstList: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.Add(3, 4)
				l.Add(4, 5)
				return l
			},
			secondList: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.Add(3, 4)
				l.Add(4, 5)
				return l
			},
			assertion: func(l List, err error) {
				if err != nil {
					t.Errorf("expected error free, but got %v", err)
					t.FailNow()
				} else {
					lc := string(l.(*list).fileContent())
					if lc != "1\n1\n2\n3\n4\n5\n1\n2\n3\n4\n5\n" {
						t.Errorf("unexpected merged content: %s", lc)
						t.FailNow()
					}
				}
			},
		},
		// 2.3 string (5 element) with number (5 element): expect incompatible type error
		{
			firstList: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.Add(3, 4)
				l.Add(4, 5)
				return l
			},
			secondList: func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "b")
				l.Add(2, "c")
				l.Add(3, "d")
				l.Add(4, "e")
				return l
			},
			assertion: func(l List, err error) {
				if err == nil {
					t.Error("expected error free, but didn't get any")
					t.FailNow()
				}
			},
		},
		// 2.4 number (5 element) with string (5 element): expect incompatible type error
		{
			firstList: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.Add(3, 4)
				l.Add(4, 5)
				return l
			},
			secondList: func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "b")
				l.Add(2, "c")
				l.Add(3, "d")
				l.Add(4, "e")
				return l
			},
			assertion: func(l List, err error) {
				if err == nil {
					t.Error("expected error free, but didn't get any")
					t.FailNow()
				}
			},
		},
		// 2.5 string (11 element) with string (11 element): expect string 20 element with list truncated error
		{
			firstList: func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "b")
				l.Add(2, "c")
				l.Add(3, "d")
				l.Add(4, "e")
				l.Add(5, "f")
				l.Add(6, "g")
				l.Add(7, "h")
				l.Add(8, "i")
				l.Add(9, "j")
				l.Add(10, "k")
				return l
			},
			secondList: func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "b")
				l.Add(2, "c")
				l.Add(3, "d")
				l.Add(4, "e")
				l.Add(5, "f")
				l.Add(6, "g")
				l.Add(7, "h")
				l.Add(8, "i")
				l.Add(9, "j")
				l.Add(10, "k")
				return l
			},
			assertion: func(l List, err error) {
				if err == nil {
					t.Error("expected error free, but didn't get any")
					t.FailNow()
				}

				lc := string(l.(*list).fileContent())
				if lc != "2\na\nb\nc\nd\ne\nf\ng\nh\ni\nj\nk\na\nb\nc\nd\ne\nf\ng\nh\ni\n" {
					t.Errorf("unexpected merged content: %s", lc)
					t.FailNow()
				}
			},
		},
		// 2.6 number (11 element) with number (11 element): expect number 20 element with list truncated error
		{
			firstList: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.Add(3, 4)
				l.Add(4, 5)
				l.Add(5, 6)
				l.Add(6, 7)
				l.Add(7, 8)
				l.Add(8, 9)
				l.Add(9, 10)
				l.Add(10, 11)
				return l
			},
			secondList: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				l.Add(3, 4)
				l.Add(4, 5)
				l.Add(5, 6)
				l.Add(6, 7)
				l.Add(7, 8)
				l.Add(8, 9)
				l.Add(9, 10)
				l.Add(10, 11)
				return l
			},
			assertion: func(l List, err error) {
				if err == nil {
					t.Error("expected error free, but didn't get any")
					t.FailNow()
				}

				lc := string(l.(*list).fileContent())
				if lc != "1\n1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n1\n2\n3\n4\n5\n6\n7\n8\n9\n" {
					t.Errorf("unexpected merged content: %s", lc)
					t.FailNow()
				}
			},
		},
	} {
		l1 := test.firstList()
		l2 := test.secondList()
		result, err := l1.Merge(l2)
		test.assertion(result, err)
	}
}

func TestList_IsSorted(t *testing.T) {
	for _, test := range []struct {
		getList   func() List
		asc       bool
		assertion func(bool, error)
	}{
		// 3.1 check sorted string list: expect true
		{
			getList: func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "b")
				l.Add(2, "c")
				return l
			},
			asc: true,
			assertion: func(r bool, err error) {
				if err != nil {
					t.Errorf("expected error free, but got %v", err)
					t.FailNow()
				}

				if r != true {
					t.Errorf("expected sorted, but got %v", r)
					t.FailNow()
				}
			},
		},
		// 3.2 check sorted number list: expect true
		{
			getList: func() List {
				l := NewNumList()
				l.Add(0, 3)
				l.Add(1, 2)
				l.Add(2, 1)
				return l
			},
			asc: false,
			assertion: func(r bool, err error) {
				if err != nil {
					t.Errorf("expected error free, but got %v", err)
					t.FailNow()
				}

				if r != true {
					t.Errorf("expected sorted, but got %v", r)
					t.FailNow()
				}
			},
		},
		// 3.3 check unsorted string list: expect false
		{
			getList: func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "c")
				l.Add(2, "b")
				return l
			},
			asc: true,
			assertion: func(r bool, err error) {
				if err != nil {
					t.Errorf("expected error free, but got %v", err)
					t.FailNow()
				}

				if r != false {
					t.Errorf("expected unsorted, but got %v", r)
					t.FailNow()
				}
			},
		},
		// 3.4 check unsorted number list: expect false
		{

			getList: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 1)
				return l
			},
			asc: false,
			assertion: func(r bool, err error) {
				if err != nil {
					t.Errorf("expected error free, but got %v", err)
					t.FailNow()
				}

				if r != false {
					t.Errorf("expected unsorted, but got %v", r)
					t.FailNow()
				}
			},
		},
		// 3.5 check sorted generic list, expect error: cannot determine sort logic
		{

			getList: func() List {
				l := NewList()
				l.Add(0, "a")
				l.Add(1, 2)
				l.Add(2, "c")
				return l
			},
			asc: true,
			assertion: func(r bool, err error) {
				if err == nil {
					t.Error("expected error, but didn't get any")
					t.FailNow()
				}
			},
		},
	} {
		l := test.getList()
		r, err := l.IsSorted(test.asc)
		test.assertion(r, err)
	}
}
