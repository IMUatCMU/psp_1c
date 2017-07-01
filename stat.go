package main

import "errors"

type Stat interface {
	Count() (int, error)
	Average() (float64, error)
	Max() (int, error)
	Min() (int, error)
	Mean() (float64, error)
	Std() (float64, error)
}

var (
	ErrEmpty        = errors.New("empty")
	ErrNotSupported = errors.New("not_supported")
)
