// +build !windows

package sdl

//#include "sdl.h"
import "C"

import "unsafe"

func GLSetAttribute(a GLAttr, value int32) error {
	errc := C.SDL_GL_SetAttribute(C.SDL_GLattr(a), C.int(value))
	if errc != 0 {
		return Error("SDL: GL attribute could not be set")
	}

	return nil
}

func GLGetAttribute(a GLAttr, values *int32) error {
	errc := C.SDL_GL_GetAttribute(C.SDL_GLattr(a), (*C.int)(values))
	if errc != 0 {
		return GetError()
	}
	return nil
}

func CreateWindow(title string, x, y int32, w, h int32, f WindowFlags) (Window, error) {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))

	win := C.SDL_CreateWindow(t, C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(f))
	if win == nil {
		return Window{0}, GetError()
	}
	return Window{uintptr(unsafe.Pointer(win))}, nil
}

func GLCreateContext(w Window) (GLContext, error) {
	c := C.SDL_GL_CreateContext((*C.SDL_Window)(unsafe.Pointer(w.uintptr)))
	if c == nil {
		return GLContext{0}, GetError()
	}
	return GLContext{uintptr(unsafe.Pointer(c))}, nil
}

func GLSetSwapInterval(interval int32) error {
	errc := C.SDL_GL_SetSwapInterval(C.int(interval))
	if errc != 0 {
		return GetError()
	}
	return nil
}

func GLGetSwapInterval() int32 {
	return int32(C.SDL_GL_GetSwapInterval())
}

func GLSwapWindow(w Window) {
	C.SDL_GL_SwapWindow((*C.SDL_Window)(unsafe.Pointer(w.uintptr)))
}

func DestroyWindow(w Window) {
	C.SDL_DestroyWindow((*C.SDL_Window)(unsafe.Pointer(w.uintptr)))
}
