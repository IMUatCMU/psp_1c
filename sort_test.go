package main

import (
	"math/rand"
	"os"
	"sort"
	"testing"
)

var (
	n100        []int = make([]int, 100)
	n100Sorted  []int = make([]int, 100)
	n100RSorted []int = make([]int, 100)

	n1000        []int = make([]int, 1000)
	n1000Sorted  []int = make([]int, 1000)
	n1000RSorted []int = make([]int, 1000)

	n10000        []int = make([]int, 10000)
	n10000Sorted  []int = make([]int, 10000)
	n10000RSorted []int = make([]int, 10000)
)

func TestMain(m *testing.M) {
	for i := 0; i < 100; i++ {
		n100[i] = rand.Intn(99999)
	}
	copy(n100Sorted, n100)
	copy(n100RSorted, n100)
	sort.Sort(sort.IntSlice(n100Sorted))
	sort.Sort(sort.Reverse(sort.IntSlice(n100RSorted)))

	for i := 0; i < 1000; i++ {
		n1000[i] = rand.Intn(99999)
	}
	copy(n1000Sorted, n1000)
	copy(n1000RSorted, n1000)
	sort.Sort(sort.IntSlice(n1000Sorted))
	sort.Sort(sort.Reverse(sort.IntSlice(n1000RSorted)))

	for i := 0; i < 10000; i++ {
		n10000[i] = rand.Intn(99999)
	}
	copy(n10000Sorted, n10000)
	copy(n10000RSorted, n10000)
	sort.Sort(sort.IntSlice(n10000Sorted))
	sort.Sort(sort.Reverse(sort.IntSlice(n10000RSorted)))

	code := m.Run()
	os.Exit(code)
}

func TestAdd100ToList(t *testing.T) {
	l := NewNumList()
	for i := 0; i < 100; i++ {
		assertNil(t, l.Add(i, i))
	}
}

func TestSortOnEmptyList(t *testing.T) {
	l := NewNumList()
	assertEqual(t, "empty", l.Sort(true).Error())
}

func TestAddRealValuesToList(t *testing.T) {
	l := NewNumList()
	for i, n := range []interface{}{1, 2.1, 3.0} {
		l.Add(i, n)
	}
	for i, n := range []int{1, 2, 3} {
		assertEqual(t, n, l.(*list).data[i].(int))
	}
}

func TestQuickSort(t *testing.T) {
	for _, test := range []struct {
		name    string
		getList func() List
		asc     bool
		assert  func(t0 *testing.T, l List)
	}{
		// n=100
		{
			name: "sort 100 random asc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n100 {
					l.Add(i, n)
				}
				return l
			},
			asc: true,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 100; i++ {
					assertEqual(t0, n100Sorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 100 sorted asc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n100Sorted {
					l.Add(i, n)
				}
				return l
			},
			asc: true,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 100; i++ {
					assertEqual(t0, n100Sorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 100 reverse sorted asc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n100RSorted {
					l.Add(i, n)
				}
				return l
			},
			asc: true,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 100; i++ {
					assertEqual(t0, n100Sorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 100 random desc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n100 {
					l.Add(i, n)
				}
				return l
			},
			asc: false,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 100; i++ {
					assertEqual(t0, n100RSorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 100 sorted desc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n100Sorted {
					l.Add(i, n)
				}
				return l
			},
			asc: false,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 100; i++ {
					assertEqual(t0, n100RSorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 100 reverse sorted desc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n100RSorted {
					l.Add(i, n)
				}
				return l
			},
			asc: false,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 100; i++ {
					assertEqual(t0, n100RSorted[i], l.(*list).data[i].(int))
				}
			},
		},
		// n=1000
		{
			name: "sort 1000 random asc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n1000 {
					l.Add(i, n)
				}
				return l
			},
			asc: true,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 1000; i++ {
					assertEqual(t0, n1000Sorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 1000 sorted asc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n1000Sorted {
					l.Add(i, n)
				}
				return l
			},
			asc: true,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 1000; i++ {
					assertEqual(t0, n1000Sorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 1000 reverse sorted asc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n1000RSorted {
					l.Add(i, n)
				}
				return l
			},
			asc: true,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 1000; i++ {
					assertEqual(t0, n1000Sorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 1000 random desc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n1000 {
					l.Add(i, n)
				}
				return l
			},
			asc: false,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 1000; i++ {
					assertEqual(t0, n1000RSorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 1000 sorted desc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n1000Sorted {
					l.Add(i, n)
				}
				return l
			},
			asc: false,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 1000; i++ {
					assertEqual(t0, n1000RSorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 1000 reverse sorted desc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n1000RSorted {
					l.Add(i, n)
				}
				return l
			},
			asc: false,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 1000; i++ {
					assertEqual(t0, n1000RSorted[i], l.(*list).data[i].(int))
				}
			},
		},
		// n=10000
		{
			name: "sort 10000 random asc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n10000 {
					l.Add(i, n)
				}
				return l
			},
			asc: true,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 10000; i++ {
					assertEqual(t0, n10000Sorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 10000 sorted asc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n10000Sorted {
					l.Add(i, n)
				}
				return l
			},
			asc: true,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 10000; i++ {
					assertEqual(t0, n10000Sorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 10000 reverse sorted asc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n10000RSorted {
					l.Add(i, n)
				}
				return l
			},
			asc: true,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 10000; i++ {
					assertEqual(t0, n10000Sorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 10000 random desc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n10000 {
					l.Add(i, n)
				}
				return l
			},
			asc: false,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 10000; i++ {
					assertEqual(t0, n10000RSorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 10000 sorted desc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n10000Sorted {
					l.Add(i, n)
				}
				return l
			},
			asc: false,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 10000; i++ {
					assertEqual(t0, n10000RSorted[i], l.(*list).data[i].(int))
				}
			},
		},
		{
			name: "sort 10000 reverse sorted desc",
			getList: func() List {
				l := NewNumList()
				for i, n := range n10000RSorted {
					l.Add(i, n)
				}
				return l
			},
			asc: false,
			assert: func(t0 *testing.T, l List) {
				for i := 0; i < 10000; i++ {
					assertEqual(t0, n10000RSorted[i], l.(*list).data[i].(int))
				}
			},
		},
	} {
		t.Run(test.name, func(t0 *testing.T) {
			l := test.getList()
			err := l.Sort(test.asc)
			if err != nil {
				t.Error("sort error")
				t.FailNow()
			}
			test.assert(t0, l)
		})
	}
}
