package jni

// #include <jni.h>
import "C"

import (
	"unsafe"
)

// Note:
// If there are `func init()`s, they will be called before `JNI_Onload()`.

var defaultVM *C.JavaVM

//export JNI_OnLoad
func JNI_OnLoad(vm *C.JavaVM, reserved unsafe.Pointer) C.jint {
	defaultVM = vm
	return C.JNI_VERSION_1_6
}

/*
//export JNI_OnLoad_gojni
func JNI_OnLoad_gojni(vm *C.JavaVM, reserved unsafe.Pointer) C.jint {
	fmt.Println("JNI_OnLoad_gojni called")
	return C.JNI_VERSION_1_8
}
*/

func GetDefaultJavaVM() *C.JavaVM {
	return defaultVM
}

// TODO: Implement this!
func GetDefaultJNIEnv() *C.JNIEnv {
	//return GetEnv(defaultVM, nil, C.JNI_VERSION_1_6)
	return nil
}
