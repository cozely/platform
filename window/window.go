// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package window

/*
#cgo windows LDFLAGS: -lSDL2
#cgo linux freebsd darwin pkg-config: sdl2

#include "sdl.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

func sdlInit() error {
	if C.SDL_WasInit(C.SDL_INIT_VIDEO) == 0 {
		cerr := C.SDL_Init(C.SDL_INIT_VIDEO)
		if cerr != 0 {
			return fmt.Errorf("window.OpenWindow: failed to initialize SDL: %v", sdlError())
		}
	}
	return nil
}

// Window represents a system window and its context.
type Window struct {
	window  *C.SDL_Window
	context C.SDL_GLContext

	title         string
	size          Coord
	monitor       int32
	multisample   int32
	debug         bool
	vsync         bool
	fullscreen    bool
	desktop       bool
	hasFocus      bool
	hasMouseFocus bool
	opened        bool
}

// New creates the game window and its associated OpenGL context.
func New(o ...Option) (*Window, error) {
	var err error

	err = sdlInit()
	if err != nil {
		return nil, fmt.Errorf("window.New: %v", err)
	}

	w := Window{
		title: "Untitled",
		size:  Coord{X: 1280, Y: 720},
		debug: true,
	}
	for _, o := range o {
		err := o(&w)
		if err != nil {
			return nil, fmt.Errorf("window.New: %v", err)
		}
	}
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MAJOR_VERSION, 4)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_MINOR_VERSION, 6)
	C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_PROFILE_MASK, C.SDL_GL_CONTEXT_PROFILE_CORE)
	// C.SDL_GL_SetAttribute(C.SDL_GL_DEPTH_SIZE, 16)
	C.SDL_GL_SetAttribute(C.SDL_GL_DOUBLEBUFFER, 1)
	if /*Window.Multisample > 0*/ false {
		C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLEBUFFERS, 1)
		C.SDL_GL_SetAttribute(C.SDL_GL_MULTISAMPLESAMPLES, C.int(w.multisample))
	}

	if w.debug {
		C.SDL_GL_SetAttribute(C.SDL_GL_CONTEXT_FLAGS, C.SDL_GL_CONTEXT_DEBUG_FLAG)
	}

	t := C.CString(w.title)
	defer C.free(unsafe.Pointer(t))

	var fs uint32
	if w.fullscreen {
		if w.desktop {
			fs = C.SDL_WINDOW_FULLSCREEN_DESKTOP
		} else {
			fs = C.SDL_WINDOW_FULLSCREEN
		}
	}
	fl := C.SDL_WINDOW_OPENGL | C.SDL_WINDOW_RESIZABLE | C.Uint32(fs)

	w.window = C.SDL_CreateWindow(
		t,
		C.int(C.SDL_WINDOWPOS_CENTERED_MASK|w.monitor),
		C.int(C.SDL_WINDOWPOS_CENTERED_MASK|w.monitor),
		C.int(w.size.X),
		C.int(w.size.Y),
		fl,
	)
	if w.window == nil {
		return nil, fmt.Errorf("window.New: %v", sdlError())
	}

	w.context = C.SDL_GL_CreateContext(w.window)
	if w.context == nil {
		return nil, fmt.Errorf("window.New: %s", sdlError())
	}

	if w.vsync {
		cerr := C.SDL_GL_SetSwapInterval(-1)
		if cerr != 0 {
			cerr := C.SDL_GL_SetSwapInterval(1)
			if cerr != 0 {
				return nil, fmt.Errorf("window.New: %v", sdlError())
			}
		}
	} else {
		cerr := C.SDL_GL_SetSwapInterval(0)
		if cerr != 0 {
			return nil, fmt.Errorf("window.New: %v", sdlError())
		}
	}

	w.opened = true
	//logOpenGLInfos()
	return &w, nil
}

// Present asks the system to display the content of the window (e.g. by
// swapping OpenGL buffers).
func (w *Window) Present() {
	C.SDL_GL_SwapWindow(w.window)
}

// Close destroys the window
func (w *Window) Close() {
	C.SDL_DestroyWindow(w.window)
}

// HasFocus returns true if the game windows has focus.
func (w *Window) HasFocus() bool {
	return w.hasFocus
}

// HasMouseFocus returns true if the mouse is currently inside the game window.
func (w *Window) HasMouseFocus() bool {
	return w.hasMouseFocus
}

var size Coord

// Size returns the size of the window in (screen) pixels.
func (w *Window) Size() Coord {
	return w.size
}

// sdlError returns nil or the current SDL Error wrapped in a Go error.
func sdlError() error {
	if s := C.SDL_GetError(); s != nil {
		return errors.New(C.GoString(s))
	}
	return nil
}

type Option func(*Window) error

func Title(s string) Option {
	return func(w *Window) error {
		w.title = s
		if w.opened {
			cs := C.CString(s)
			defer C.free(unsafe.Pointer(cs))
			C.SDL_SetWindowTitle(w.window, cs)
		}
		return nil
	}
}

func Size(x, y int32) Option {
	return func(w *Window) error {
		if w.opened {
			//TODO: implement
			return errors.New("window.VSync: not implemented for opened windows")
		}
		w.size = Coord{x, y}
		return nil
	}
}

func Fullscreen(fullscreen bool, windowed bool) Option {
	return func(w *Window) error {
		if w.opened {
			//TODO: implement
			return errors.New("window.VSync: not implemented for opened windows")
		}
		w.fullscreen = fullscreen
		w.desktop = !windowed
		return nil
	}
}

func Monitor(n int) Option {
	return func(w *Window) error {
		if w.opened {
			return errors.New("window.VSync: not implemented for opened windows")
		}
		w.monitor = int32(n)
		return nil
	}
}

func VSync(enable bool) Option {
	return func(w *Window) error {
		if w.opened {
			return errors.New("window.VSync: not implemented for opened windows")
		}
		w.vsync = enable
		return nil
	}
}

func Debug(enable bool) Option {
	return func(w *Window) error {
		if w.opened {
			return errors.New("window.Debug: not implemented for opened windows")
		}
		w.debug = enable
		return nil
	}
}
