package sdl

type InitFlags uint32

const (
	InitTimer          InitFlags = 0x00000001
	InitAudio                    = 0x00000010
	InitVideo                    = 0x00000020 // InitVideo implies InitEvents
	InitJoystick                 = 0x00000200 // InitJoystick implies InitEvents
	InitHaptic                   = 0x00001000
	InitGamecontroller           = 0x00002000 // InitGameController implies InitJoystick
	InitEvents                   = 0x00004000
	InitSensor                   = 0x00008000
	InitEverything               = InitTimer | InitAudio | InitVideo | InitEvents | InitJoystick | InitHaptic | InitGamecontroller | InitSensor
)
