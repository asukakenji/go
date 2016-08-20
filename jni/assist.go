package jni

/*
#include <jni.h>
*/
import "C"

import (
	"unsafe"
)

// These functions exist to help exporting functions in the low-level API.
//
// Functions with C type as a parameter are not callable from another package.
// For instance, the compiler treats C.jclass in this package and C.jclass in
// the client's package as different types. Therefore they are incompatiable,
// and cannot be casted directly between each other.
//
// See http://stackoverflow.com/q/39013201/142239 for the discussion.
//
// Therefore, the last resort is to cast to unsafe.Pointer and cast back to the
// actual type. However, since C types in this package (the actual types) are
// not accessible from any other packages (for instance, the compiler treats
// C.jclass in this package as jni.C.jclass), there must be functions in this
// package to cast back the unsafe.Pointer to the actual type.

func GetCJBoolean(z unsafe.Pointer) *C.jboolean {
	return (*C.jboolean)(z)
}

func GetCJByte(b unsafe.Pointer) *C.jbyte {
	return (*C.jbyte)(b)
}

func GetCJChar(c unsafe.Pointer) *C.jchar {
	return (*C.jchar)(c)
}

func GetCJShort(s unsafe.Pointer) *C.jshort {
	return (*C.jshort)(s)
}

func GetCJInt(i unsafe.Pointer) *C.jint {
	return (*C.jint)(i)
}

func GetCJLong(j unsafe.Pointer) *C.jlong {
	return (*C.jlong)(j)
}

func GetCJFloat(f unsafe.Pointer) *C.jfloat {
	return (*C.jfloat)(f)
}

func GetCJDouble(d unsafe.Pointer) *C.jdouble {
	return (*C.jdouble)(d)
}

func GetCJSize(i unsafe.Pointer) *C.jsize {
	return (*C.jsize)(i)
}

func GetCJObject(obj unsafe.Pointer) C.jobject {
	return C.jobject(obj)
}

func GetCJClass(clazz unsafe.Pointer) C.jclass {
	return C.jclass(clazz)
}

func GetCJThrowable(throw unsafe.Pointer) C.jthrowable {
	return C.jthrowable(throw)
}

func GetCJString(str unsafe.Pointer) C.jstring {
	return C.jstring(str)
}

func GetCJArray(arr unsafe.Pointer) C.jarray {
	return C.jarray(arr)
}

func GetCJBooleanArray(arr unsafe.Pointer) C.jbooleanArray {
	return C.jbooleanArray(arr)
}

func GetCJByteArray(arr unsafe.Pointer) C.jbyteArray {
	return C.jbyteArray(arr)
}

func GetCJCharArray(arr unsafe.Pointer) C.jcharArray {
	return C.jcharArray(arr)
}

func GetCJShortArray(arr unsafe.Pointer) C.jshortArray {
	return C.jshortArray(arr)
}

func GetCJIntArray(arr unsafe.Pointer) C.jintArray {
	return C.jintArray(arr)
}

func GetCJLongArray(arr unsafe.Pointer) C.jlongArray {
	return C.jlongArray(arr)
}

func GetCJFloatArray(arr unsafe.Pointer) C.jfloatArray {
	return C.jfloatArray(arr)
}

func GetCJDoubleArray(arr unsafe.Pointer) C.jdoubleArray {
	return C.jdoubleArray(arr)
}

func GetCJObjectArray(arr unsafe.Pointer) C.jobjectArray {
	return C.jobjectArray(arr)
}

func GetCJWeak(weak unsafe.Pointer) C.jweak {
	return C.jweak(weak)
}

func GetCJValue(v unsafe.Pointer) *C.jvalue {
	return (*C.jvalue)(v)
}

func GetCJFieldID(fieldID unsafe.Pointer) C.jfieldID {
	return C.jfieldID(fieldID)
}

func GetCJMethodID(methodID unsafe.Pointer) C.jmethodID {
	return C.jmethodID(methodID)
}

func GetCJNIEnv(env unsafe.Pointer) *C.JNIEnv {
	return (*C.JNIEnv)(env)
}

func GetCJavaVM(vm unsafe.Pointer) *C.JavaVM {
	return (*C.JavaVM)(vm)
}

func GetCJNINativeMethod(method unsafe.Pointer) *C.JNINativeMethod {
	return (*C.JNINativeMethod)(method)
}

func GetPtrCChar(pc unsafe.Pointer) *C.char {
	return (*C.char)(pc)
}

func GetPtrCJValue(pv unsafe.Pointer) *C.jvalue {
	return (*C.jvalue)(pv)
}

func UnwrapPtrCChar(f func(unsafe.Pointer)) func(*C.char) {
	return func(pc *C.char) {
		f(unsafe.Pointer(pc))
	}
}

func UnwrapPtrCJValue(f func(unsafe.Pointer)) func(*C.jvalue) {
	return func(pv *C.jvalue) {
		f(unsafe.Pointer(pv))
	}
}
