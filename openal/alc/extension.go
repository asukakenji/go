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

func (d Device) IsExtensionPresent(extName string) bool {
	cExtName := C.CString(extName)
	defer C.free(unsafe.Pointer(cExtName))
	return GoBool(C.alcIsExtensionPresent(d.CDevice(), cExtName))
}

func (d Device) GetProcAddress(funcName string) unsafe.Pointer {
	cFuncName := C.CString(funcName)
	defer C.free(unsafe.Pointer(cFuncName))
	return C.alcGetProcAddress(d.CDevice(), cFuncName)
}

func (d Device) GetEnumValue(device Device, enumName string) Enum {
	cEnumName := C.CString(enumName)
	defer C.free(unsafe.Pointer(cEnumName))
	return Enum(C.alcGetEnumValue(d.CDevice(), cEnumName))
}
