// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package window

import (
	"fmt"

	"github.com/cozely/platform/internal"
)

// Window represents a platform and its context.
type Window struct {
	internal internal.Window

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

	err = internal.Setup()
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

	w.internal, err = internal.NewWindow(
		w.title,
		w.size.X, w.size.Y,
		0,
		w.monitor,
		w.fullscreen,
		w.desktop,
		w.vsync,
		w.debug,
	)

	w.opened = true
	//logOpenGLInfos()
	return &w, nil
}

// Present asks the system to display the content of the window (e.g. by
// swapping OpenGL buffers).
func (w *Window) Present() {
	w.internal.Present()
}

// Close destroys the window.
func (w *Window) Close() {
	w.internal.Close()
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
