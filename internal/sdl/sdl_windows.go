package sdl

import (
	"unsafe"
	"x/sys/syscall"
)

var dll = syscall.NewLazyDLL("SDL2.dll")

var (
	SDL_Init          = dll.NewProc("SDL_Init")
	SDL_InitSubSystem = dll.NewProc("SDL_InitSubSystem")
)

func Init(f InitFlags) error {
	errc, _, _ := SDL_Init.Call(uintptr(f))
	if errc != 0 {
		return GetError()
	}
	return nil
}

func InitSubSystem(f InitFlags) error {
	errc, _, _ := SDL_InitSubSystem.Call(uintptr(f))
	if errc != 0 {
		return GetError()
	}
	return nil
}

func QuitSubSystem(f InitFlags) {
	SDL_QuitSubSystem.Call(uintptr(f))
}

func WasInit(f InitFlags) InitFlags {
	ret, _, _ := SDL_WasInit.Call(uintptr(f))
	return InitFlags(ret)
}

func Quit() {
	SDL_Quit.Call()
}
