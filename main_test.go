package main

import (
	"fmt"
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
