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

type CaptureDevice uintptr

func OpenCaptureDevice(deviceName string, frequency Int, format Enum, bufferSize Sizei) CaptureDevice {
	cDeviceName := C.CString(deviceName)
	defer C.free(unsafe.Pointer(cDeviceName))
	return GoCaptureDevice(C.alcCaptureOpenDevice(cDeviceName, C.ALCuint(frequency), C.ALCenum(format), C.ALCsizei(bufferSize)))
}

func (d CaptureDevice) Close() bool {
	return GoBool(C.alcCaptureCloseDevice(d.CCaptureDevice()))
}

func (d CaptureDevice) Start() {
	C.alcCaptureStart(d.CCaptureDevice())
}

func (d CaptureDevice) Stop() {
	C.alcCaptureStop(d.CCaptureDevice())
}

func (d CaptureDevice) Samples(buffer []byte, samples Sizei) {
	cBuffer := unsafe.Pointer(&buffer[0])
	C.alcCaptureSamples(d.CCaptureDevice(), cBuffer, C.ALCsizei(samples))
}

// GoCaptureDevice : C capture device to Go capture device.
func GoCaptureDevice(cCaptureDevice *C.ALCdevice) CaptureDevice {
	return CaptureDevice(unsafe.Pointer(cCaptureDevice))
}

// CCaptureDevice : Go capture device to C capture device.
func (d CaptureDevice) CCaptureDevice() *C.ALCdevice {
	return (*C.ALCdevice)(unsafe.Pointer(d))
}
