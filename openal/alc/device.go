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
import "unsafe"

// Device represents an OpenAL device.
type Device uintptr

// OpenDevice opens a device by name.
func OpenDevice(deviceName string) Device {
	cDeviceName := C.CString(deviceName)
	defer C.free(unsafe.Pointer(cDeviceName))
	return GoDevice(C.alcOpenDevice(cDeviceName))
}

// OpenDefaultDevice opens the default device.
func OpenDefaultDevice() Device {
	return GoDevice(C.alcOpenDevice(nil))
}

// Close closes device d.
func (d Device) Close() bool {
	return GoBool(C.alcCloseDevice(d.CDevice()))
}

// GoDevice : C device to Go device.
func GoDevice(cDevice *C.ALCdevice) Device {
	return Device(unsafe.Pointer(cDevice))
}

// CDevice : Go device to C device.
func (d Device) CDevice() *C.ALCdevice {
	return (*C.ALCdevice)(unsafe.Pointer(d))
}
