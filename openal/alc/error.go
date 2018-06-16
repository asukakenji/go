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
//
// // free()
// #include <stdlib.h>
import "C"

// GetError retrieves the current context error state.
func GetError(d Device) Enum {
	return GoEnum(C.alcGetError(d.CDevice()))
}
