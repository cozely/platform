package window

import (
	"fmt"

	"github.com/cozely/platform/internal/sdl"
)

// InfoString returns information about the OpenGL context
func InfoString() string {
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

	return s
}

func sdlGLAttribute(attr sdl.GLAttr) (int32, error) {
	var v int32
	err := sdl.GLGetAttribute(attr, &v)
	if err != nil {
		return 0, err
	}
	return v, nil
}
