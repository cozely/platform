package sdl

type pointer struct {
	uintptr
}

type Error string

func (s Error) Error() string {
	return string(s)
}
