// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package window

import (
	"fmt"

	"github.com/cozely/platform/internal/sdl"
)

func setup() error {
	if sdl.WasInit(sdl.InitVideo) == 0 {
		err := sdl.Init(sdl.InitVideo)
		if err != nil {
			return err
		}
	}
	return nil
}

// Window represents a platform and its context.
type Window struct {
	handle  sdl.Window
	context sdl.GLContext

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

// New creates a window and its associated context.
func New(o ...Option) (*Window, error) {
	var err error

	err = setup()
	if err != nil {
		return nil, fmt.Errorf("internal.New: %v", err)
	}

	w := Window{
		title: "Untitled",
		size:  Coord{X: 1280, Y: 720},
		debug: true,
	}
	for _, o := range o {
		err := o(&w)
		if err != nil {
			return nil, fmt.Errorf("internal.New: %v", err)
		}
	}

	sdl.GLSetAttribute(sdl.GLContextMajorVersion, 3)
	sdl.GLSetAttribute(sdl.GLContextMinorVersion, 0)
	sdl.GLSetAttribute(sdl.GLContextProfileMask, sdl.GLContextProfileES)
	// sdl.GLSetAttribute(sdl.GLDepthSize, 16)
	sdl.GLSetAttribute(sdl.GLDoubleBuffer, 1)
	multisample := int32(0)
	if multisample > 0 {
		sdl.GLSetAttribute(sdl.GLMultisampleBuffers, 1)
		sdl.GLSetAttribute(sdl.GLMultisampleSamples, multisample)
	}

	if w.debug {
		sdl.GLSetAttribute(sdl.GLContextFlags, sdl.GLContextDebugFlag)
	}

	flags := sdl.WindowOpenGL | sdl.WindowResizable
	if w.fullscreen {
		if w.desktop {
			flags |= sdl.WindowFullscreenDesktop
		} else {
			flags |= sdl.WindowFullscreen
		}
	}

	w.handle, err = sdl.CreateWindow(
		w.title,
		sdl.WindowPosCenteredMask|w.monitor,
		sdl.WindowPosCenteredMask|w.monitor,
		w.size.X,
		w.size.Y,
		flags,
	)
	if err != nil {
		return nil, err
	}

	w.context, err = sdl.GLCreateContext(w.handle)
	if err != nil {
		return nil, err
	}

	if w.vsync {
		err = sdl.GLSetSwapInterval(-1)
		if err != nil {
			err = sdl.GLSetSwapInterval(1)
			if err != nil {
				return &w, err
			}
		}
	} else {
		err = sdl.GLSetSwapInterval(0)
		if err != nil {
			return &w, err
		}
	}

	logOpenGLInfos()

	w.opened = true

	return &w, nil
}

// Present asks the system to display the content of the window (e.g. by
// swapping OpenGL buffers).
func (w *Window) Present() {
	sdl.GLSwapWindow(w.handle)
}

// Close destroys the window.
func (w *Window) Close() {
	sdl.DestroyWindow(w.handle)
}

// HasFocus returns true if the window has focus.
func (w *Window) HasFocus() bool {
	return w.hasFocus
}

// HasMouseFocus returns true if the mouse is currently inside the window.
func (w *Window) HasMouseFocus() bool {
	return w.hasMouseFocus
}

// Size returns the size of the window in (screen) pixels.
func (w *Window) Size() Coord {
	return w.size
}

// logOpenGLInfos displays information about the OpenGL context
func logOpenGLInfos() {
	s := "OpenGL "

	prof, err1 := sdlGLAttribute(sdl.GLContextProfileMask)
	switch {
	case err1 != nil:
		s += "(error) "
	case prof == sdl.GLContextProfileES:
		s += "ES "
	case prof == sdl.GLContextProfileCore:
		s += "Core "
	case prof == sdl.GLContextProfileCompatibility:
		s += "Compatibility "
	default:
		s += "(unknown) "
	}

	maj, err1 := sdlGLAttribute(sdl.GLContextMajorVersion)
	min, err2 := sdlGLAttribute(sdl.GLContextMinorVersion)
	if err1 == nil && err2 == nil {
		s += fmt.Sprintf("%d.%d", maj, min)
	}

	db, err1 := sdlGLAttribute(sdl.GLDoubleBuffer)
	if err1 == nil {
		if db != 0 {
			s += ", double buffer"
		} else {
			s += ", NO double buffer"
		}
	}

	av, err1 := sdlGLAttribute(sdl.GLAcceleratedVisual)
	if err1 == nil {
		if av != 0 {
			s += ", accelerated"
		} else {
			s += ", NOT accelerated"
		}
	}

	sw := sdl.GLGetSwapInterval()
	if sw > 0 {
		if sw != 0 {
			s += ", vsync"
		} else {
			s += ", NO vsync"
		}
	}
	println(s)
}

func sdlGLAttribute(attr sdl.GLAttr) (int32, error) {
	var v int32
	err := sdl.GLGetAttribute(attr, &v)
	if err != nil {
		return 0, err
	}
	return v, nil
}
