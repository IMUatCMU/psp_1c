package main

type stringStat struct {
	l *list
}

func (s stringStat) Count() (int, error) {
	return len(s.l.data), nil
}

func (s stringStat) Average() (float64, error) {
	if len(s.l.data) == 0 {
		return 0.0, ErrEmpty
	}
	var sum int = 0
	for _, e := range s.l.data {
		sum += len(e.(string))
	}
	return float64(sum) / float64(len(s.l.data)), nil
}

func (s stringStat) Max() (int, error) {
	if len(s.l.data) == 0 {
		return 0.0, ErrEmpty
	}
	var max int = len(s.l.data[0].(string))
	for _, e := range s.l.data {
		if len(e.(string)) > max {
			max = len(e.(string))
		}
	}
	return max, nil
}

func (s stringStat) Min() (int, error) {
	if len(s.l.data) == 0 {
		return 0.0, ErrEmpty
	}
	var min int = len(s.l.data[0].(string))
	for _, e := range s.l.data {
		if len(e.(string)) < min {
			min = len(e.(string))
		}
	}
	return min, nil
}

func (s stringStat) Mean() (float64, error) {
	return 0.0, ErrNotSupported
}

func (s stringStat) Std() (float64, error) {
	return 0.0, ErrNotSupported
}
