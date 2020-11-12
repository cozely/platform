package sdl

import "unsafe"

var (
	SDL_GL_SetAttribute    = dll.NewProc("SDL_GL_SetAttribute")
	SDL_GL_GetAttribute    = dll.NewProc("SDL_GL_GetAttribute")
	SDL_GL_CreateContext   = dll.NewProc("SDL_GL_CreateContext")
	SDL_GL_SetSwapInterval = dll.NewProc("SDL_GL_SetSwapInterval")
	SDL_GL_GetSwapInterval = dll.NewProc("SDL_GL_GetSwapInterval")
	SDL_GL_SwapWindow      = dll.NewProc("SDL_GL_SwapWindow")
	SDL_CreateWindow       = dll.NewProc("SDL_CreateWindow")
	SDL_DestroyWindow      = dll.NewProc("SDL_DestroyWindow")
)

func GLSetAttribute(a GLAttr, value int32) error {
	errc, _, _ := SDL_GL_SetAttribute.Call(uintptr(a), uintptr(value))
	if errc != 0 {
		return Error("SDL: GL attribute could not be set")
	}
	return nil
}

func GLGetAttribute(a GLAttr, values *int32) error {
	errc, _, _ := SDL_GL_GetAttribute.Call(uintptr(a), uintptr(unsafe.Pointer(values)))
	if errc != 0 {
		return GetError()
	}
	return nil
}

func CreateWindow(title string, x, y int32, w, h int32, f WindowFlags) (Window, error) {
	t := cString(title)

	win, _, _ := SDL_CreateWindow.Call(t, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(f))
	if win == 0 {
		return Window{0}, GetError()
	}
	return Window{uintptr(unsafe.Pointer(win))}, nil
}

func GLCreateContext(w Window) (GLContext, error) {
	c, _, _ := SDL_GL_CreateContext.Call(w.uintptr)
	if c == 0 {
		return GLContext{0}, GetError()
	}
	return GLContext{c}, nil
}

func GLSetSwapInterval(interval int32) error {
	errc, _, _ := SDL_GL_SetSwapInterval.Call(uintptr(interval))
	if errc != 0 {
		return GetError()
	}
	return nil
}

func GLGetSwapInterval() int32 {
	r, _, _ := SDL_GL_GetSwapInterval.Call()
	return int32(r)
}

func GLSwapWindow(w Window) {
	SDL_GL_SwapWindow.Call(w.uintptr)
}

func DestroyWindow(w Window) {
	SDL_DestroyWindow.Call(w.uintptr)
}
