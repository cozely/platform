package sdl

// The structure that defines a display mode.
type DisplayMode struct {
	Format      uint32  // pixel format
	W           int32   // in screen coordinates
	H           int32   // in screen coordinates
	RefreshRate int32   // zero for unspecified
	DriverData  pointer // driver-specific, initialize to 0
}

// The type used to identify a window
type Window pointer

// The flags on a window.
type WindowFlags uint32

const (
	WindowFullscreen        WindowFlags = 0x00000001
	WindowOpenGL            WindowFlags = 0x00000002
	WindowShown             WindowFlags = 0X00000004
	WindowHidden            WindowFlags = 0x00000008
	WindowBorderless        WindowFlags = 0x00000010
	WindowResizable         WindowFlags = 0x00000020
	WindowMinimized         WindowFlags = 0x00000040
	WindowMaximized         WindowFlags = 0x00000080
	WindowInputGrabbed      WindowFlags = 0x00000100
	WindowInputFocus        WindowFlags = 0x00000200
	WindowMouseFocus        WindowFlags = 0x00000400
	WindowFullscreenDesktop WindowFlags = (WindowFullscreen | 0x00001000)
	WindowForeign           WindowFlags = 0x00000800 // i.e. not created by SDL
	WindowAllowHighdpi      WindowFlags = 0x00002000 // should be created in high-DPI mode if supported
	WindowMouseCapture      WindowFlags = 0x00004000 // has mouse captured (unrelated to InputGrabbed)
	WindowAlwaysOnTop       WindowFlags = 0x00008000
	WindowSkipTaskbar       WindowFlags = 0x00010000
	WindowUtility           WindowFlags = 0x00020000
	WindowTooltip           WindowFlags = 0x00040000
	WindowPopupMenu         WindowFlags = 0x00080000
	WindowVulkan            WindowFlags = 0x10000000
)

const (
	WindowPosUndefinedMask int32 = 0x1FFF0000
	WindowPosUndefined           = WindowPosUndefinedMask | 0
	WindowPosCenteredMask        = 0x2FFF0000
	WindowPosCentered            = WindowPosCenteredMask | 0
)

func WindowPosUndefinedDisplay(d int32) int32 {
	return WindowPosUndefinedMask | d
}

func WindowPosCenteredDisplay(d int32) int32 {
	return WindowPosCenteredMask | d
}

// An opaque handle to an OpenGL context.
type GLContext pointer

type GLAttr int32

const (
	GLRedSize GLAttr = iota
	GLGreenSize
	GLBlueSize
	GLAlphaSize
	GLBufferSize
	GLDoubleBuffer
	GLDepthSize
	GLStencilSize
	GLAccumRedSize
	GLAccumGreenSize
	GLAccumBlueSize
	GLAccumAlphaSize
	GLStereo
	GLMultisampleBuffers
	GLMultisampleSamples
	GLAcceleratedVisual
	GLRetainedBacking
	GLContextMajorVersion
	GLContextMinorVersion
	GLContextEGL
	GLContextFlags
	GLContextProfileMask
	GLShareWithCurrentContext
	GLFramebufferSRGBCapable
	GLContextReleaseBehavior
	GLContextResetNotification
	GLContextNoError
)

const (
	GLContextProfileCore          int32 = 0x0001
	GLContextProfileCompatibility int32 = 0x0002
	GLContextProfileES            int32 = 0x0004 // GLX_CONTEXT_ES2_PROFILE_BIT_EXT
)

const (
	GLContextDebugFlag             int32 = 0x0001
	GLContextForwardCompatibleFlag int32 = 0x0002
	GLContextRobustAccessFlag      int32 = 0x0004
	GLContextResetIsolationFlag    int32 = 0x0008
)
