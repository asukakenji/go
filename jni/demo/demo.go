package main

/*
#include "Main.h"

#include <jni.h>
*/
import "C"

import (
	"fmt"

	"github.com/asukakenji/go/jni"
)

//export Java_Main_test
func Java_Main_test(env *C.JNIEnv, clazz C.jclass) {
	fmt.Println(jni.GetVersion(env))
	fmt.Println("Hello, world from a Go function!")
	test2(env, clazz)
}

func test2(env *C.JNIEnv, clazz C.jclass) {
	jni.WithCString("test2", func(c_name *C.char) {
		jni.WithCString("()V", func(c_sig *C.char) {
			if mid := jni.GetStaticMethodID(env, clazz, c_name, c_sig); mid != nil {
				jni.CallStaticVoidMethodA(env, clazz, mid, nil)
			}
		})
	})
}

func main() {
}
