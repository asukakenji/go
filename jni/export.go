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
	fmt.Println(GetVersion(env))
	fmt.Println("Hello, world from a Go function!")
	test2(env, clazz)
}

func test2(env *C.JNIEnv, clazz C.jclass) {
	if name := C.CString("test2"); name != nil {
		defer C.free(unsafe.Pointer(name))
		if sig := C.CString("()V"); sig != nil {
			defer C.free(unsafe.Pointer(sig))
			if mid := C._GoJniGetStaticMethodID(env, clazz, name, sig); mid != nil {
				C._GoJniCallStaticVoidMethod(env, clazz, mid)
			}
		}
	}
}

func main() {
}



// jni.h:
//     /*
//      * JNI Types
//      */

// jni.h:
//     typedef jint jsize;
/* No need to do anything */

// jni.h:
//     struct _jobject;
//     typedef struct _jobject *jobject;
type Object struct {
	peer C.jobject
}

// jni.h:
//     typedef jobject jclass;
type Class Object

// jni.h:
//     typedef jobject jthrowable;
type Throwable Object

// jni.h:
//     typedef jobject jstring;
// type String Object

// jni.h:
//     typedef jobject jarray;
// type Array Object

// jni.h:
//     typedef jarray jbooleanArray;
// type BooleanArray Array

// jni.h:
//     typedef jarray jbyteArray;
// type ByteArray Array

// jni.h:
//     typedef jarray jcharArray;
// type CharArray Array

// jni.h:
//     typedef jarray jshortArray;
// type ShortArray Array

// jni.h:
//     typedef jarray jintArray;
// type IntArray Array

// jni.h:
//     typedef jarray jlongArray;
// type LongArray Array

// jni.h:
//     typedef jarray jfloatArray;
// type FloatArray Array

// jni.h:
//     typedef jarray jdoubleArray;
// type DoubleArray Array

// jni.h:
//     typedef jarray jobjectArray;
// type ObjectArray Array

// jni.h
//     typedef jobject jweak;
type Weak Object

// jni.h
//     typedef union jvalue {
//         jboolean z;
//         jbyte    b;
//         jchar    c;
//         jshort   s;
//         jint     i;
//         jlong    j;
//         jfloat   f;
//         jdouble  d;
//         jobject  l;
//     } jvalue;
/* Not Implemented */

// jni.h:
//     struct _jfieldID;
//     typedef struct _jfieldID *jfieldID;
type FieldID struct {
	peer C.jfieldID
}

// jni.h:
//     struct _jmethodID;
//     typedef struct _jmethodID *jmethodID;
type MethodID struct {
	peer C.jmethodID
}

// jni.h:
//     /* Return values from jobjectRefType */
//     typedef enum _jobjectType {
//         JNIInvalidRefType    = 0,
//         JNILocalRefType      = 1,
//         JNIGlobalRefType     = 2,
//         JNIWeakGlobalRefType = 3
//     } jobjectRefType;
const (
	JNIInvalidRefType = C.JNIInvalidRefType
	JNILocalRefType = C.JNILocalRefType
	JNIGlobalRefType = C.JNIGlobalRefType
	JNIWeakGlobalRefType = C.JNIWeakGlobalRefType
)

// jni.h:
//     /*
//      * jboolean constants
//      */
//     
//     #define JNI_FALSE 0
//     #define JNI_TRUE 1
const (
	JNI_FALSE = 0
	JNI_TRUE = 1
)

// jni.h:
//     /*
//      * possible return values for JNI functions.
//      */
//     
//     #define JNI_OK           0                 /* success */
//     #define JNI_ERR          (-1)              /* unknown error */
//     #define JNI_EDETACHED    (-2)              /* thread detached from the VM */
//     #define JNI_EVERSION     (-3)              /* JNI version error */
//     #define JNI_ENOMEM       (-4)              /* not enough memory */
//     #define JNI_EEXIST       (-5)              /* VM already created */
//     #define JNI_EINVAL       (-6)              /* invalid arguments */
const (
	JNI_OK = 0
	JNI_ERR = -1
	JNI_EDETACHED = -2
	JNI_EVERSION = -3
	JNI_ENOMEM = -4
	JNI_EEXIST = -5
	JNI_EINVAL = -6
)

