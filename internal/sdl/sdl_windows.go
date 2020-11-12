package sdl

import (
	"golang.org/x/sys/windows"
)

var dll = windows.NewLazyDLL("SDL2.dll")

var (
	SDL_Init          = dll.NewProc("SDL_Init")
	SDL_InitSubSystem = dll.NewProc("SDL_InitSubSystem")
	SDL_QuitSubSystem = dll.NewProc("SDL_QuitSubSystem")
	SDL_WasInit       = dll.NewProc("SDL_WasInit")
	SDL_Quit          = dll.NewProc("SDL_Quit")
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
