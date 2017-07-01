package main

type noStat struct{}

func (s noStat) Count() (int, error) {
	return 0, ErrNotSupported
}

func (s noStat) Average() (float64, error) {
	return 0.0, ErrNotSupported
}

func (s noStat) Max() (int, error) {
	return 0, ErrNotSupported
}

func (s noStat) Min() (int, error) {
	return 0, ErrNotSupported
}

func (s noStat) Mean() (float64, error) {
	return 0.0, ErrNotSupported
}

func (s noStat) Std() (float64, error) {
	return 0.0, ErrNotSupported
}
