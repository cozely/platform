package internal

import (
	"errors"
	"fmt"
	"unsafe"
)

/*
#cgo windows LDFLAGS: -lSDL2
#cgo linux freebsd darwin pkg-config: sdl2

#if defined(__WIN32)
#include <SDL2/SDL.h>
#else
#include <SDL.h>
#endif
*/
import "C"

func Setup() error {
	if C.SDL_WasInit(C.SDL_INIT_VIDEO) == 0 {
		cerr := C.SDL_Init(C.SDL_INIT_VIDEO)
		if cerr != 0 {
			return fmt.Errorf("platform.internal.Setup: failed to initialize SDL: %v", sdlError())
		}
	}
	return nil
}

// sdlError returns nil or the current SDL Error wrapped in a Go error.
func sdlError() error {
	if s := C.SDL_GetError(); s != nil {
		return errors.New(C.GoString(s))
	}
	return nil
}

type Window struct {
	handle  *C.SDL_Window
	context C.SDL_GLContext
}

func NewWindow(title string, w, h int32, multisample int32, monitor int32, fullscreen, desktop, vsync, debug bool) (win Window, err error) {

	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MAJOR_VERSION, 4)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MINOR_VERSION, 6)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_PROFILE_MASK, C.SDL_GL_CONTEXT_PROFILE_CORE)
	// C.SDL_GL_SetAttribute(C.SDL_GL_DEPTH_SIZE, 16)
	C.SDL_GL_SetAttribute(C.SDL_GL_DOUBLEBUFFER, 1)
	if multisample > 0 {
		C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLEBUFFERS, 1)
		C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLESAMPLES, C.int(multisample))
	}

	if debug {
		C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_FLAGS, C.SDL_GL_CONTEXT_DEBUG_FLAG)
	}

	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))

	var fs uint32
	if fullscreen {
		if desktop {
			fs = C.SDL_WINDOW_FULLSCREEN_DESKTOP
		} else {
			fs = C.SDL_WINDOW_FULLSCREEN
		}
	}
	fl := C.SDL_WINDOW_OPENGL | C.SDL_WINDOW_RESIZABLE | C.Uint32(fs)

	win.handle = C.SDL_CreateWindow(
		t,
		C.int(C.SDL_WINDOWPOS_CENTERED_MASK|monitor),
		C.int(C.SDL_WINDOWPOS_CENTERED_MASK|monitor),
		C.int(w),
		C.int(h),
		fl,
	)
	if win.handle == nil {
		return win, fmt.Errorf("platform.internal.NewWindow: %v", sdlError())
	}

	win.context = C.SDL_GL_CreateContext(win.handle)
	if win.context == nil {
		return win, fmt.Errorf("platform.internal.NewWindow: %s", sdlError())
	}

	if vsync {
		cerr := C.SDL_GL_SetSwapInterval(-1)
		if cerr != 0 {
			cerr := C.SDL_GL_SetSwapInterval(1)
			if cerr != 0 {
				return win, fmt.Errorf("platform.internal.NewWindow: %v", sdlError())
			}
		}
	} else {
		cerr := C.SDL_GL_SetSwapInterval(0)
		if cerr != 0 {
			return win, fmt.Errorf("platform.internal.NewWindow: %v", sdlError())
		}
	}

	return win, nil
}

func (w Window) Present() {
	C.SDL_GL_SwapWindow(w.handle)
}

func (w Window) Close() {
	C.SDL_DestroyWindow(w.handle)
}
