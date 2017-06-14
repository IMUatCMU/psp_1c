package main

import (
	"fmt"
	"reflect"
	"testing"
	"os"
	"io/ioutil"
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
	for i, test := range []struct{
		fileName 	string
		list 		func() List
		expectedFile	string
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
			"a\nb\nc\n",
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
			"1\n2\n3\n",
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
			"a\nb\nc\n",
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
			"a\nb\nc\n",
		},
	}{
		t.Logf("testing case %d\n", i + 1)

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