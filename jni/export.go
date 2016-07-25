package main

// #include "Main.h"
//
// #include "bridge.h"
//
// // JNIEnv, jclass
// #include <jni.h>
//
// // free()
// #include <stdlib.h>
import "C"

import (
	"fmt"
	"unsafe"
)

//export Java_Main_test
func Java_Main_test(env *C.JNIEnv, clazz C.jclass) {
	fmt.Println("Hello, world from a Go function!")
	test2(env, clazz)
}

func test2(env *C.JNIEnv, clazz C.jclass) {
	if name := C.CString("test2"); name != nil {
		defer C.free(unsafe.Pointer(name))
		if sig := C.CString("()V"); sig != nil {
			defer C.free(unsafe.Pointer(sig))
			if mid := C.GoGetStaticMethodID(env, clazz, name, sig); mid != nil {
				C.GoCallStaticVoidMethod(env, clazz, mid)
			}
		}
	}
}

func main() {
}
