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

func GetString(device Device, param Enum) string {
	return C.GoString(C.alcGetString(device.CDevice(), C.ALCenum(param)))
}

func GetIntegerv(device Device, param Enum, size Sizei) []Int {
	data := make([]Int, size)
	C.alcGetIntegerv(device.CDevice(), C.ALCenum(param), C.ALCsizei(size), (*C.ALCint)(&data[0]))
	return data
}

// TODO: Add functions for known params
