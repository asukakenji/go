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
import (
	"fmt"
	"unsafe"
)

// Context represents an OpenAL context.
type Context uintptr

// CreateContext creates a context using a specified device.
func CreateContext(device Device, options map[Int]interface{}) Context {
	attrList := make([]C.ALCint, len(options)*2+1)
	index := 0
	for key, value := range options {
		attrList[index] = C.ALCint(key)
		index++
		switch key {
		case Frequency, Refresh, MonoSources, StereoSources:
			attrList[index] = C.ALCint(value.(Int))
		case Sync:
			attrList[index] = C.ALCint(value.(Boolean))
		default:
			panic(fmt.Sprintf("Unknown key: %d", key))
		}
		index++
	}
	attrList[index] = C.ALCint(0)
	return GoContext(C.alcCreateContext(device.CDevice(), &attrList[0]))
}

// CreateDefaultContext creates a context using a specified device and default attributes.
func CreateDefaultContext(device Device) Context {
	return GoContext(C.alcCreateContext(device.CDevice(), nil))
}

// MakeContextCurrent makes a specified context the current context.
func MakeContextCurrent(c Context) bool {
	return GoBool(C.alcMakeContextCurrent(c.CContext()))
}

// GetCurrentContext retrieves the current context.
func GetCurrentContext() Context {
	return GoContext(C.alcGetCurrentContext())
}

// Process tells context c to begin processing.
func (c Context) Process() {
	C.alcProcessContext(c.CContext())
}

// Suspend suspends processing on context c.
func (c Context) Suspend() {
	C.alcSuspendContext(c.CContext())
}

// Destroy destroys context c.
func (c Context) Destroy() {
	C.alcDestroyContext(c.CContext())
}

// Device retrieves context c's device.
func (c Context) Device() Device {
	return GoDevice(C.alcGetContextsDevice(c.CContext()))
}

// GoContext : C context to Go context.
func GoContext(cContext *C.ALCcontext) Context {
	return Context(unsafe.Pointer(cContext))
}

// CContext : Go context to C context.
func (c Context) CContext() *C.ALCcontext {
	return (*C.ALCcontext)(unsafe.Pointer(c))
}
