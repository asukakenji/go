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

// Boolean represents a 8-bit boolean value.
type Boolean int8

// Char represents a character.
type Char int8

// Byte represents a signed 8-bit 2's complement integer.
type Byte int8

// Ubyte represents an unsigned 8-bit integer.
type Ubyte uint8

// Short represents a signed 16-bit 2's complement integer.
type Short int16

// Ushort represents an unsigned 16-bit integer.
type Ushort uint16

// Int represents a signed 32-bit 2's complement integer.
type Int int32

// Uint represents an unsigned 32-bit integer.
type Uint uint32

// Sizei represents a non-negative 32-bit binary integer size.
type Sizei int32

// Enum represents an enumerated 32-bit value.
type Enum int32

// Float represents a 32-bit IEEE754 floating-point value.
type Float float32

// Double represents a 64-bit IEEE754 floating-point value.
type Double float64

const (
	// False represents Boolean False.
	False Boolean = 0
	// True represents Boolean True.
	True Boolean = 1
)

const (
	// Frequency represents output frequency. It is specified in <int> Hz.
	Frequency Int = C.ALC_FREQUENCY
	// Refresh represents update rate of context processing. It is specified in <int> Hz.
	Refresh Int = C.ALC_REFRESH
	// Sync is a flag indicating a synchronous context. It is specified using a bool value.
	Sync Int = C.ALC_SYNC
	// MonoSources represents the requested number of mono sources. It is specified by <int> number of requested mono (3D) sources.
	MonoSources Int = C.ALC_MONO_SOURCES
	// StereoSources represents the requested number of stereo sources. It is specified by <int> number of requested stereo sources.
	StereoSources Int = C.ALC_STEREO_SOURCES
)

const (
	// NoError indicates there is not currently an error.
	NoError Enum = C.ALC_NO_ERROR
	// InvalidDevice indicates a bad device was passed to an OpenAL function.
	InvalidDevice Enum = C.ALC_INVALID_DEVICE
	// InvalidContext indicates a bad context was passed to an OpenAL function.
	InvalidContext Enum = C.ALC_INVALID_CONTEXT
	// InvalidEnum indicates an unknown enum value was passed to an OpenAL function.
	InvalidEnum Enum = C.ALC_INVALID_ENUM
	// InvalidValue indicate an invalid value was passed to an OpenAL function.
	InvalidValue Enum = C.ALC_INVALID_VALUE
	// OutOfMemory indicates the requested operation resulted in OpenAL running out of memory.
	OutOfMemory Enum = C.ALC_OUT_OF_MEMORY
)

const (
	DefaultDeviceSpecifier        Enum = C.ALC_DEFAULT_DEVICE_SPECIFIER
	DeviceSpecifier               Enum = C.ALC_DEVICE_SPECIFIER
	Extensions                    Enum = C.ALC_EXTENSIONS
	CaptureDeviceSpecifier        Enum = C.ALC_CAPTURE_DEVICE_SPECIFIER
	CaptureDefaultDeviceSpecifier Enum = C.ALC_CAPTURE_DEFAULT_DEVICE_SPECIFIER
)

const (
	MajorVersion   Enum = C.ALC_MAJOR_VERSION
	MinorVersion   Enum = C.ALC_MINOR_VERSION
	AttributesSize Enum = C.ALC_ATTRIBUTES_SIZE
	AllAttributes  Enum = C.ALC_ALL_ATTRIBUTES
)

// GoBool : C bool to Go bool
func GoBool(t C.ALCboolean) bool {
	if t == C.ALC_FALSE {
		return false
	}
	return true
}

// CBool : Go bool to C bool
func CBool(t bool) C.ALCboolean {
	if t {
		return C.ALC_TRUE
	}
	return C.ALC_FALSE
}
