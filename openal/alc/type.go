package alc

// #cgo darwin  LDFLAGS: -framework OpenAL
// #cgo freebsd LDFLAGS: -lopenal
//
// #ifdef __APPLE__
// #include <OpenAL/al.h>
// #include <OpenAL/alc.h>
// #else
// #include <AL/al.h>
// #include <AL/alc.h>
// #endif
import "C"

type Boolean int8
type Char int8
type Byte int8
type Ubyte uint8
type Short int16
type Ushort uint16
type Int int32
type Uint uint32
type Sizei int32
type Enum int32
type Float float32
type Double float64

const (
	False Boolean = 0
	True  Boolean = 1
)

const (
	Frequency     Int = C.ALC_FREQUENCY
	Refresh       Int = C.ALC_REFRESH
	Sync          Int = C.ALC_SYNC
	MonoSources   Int = C.ALC_MONO_SOURCES
	StereoSources Int = C.ALC_STEREO_SOURCES
)

const (
	NoError        Enum = C.ALC_NO_ERROR
	InvalidDevice  Enum = C.ALC_INVALID_DEVICE
	InvalidContext Enum = C.ALC_INVALID_CONTEXT
	InvalidEnum    Enum = C.ALC_INVALID_ENUM
	InvalidValue   Enum = C.ALC_INVALID_VALUE
	OutOfMemory    Enum = C.ALC_OUT_OF_MEMORY
)

func GoBool(t C.ALCboolean) bool {
	if t == C.ALC_FALSE {
		return false
	}
	return true
}

func CBool(t bool) C.ALCboolean {
	if t {
		return C.ALC_TRUE
	}
	return C.ALC_FALSE
}

func GoEnum(e C.ALCenum) Enum {
	return Enum(e)
}

func CEnum(e Enum) C.ALCenum {
	return C.ALCenum(e)
}
