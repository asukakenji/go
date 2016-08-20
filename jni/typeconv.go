package jni

/*
// free()
#include <stdlib.h>

#include <jni.h>
*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

/* Support Functions - Primitive Conversion (High Level) */

func JavaBoolean(z bool) C.jboolean {
	if z {
		return C.JNI_TRUE
	}
	return C.JNI_FALSE
}

func JavaByte(b int8) C.jbyte {
	return C.jbyte(b)
}

func JavaChar(c uint16) C.jchar {
	return C.jchar(c)
}

func JavaShort(s int16) C.jshort {
	return C.jshort(s)
}

var JavaInt = func() func(int) C.jint {
	// From src/strconv/itoa.go (strconv.formatBits()):
	//     if ^uintptr(0)>>32 == 0 {
	//         ...
	//     }
	if ^uint(0)>>32 == 0 {
		return JavaIntFor32Bit
	}
	return JavaIntFor64Bit
}()

// From https://golang.org/ref/spec#Numeric_types:
//     uint       either 32 or 64 bits
//     int        same size as uint
//     uintptr    an unsigned integer large enough to store the
//                uninterpreted bits of a pointer value
func JavaIntFor64Bit(i int) C.jint {
	// From src/math/const.go:
	//     const (
	//         MaxInt32 = 1<<31 - 1
	//         MinInt32 = -1 << 31
	//     )
	if i < (-1<<31) || i > (1<<31-1) {
		panic(fmt.Errorf("JavaInt(): input is not a valid signed 32-bit integer: %d", i))
	}
	return C.jint(int32(i))
}

func JavaIntFor32Bit(i int) C.jint {
	return C.jint(int32(i))
}

func JavaIntRaw(i int32) C.jint {
	return C.jint(i)
}

func JavaLong(j int64) C.jlong {
	return C.jlong(j)
}

func JavaFloat(f float32) C.jfloat {
	return C.jfloat(f)
}

func JavaDouble(d float64) C.jdouble {
	return C.jdouble(d)
}

func JavaSize(i int) C.jsize {
	return C.jsize(JavaInt(i))
}

func JavaSizeRaw(i int32) C.jsize {
	return C.jsize(JavaIntRaw(i))
}

func GoBool(z C.jboolean) bool {
	return z != C.JNI_FALSE
}

func GoInt8(b C.jbyte) int8 {
	return int8(b)
}

func GoUint16(c C.jchar) uint16 {
	return uint16(c)
}

func GoInt16(s C.jshort) int16 {
	return int16(s)
}

// From https://golang.org/ref/spec#Numeric_types:
//     uint       either 32 or 64 bits
//     int        same size as uint
//     uintptr    an unsigned integer large enough to store the
//                uninterpreted bits of a pointer value
func GoInt(i C.jint) int {
	return int(int32(i))
}

func GoIntFromSize(i C.jsize) int {
	return GoInt(C.jint(i))
}

func GoInt32(i C.jint) int32 {
	return int32(i)
}

func GoInt32FromSize(i C.jsize) int32 {
	return GoInt32(C.jint(i))
}

func GoInt64(j C.jlong) int64 {
	return int64(j)
}

func GoFloat32(f C.jfloat) float32 {
	return float32(f)
}

func GoFloat64(d C.jdouble) float64 {
	return float64(d)
}

/* Support Functions - String Conversion (High Level) */

// For example, see:
//     src/net/cgo_unix.go
//     src/os/user/lookup_unix.go
func WithCString(s string, f func(*C.char)) {
	pc := C.CString(s)
	defer C.free(unsafe.Pointer(pc))
	f(pc)
}

/* Support Functions - jvalue Conversion (Low Level) */

func JValueArray(args ...interface{}) []C.jvalue {
	result := make([]C.jvalue, len(args))
	for i, v := range args {
		switch v2 := v.(type) {
		case C.jboolean:
			result[i] = JValueFromJBoolean(v2)
		case C.jbyte:
			result[i] = JValueFromJByte(v2)
		case C.jchar:
			result[i] = JValueFromJChar(v2)
		case C.jshort:
			result[i] = JValueFromJShort(v2)
		case C.jint:
			result[i] = JValueFromJInt(v2)
		case C.jlong:
			result[i] = JValueFromJLong(v2)
		case C.jfloat:
			result[i] = JValueFromJFloat(v2)
		case C.jdouble:
			result[i] = JValueFromJDouble(v2)
		case C.jsize:
			result[i] = JValueFromJInt(C.jint(v2))
		case C.jobject:
			result[i] = JValueFromJObject(v2)
		case C.jclass:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jthrowable:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jstring:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jarray:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jbooleanArray:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jbyteArray:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jcharArray:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jshortArray:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jintArray:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jlongArray:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jfloatArray:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jdoubleArray:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jobjectArray:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jweak:
			result[i] = JValueFromJObject(C.jobject(v2))
		case C.jvalue:
			result[i] = v2
		default:
			panic("Unknown type")
		}
	}
	return result
}

func WithJValueArray(va []C.jvalue, f func(*C.jvalue)) {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&va))
	c_va := (*C.jvalue)(unsafe.Pointer(header.Data))
	f(c_va)
}
