package main

import "math"

type numberStat struct {
	l *list
}

func (s numberStat) Count() (int, error) {
	return len(s.l.data), nil
}

func (s numberStat) Average() (float64, error) {
	if len(s.l.data) == 0 {
		return 0.0, ErrEmpty
	}
	var sum int = 0
	for _, e := range s.l.data {
		sum += e.(int)
	}
	return float64(sum) / float64(len(s.l.data)), nil
}

func (s numberStat) Max() (int, error) {
	if len(s.l.data) == 0 {
		return 0.0, ErrEmpty
	}
	var max int = s.l.data[0].(int)
	for _, e := range s.l.data {
		if e.(int) > max {
			max = e.(int)
		}
	}
	return max, nil
}

func (s numberStat) Min() (int, error) {
	if len(s.l.data) == 0 {
		return 0.0, ErrEmpty
	}
	var min int = s.l.data[0].(int)
	for _, e := range s.l.data {
		if e.(int) < min {
			min = e.(int)
		}
	}
	return min, nil
}

func (s numberStat) Mean() (float64, error) {
	return s.Average()
}

func (s numberStat) Std() (float64, error) {
	if len(s.l.data) == 0 {
		return 0.0, ErrEmpty
	}
	if mean, err := s.Mean(); err != nil {
		return 0.0, err
	} else {
		var sd float64
		for _, e := range s.l.data {
			sd += math.Pow(float64(e.(int))-mean, 2)
		}
		sd = math.Sqrt(sd / float64(len(s.l.data)))
		return sd, nil
	}
}
