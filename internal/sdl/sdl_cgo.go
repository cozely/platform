// +build !windows

package sdl

/*
#cgo windows LDFLAGS: -lSDL2
#cgo linux freebsd darwin pkg-config: sdl2

#include "sdl.h"
*/
import "C"

func Init(f InitFlags) error {
	errc := C.SDL_Init(C.Uint32(f))
	if errc != 0 {
		return GetError()
	}
	return nil
}

func InitSubSystem(f InitFlags) error {
	errc := C.SDL_InitSubSystem(C.Uint32(f))
	if errc != 0 {
		return GetError()
	}
	return nil
}

func QuitSubSystem(f InitFlags) {
	C.SDL_QuitSubSystem(C.Uint32(f))
}

func WasInit(f InitFlags) InitFlags {
	return InitFlags(C.SDL_WasInit(C.Uint32(f)))
}

func Quit() {
	C.SDL_Quit()
}