// jni.h:
//     /*
//      * used in ReleaseScalarArrayElements
//      */
//     
//     #define JNI_COMMIT 1
//     #define JNI_ABORT 2
const (
	JNI_COMMIT = 1
	JNI_ABORT = 2
)

// jni.h:
//     /*
//      * used in RegisterNatives to describe native method name, signature,
//      * and function pointer.
//      */
//     
//     typedef struct {
//         char *name;
//         char *signature;
//         void *fnPtr;
//     } JNINativeMethod;
/* Not Implemented */

// jni.h:
//     /*
//      * JNI Native Method Interface.
//      */
//     
//     struct JNINativeInterface_;
//     
//     struct JNIEnv_;
//
//     typedef const struct JNINativeInterface_ *JNIEnv;

// jni.h:
//     /*
//      * JNI Invocation Interface.
//      */
//     
//     struct JNIInvokeInterface_;
//     
//     struct JavaVM_;
//     
//     typedef const struct JNIInvokeInterface_ *JavaVM;

// jni.h:
//     jint (JNICALL *GetVersion)(JNIEnv *env);
func GetVersion(env *C.JNIEnv) C.jint {
	return C._GoJniGetVersion(env)
}

// jni.h:
//     jclass (JNICALL *DefineClass)(JNIEnv *env, const char *name, jobject loader, const jbyte *buf, jsize len);
func DefineClass(env *C.JNIEnv, name *C.char, loader C.jobject, buf *C.jbyte, _len C.jsize) C.jclass {
	return C._GoJniDefineClass(env, name, loader, buf, _len)
}

// jni.h:
//     jclass (JNICALL *FindClass)(JNIEnv *env, const char *name);
func FindClass(env *C.JNIEnv, name *C.char) C.jclass {
	return C._GoJniFindClass(env, name)
}

// jni.h:
//     jmethodID (JNICALL *FromReflectedMethod)(JNIEnv *env, jobject method);
func FromReflectedMethod(env *C.JNIEnv, method C.jobject) C.jmethodID {
	return C._GoJniFromReflectedMethod(env, method)
}

// jni.h:
//     jfieldID (JNICALL *FromReflectedField)(JNIEnv *env, jobject field);
func FromReflectedField(env *C.JNIEnv, field C.jobject) C.jfieldID {
	return C._GoJniFromReflectedField(env, field)
}

// jni.h:
//     jobject (JNICALL *ToReflectedMethod)(JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic);
func ToReflectedMethod(env *C.JNIEnv, cls C.jclass, methodID C.jmethodID, isStatic C.jboolean) C.jobject {
	return C._GoJniToReflectedMethod(env, cls, methodID, isStatic)
}

// jni.h:
//     jclass (JNICALL *GetSuperclass)(JNIEnv *env, jclass sub);
func GetSuperclass(env *C.JNIEnv, sub C.jclass) C.jclass {
	return C._GoJniGetSuperclass(env, sub)
}

// jni.h:
//     jboolean (JNICALL *IsAssignableFrom)(JNIEnv *env, jclass sub, jclass sup);
func IsAssignableFrom(env *C.JNIEnv, sub, sup C.jclass) C.jboolean {
	return C._GoJniIsAssignableFrom(env, sub, sup)
}

// jni.h:
//     jobject (JNICALL *ToReflectedField)(JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic);
func ToReflectedField(env *C.JNIEnv, cls C.jclass, fieldID C.jfieldID, isStatic C.jboolean) C.jobject {
	return C._GoJniToReflectedField(env, cls, fieldID, isStatic)
}



type Env struct {
	peer *C.JNIEnv
}

// jni.h:
//     #define JNI_VERSION_1_1 0x00010001
//     #define JNI_VERSION_1_2 0x00010002
//     #define JNI_VERSION_1_4 0x00010004
//     #define JNI_VERSION_1_6 0x00010006
//     #define JNI_VERSION_1_8 0x00010008
const (
	JNI_VERSION_1_1 = C.JNI_VERSION_1_1
	JNI_VERSION_1_2 = C.JNI_VERSION_1_2
	JNI_VERSION_1_4 = C.JNI_VERSION_1_4
	JNI_VERSION_1_6 = C.JNI_VERSION_1_6
	JNI_VERSION_1_8 = C.JNI_VERSION_1_8
)
