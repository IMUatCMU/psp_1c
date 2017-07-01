package main

import (
	"testing"
)

func TestStat(t *testing.T) {
	for _, test := range []struct {
		name        string
		getList     func() List
		assertCount func(t0 *testing.T, count int, err error)
		assertAvg   func(t0 *testing.T, avg float64, err error)
		assertMax   func(t0 *testing.T, max int, err error)
		assertMin   func(t0 *testing.T, min int, err error)
		assertMean  func(t0 *testing.T, mean float64, err error)
		assertStd   func(t0 *testing.T, std float64, err error)
	}{
		{
			name: "string stat normal case",
			getList: func() List {
				l := NewStringList()
				l.Add(0, "a")
				l.Add(1, "ab")
				l.Add(2, "abc")
				return l
			},
			assertCount: func(t0 *testing.T, count int, err error) {
				assertNil(t, err)
				assertEqual(t0, 3, count)
			},
			assertAvg: func(t0 *testing.T, avg float64, err error) {
				assertNil(t, err)
				assertEqual(t0, float64(2.0), avg)
			},
			assertMax: func(t0 *testing.T, max int, err error) {
				assertNil(t, err)
				assertEqual(t0, 3, max)
			},
			assertMin: func(t0 *testing.T, min int, err error) {
				assertNil(t, err)
				assertEqual(t0, 1, min)
			},
			assertMean: func(t0 *testing.T, mean float64, err error) {
				assertEqual(t0, "not_supported", err.Error())
			},
			assertStd: func(t0 *testing.T, std float64, err error) {
				assertEqual(t0, "not_supported", err.Error())
			},
		},
		{
			name: "number stat normal case",
			getList: func() List {
				l := NewNumList()
				l.Add(0, 1)
				l.Add(1, 2)
				l.Add(2, 3)
				return l
			},
			assertCount: func(t0 *testing.T, count int, err error) {
				assertNil(t, err)
				assertEqual(t0, 3, count)
			},
			assertAvg: func(t0 *testing.T, avg float64, err error) {
				assertNil(t, err)
				assertEqual(t0, float64(2.0), avg)
			},
			assertMax: func(t0 *testing.T, max int, err error) {
				assertNil(t, err)
				assertEqual(t0, 3, max)
			},
			assertMin: func(t0 *testing.T, min int, err error) {
				assertNil(t, err)
				assertEqual(t0, 1, min)
			},
			assertMean: func(t0 *testing.T, mean float64, err error) {
				assertNil(t, err)
				assertEqual(t0, float64(2.0), mean)
			},
			assertStd: func(t0 *testing.T, std float64, err error) {
				assertNil(t, err)
				assertEqual(t0, 0.816496580927726, std)
			},
		},
		{
			name: "empty list",
			getList: func() List {
				l := NewNumList()
				return l
			},
			assertCount: func(t0 *testing.T, count int, err error) {
				assertNil(t, err)
				assertEqual(t0, 0, count)
			},
			assertAvg: func(t0 *testing.T, avg float64, err error) {
				assertEqual(t0, "empty", err.Error())
			},
			assertMax: func(t0 *testing.T, max int, err error) {
				assertEqual(t0, "empty", err.Error())
			},
			assertMin: func(t0 *testing.T, min int, err error) {
				assertEqual(t0, "empty", err.Error())
			},
			assertMean: func(t0 *testing.T, mean float64, err error) {
				assertEqual(t0, "empty", err.Error())
			},
			assertStd: func(t0 *testing.T, std float64, err error) {
				assertEqual(t0, "empty", err.Error())
			},
		},
		{
			name: "generic list",
			getList: func() List {
				l := NewList()
				return l
			},
			assertCount: func(t0 *testing.T, count int, err error) {
				assertEqual(t0, "not_supported", err.Error())
			},
			assertAvg: func(t0 *testing.T, avg float64, err error) {
				assertEqual(t0, "not_supported", err.Error())
			},
			assertMax: func(t0 *testing.T, max int, err error) {
				assertEqual(t0, "not_supported", err.Error())
			},
			assertMin: func(t0 *testing.T, min int, err error) {
				assertEqual(t0, "not_supported", err.Error())
			},
			assertMean: func(t0 *testing.T, mean float64, err error) {
				assertEqual(t0, "not_supported", err.Error())
			},
			assertStd: func(t0 *testing.T, std float64, err error) {
				assertEqual(t0, "not_supported", err.Error())
			},
		},
	} {
		t.Run(test.name, func(t0 *testing.T) {
			l := test.getList()

			count, err := l.Count()
			test.assertCount(t0, count, err)

			avg, err := l.Average()
			test.assertAvg(t0, avg, err)

			max, err := l.Max()
			test.assertMax(t0, max, err)

			min, err := l.Min()
			test.assertMin(t0, min, err)

			mean, err := l.Mean()
			test.assertMean(t0, mean, err)

			std, err := l.Std()
			test.assertStd(t0, std, err)
		})
	}
}

func assertNil(t *testing.T, any interface{}) {
	if any != nil {
		t.Errorf("expected nil, but got %v", any)
		t.FailNow()
	}
}

func assertEqual(t *testing.T, expect, got interface{}) {
	if expect != got {
		t.Errorf("expected %v, but got %v", expect, got)
		t.FailNow()
	}
}
