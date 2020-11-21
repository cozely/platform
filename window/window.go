// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package window

import (
	"fmt"

	"github.com/cozely/platform/internal/gl"
	"github.com/cozely/platform/internal/sdl"
)

func setupSDL() error {
	if sdl.WasInit(sdl.InitVideo) == 0 {
		err := sdl.Init(sdl.InitVideo)
		if err != nil {
			return err
		}
	}

	sdl.GLLoadDefaultLibrary()
	return nil
}

func setupGL() error {
	if !gl.WasInit() {
		err := gl.Init() //TODO: test if already init
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

	err = setupSDL()
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

	err = setupGL()
	if err != nil {
		return nil, err
	}

	c := struct {
		R float32
		G float32
		B float32
		A float32
	}{R: 1.0, G: 0.5, B: 0.5, A: 1.0}
	gl.ClearBufferv(gl.COLOR, 0, &c)

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
