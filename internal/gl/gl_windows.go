package gl

import (
	//"errors"
	"syscall"
	"unsafe"

	//"golang.org/x/sys/windows"

	"github.com/cozely/platform/internal/sdl"
)

/*
var (
	opengl32          = windows.NewLazySystemDLL("opengl32")
	wglGetProcAddress = opengl32.NewProc("wglGetProcAddress")
)

func getProcAddress(name string) (error, uintptr) {
	cname, err := windows.BytePtrFromString(name)
	if err != nil {
		panic(err)
	}
	if r, _, _ := wglGetProcAddress.Call(uintptr(unsafe.Pointer(cname))); r != 0 {
		return nil, r
	}
	p := opengl32.NewProc(name)
	if err := p.Find(); err != nil {
		// The proc is not found.
		return errors.New("getProcAddress: no address found for " + name), 0
	}
	return nil, p.Addr()
}
*/

var (
	glClearBufferv uintptr
)

var wasInit bool = false

func WasInit() bool {
	return wasInit
}

func Init() error {
	var err error

	err, glClearBufferv = sdl.GLGetProcAddress("glClearBufferfv")
	if err != nil {
		return err
	}

	wasInit = true

	return nil
}

func ClearBufferv(buffer Enum, drawBuffer int32, color *struct{ R, G, B, A float32 }) {
	syscall.Syscall(glClearBufferv, 3, uintptr(buffer), uintptr(drawBuffer), uintptr(unsafe.Pointer(color)))
}
