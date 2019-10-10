package sdl

//#include "sdl.h"
import "C"

func GetError() error {
	if s := C.SDL_GetError(); s != nil {
		return Error(C.GoString(s))
	}
	return nil
}

func ClearError() {
	C.SDL_ClearError()
}
