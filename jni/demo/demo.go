package main

/*
#include "Main.h"

#include <jni.h>
*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/asukakenji/go/jni"
)

//export Java_Main_test
func Java_Main_test(env *C.JNIEnv, clazz C.jclass) {
	fmt.Println(jni.GetVersion(jni.GetCJNIEnv(unsafe.Pointer(env))))
	fmt.Println("Hello, world from a Go function!")
	test2(env, clazz)
}

func test2(env *C.JNIEnv, clazz C.jclass) {
	env2 := jni.GetCJNIEnv(unsafe.Pointer(env))
	clazz2 := jni.GetCJClass(unsafe.Pointer(clazz))
	jni.WithCString("test2", jni.UnwrapPtrCChar(WrapPtrCChar(func(c_name *C.char) {
		jni.WithCString("()V", jni.UnwrapPtrCChar(WrapPtrCChar(func(c_sig *C.char) {
			c_name2 := jni.GetPtrCChar(unsafe.Pointer(c_name))
			c_sig2 := jni.GetPtrCChar(unsafe.Pointer(c_sig))
			if mid := jni.GetStaticMethodID(env2, clazz2, c_name2, c_sig2); mid != nil {
				jni.CallStaticVoidMethodA(env2, clazz2, mid, nil)
			}
		})))
	})))
}

func WrapPtrCChar(f func(*C.char)) func(unsafe.Pointer) {
	return func(pc unsafe.Pointer) {
		f((*C.char)(pc))
	}
}

func main() {
}
