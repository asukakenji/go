package main

/*
#include <jni.h>
*/
import "C"

// This must be implemented with low level API,
// since JNIEnv.CallObjectMethod() calls this method
func (t JThrowable) GetMessage() string {
	// Preparation
	obj := t.JObjectPeer()

	// Get Method ID
	// TODO: Write this!
	//env := GetEnvFromJavaVM()
	env := (*C.JNIEnv)(nil)
	clazz := GetObjectClass(env, obj)
	name := "getMessage"
	sig := "()Ljava/lang/String;"
	var methodID C.jmethodID
	WithCString(name, func(c_name *C.char) {
		WithCString(sig, func(c_sig *C.char) {
			methodID = GetMethodID(env, clazz, c_name, c_sig)
		})
	})
	// TODO: Exception handling for GetMethodID()

	// Calling CallObjectMethod
	//jmessage := C.jstring(CallObjectMethod(env, obj, methodID, nil))
	//jmessage = jmessage
	// Options:
	//     GetStringLength / GetStringChars / ReleaseStringChars (const jchar*)
	//     GetStringUTFLength / GetStringUTFChars / ReleaseStringUTFChars (const char*)
	//     GetStringRegion (jchar*)
	//     GetStringUTFRegion (char*)
	//     GetStringCritical / ReleaseStringCritical (const jchar*)
	message := ""
	// TODO: Exception handling for CallObjectMethod()

	return message
}

type Throwable struct {
	message string
}

func ToThrowable(t JThrowable) Throwable {
	return Throwable{t.GetMessage()}
}

// Implements the built-in "error" interface
// See: https://golang.org/ref/spec#Errors
func (t Throwable) Error() string {
	return t.message
}

type ClassFormatError Throwable

type ClassCircularityError Throwable

type OutOfMemoryError Throwable

type SecurityException Throwable
