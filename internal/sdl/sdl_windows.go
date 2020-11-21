package sdl

import (
	"errors"
	"unsafe"

	"golang.org/x/sys/windows"
)

var dll = windows.NewLazyDLL("SDL2.dll")

var (
	SDL_Init              = dll.NewProc("SDL_Init")
	SDL_InitSubSystem     = dll.NewProc("SDL_InitSubSystem")
	SDL_QuitSubSystem     = dll.NewProc("SDL_QuitSubSystem")
	SDL_WasInit           = dll.NewProc("SDL_WasInit")
	SDL_Quit              = dll.NewProc("SDL_Quit")
	SDL_GL_LoadLibrary    = dll.NewProc("SDL_GL_LoadLibrary")
	SDL_GL_GetProcAddress = dll.NewProc("SDL_GL_GetProcAddress")
	SDL_GL_UnloadLibrary  = dll.NewProc("SDL_GL_UnloadLibrary")
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

func GLLoadDefaultLibrary() error {
	errc, _, _ := SDL_GL_LoadLibrary.Call(0)
	if errc != 0 {
		return GetError()
	}
	return nil
}

func GLLoadLibrary(name string) error {
	cname, err := windows.BytePtrFromString(name)
	if err != nil {
		panic(err)
	}
	errc, _, _ := SDL_GL_LoadLibrary.Call(uintptr(unsafe.Pointer(cname)))
	if errc != 0 {
		return GetError()
	}
	return nil
}

func GLGetProcAddress(name string) (error, uintptr) {
	cname, err := windows.BytePtrFromString(name)
	if err != nil {
		panic(err)
	}
	addr, _, _ := SDL_GL_GetProcAddress.Call(uintptr(unsafe.Pointer(cname)))
	if addr == 0 {
		return errors.New("sdl.GLGetProcAddress: unable to find address for " + name), 0
	}
	return nil, addr
}
