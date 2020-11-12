package sdl

import (
	"bytes"
)

var (
	SDL_GetError   = dll.NewProc("SDL_GetError")
	SDL_ClearError = dll.NewProc("SDL_ClearError")
)

func goString(uintptr p) string {
	b := (*[1 << (30 - 1)]byte)(unsafe.Pointer(p))
	size = bytes.IndexByte(b, 0)
	return string(b[:size:size])
}

func GetError() error {
	s, _, _ := SDL_GetError.Call()
	if s != 0 {
		return Error(goString(s))
	}
	return nil
}

func ClearError() {
	SDL_ClearError.Call()
}
