package main

// #include "Main.h"
//
// // TODO: Should get rid of this after the 2 demo functions are removed.
// #include "jnienv.h"
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
	WithCString("test2", func(c_name *C.char) {
		WithCString("()V", func(c_sig *C.char) {
			if mid := C._GoJniGetStaticMethodID(env, clazz, c_name, c_sig); mid != nil {
				C._GoJniCallStaticVoidMethodA(env, clazz, mid, nil)
			}
		})
	})
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
type JObject struct {
	peer C.jobject
}

func NewJObject(peer C.jobject) JObject {
	return JObject{peer}
}

func (obj JObject) Peer() C.jobject {
	return obj.peer
}

// jni.h:
//     typedef jobject jclass;
type JClass JObject

func NewJClass(peer C.jclass) JClass {
	return JClass{C.jobject(peer)}
}

func (clazz JClass) Peer() C.jclass {
	return C.jclass(clazz.peer)
}

// jni.h:
//     typedef jobject jthrowable;
type JThrowable JObject

func NewJThrowable(peer C.jthrowable) JThrowable {
	return JThrowable{C.jobject(peer)}
}

func (obj JThrowable) Peer() C.jthrowable {
	return C.jthrowable(obj.peer)
}

// jni.h:
//     typedef jobject jstring;
// type JString JObject

// jni.h:
//     typedef jobject jarray;
// type JArray JObject

// jni.h:
//     typedef jarray jbooleanArray;
// type JBooleanArray JArray

// jni.h:
//     typedef jarray jbyteArray;
// type JByteArray JArray

// jni.h:
//     typedef jarray jcharArray;
// type JCharArray JArray

// jni.h:
//     typedef jarray jshortArray;
// type JShortArray JArray

// jni.h:
//     typedef jarray jintArray;
// type JIntArray JArray

// jni.h:
//     typedef jarray jlongArray;
// type JLongArray JArray

// jni.h:
//     typedef jarray jfloatArray;
// type JFloatArray JArray

// jni.h:
//     typedef jarray jdoubleArray;
// type JDoubleArray JArray

// jni.h:
//     typedef jarray jobjectArray;
// type JObjectArray JArray

// jni.h:
//     typedef jobject jweak;
type JWeak JObject

func NewJWeak(peer C.jweak) JWeak {
	return JWeak{C.jobject(peer)}
}

func (ref JWeak) Peer() C.jweak {
	return C.jweak(ref.peer)
}

// jni.h:
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
type JValue struct {
	peer C.jvalue
}

func NewJValue(peer C.jvalue) JValue {
	return JValue{peer}
}

func (value JValue) Peer() C.jvalue {
	return value.peer
}

/* See "jvalue.go" for related methods */

// jni.h:
//     struct _jfieldID;
//     typedef struct _jfieldID *jfieldID;
type JFieldID struct {
	peer C.jfieldID
}

func NewJFieldID(peer C.jfieldID) JFieldID {
	return JFieldID{peer}
}

func (fieldID JFieldID) Peer() C.jfieldID {
	return fieldID.peer
}

// jni.h:
//     struct _jmethodID;
//     typedef struct _jmethodID *jmethodID;
type JMethodID struct {
	peer C.jmethodID
}

func NewJMethodID(peer C.jmethodID) JMethodID {
	return JMethodID{peer}
}

func (methodID JMethodID) Peer() C.jmethodID {
	return methodID.peer
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
	JNIInvalidRefType    = C.JNIInvalidRefType
	JNILocalRefType      = C.JNILocalRefType
	JNIGlobalRefType     = C.JNIGlobalRefType
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
	JNI_TRUE  = 1
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
	JNI_OK        = 0
	JNI_ERR       = -1
	JNI_EDETACHED = -2
	JNI_EVERSION  = -3
	JNI_ENOMEM    = -4
	JNI_EEXIST    = -5
	JNI_EINVAL    = -6
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
	JNI_ABORT  = 2
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
//     typedef const struct JNINativeInterface_ *JNIEnv;
type JNIEnv struct {
	peer *C.JNIEnv
}

func NewJNIEnv(peer *C.JNIEnv) JNIEnv {
	return JNIEnv{peer}
}

func (env JNIEnv) Peer() *C.JNIEnv {
	return env.peer
}

// jni.h:
//     /*
//      * JNI Invocation Interface.
//      */
//
//     struct JNIInvokeInterface_;
//     typedef const struct JNIInvokeInterface_ *JavaVM;
type JavaVM struct {
	peer *C.JavaVM
}

func NewJavaVM(peer *C.JavaVM) JavaVM {
	return JavaVM{peer}
}

func (vm JavaVM) Peer() *C.JavaVM {
	return vm.peer
}

////////////////////
// JNIEnv Methods //
////////////////////

func (env JNIEnv) GetVersion() int32 {
	return GoInt32(GetVersion(env.Peer()))
}

func (env JNIEnv) DefineClass(name string, loader JObject, buf []byte, _len int32) JClass {
	// TODO: Implement this!
	panic("Not Yet Implemented")
}

func (env JNIEnv) FindClass(name string) JClass {
	var result C.jclass
	WithCString(name, func(c_name *C.char) {
		result = FindClass(env.Peer(), c_name)
	})
	return NewJClass(result)
}

func (env JNIEnv) FromReflectedMethod(method JObject) JMethodID {
	return NewJMethodID(FromReflectedMethod(env.Peer(), method.Peer()))
}

func (env JNIEnv) FromReflectedField(field JObject) JFieldID {
	return NewJFieldID(FromReflectedField(env.Peer(), field.Peer()))
}

func (env JNIEnv) ToReflectedMethod(cls JClass, methodID JMethodID, isStatic bool) JObject {
	return NewJObject(ToReflectedMethod(env.Peer(), cls.Peer(), methodID.Peer(), JavaBoolean(isStatic)))
}

func (env JNIEnv) GetSuperclass(sub JClass) JClass {
	return NewJClass(GetSuperclass(env.Peer(), sub.Peer()))
}

func (env JNIEnv) IsAssignableFrom(sub, sup JClass) bool {
	return GoBool(IsAssignableFrom(env.Peer(), sub.Peer(), sup.Peer()))
}

func (env JNIEnv) ToReflectedField(cls JClass, fieldID JFieldID, isStatic bool) JObject {
	return NewJObject(ToReflectedField(env.Peer(), cls.Peer(), fieldID.Peer(), JavaBoolean(isStatic)))
}

func (env JNIEnv) Throw(obj JThrowable) int32 {
	return GoInt32(Throw(env.Peer(), obj.Peer()))
}

func (env JNIEnv) ThrowNew(clazz JClass, msg string) int32 {
	var result C.jint
	WithCString(msg, func(c_msg *C.char) {
		result = ThrowNew(env.Peer(), clazz.Peer(), c_msg)
	})
	return GoInt32(result)
}

func (env JNIEnv) ExceptionOccurred() JThrowable {
	return NewJThrowable(ExceptionOccurred(env.Peer()))
}

func (env JNIEnv) ExceptionDescribe() {
	ExceptionDescribe(env.Peer())
}

func (env JNIEnv) ExceptionClear() {
	ExceptionClear(env.Peer())
}

func (env JNIEnv) FatalError(msg string) {
	WithCString(msg, func(c_msg *C.char) {
		FatalError(env.Peer(), c_msg)
	})
}

func (env JNIEnv) PushLocalFrame(capacity int32) int32 {
	return GoInt32(PushLocalFrame(env.Peer(), C.jint(capacity)))
}

func (env JNIEnv) PopLocalFrame(result JObject) JObject {
	return NewJObject(PopLocalFrame(env.Peer(), result.Peer()))
}

func (env JNIEnv) NewGlobalRef(lobj JObject) JObject {
	return NewJObject(NewGlobalRef(env.Peer(), lobj.Peer()))
}

func (env JNIEnv) DeleteGlobalRef(gref JObject) {
	DeleteGlobalRef(env.Peer(), gref.Peer())
}

func (env JNIEnv) DeleteLocalRef(obj JObject) {
	DeleteLocalRef(env.Peer(), obj.Peer())
}

func (env JNIEnv) IsSameObject(obj1, obj2 JObject) bool {
	return GoBool(IsSameObject(env.Peer(), obj1.Peer(), obj2.Peer()))
}

func (env JNIEnv) NewLocalRef(ref JObject) JObject {
	return NewJObject(NewLocalRef(env.Peer(), ref.Peer()))
}

func (env JNIEnv) EnsureLocalCapacity(capacity int32) int32 {
	return GoInt32(EnsureLocalCapacity(env.Peer(), C.jint(capacity)))
}

func (env JNIEnv) AllocObject(clazz JClass) JObject {
	return NewJObject(AllocObject(env.Peer(), clazz.Peer()))
}

func (env JNIEnv) NewObject(clazz JClass, methodID JMethodID, args []interface{}) JObject {
	// TODO: Implement this!
	panic("Not Yet Implemented")
}

func (env JNIEnv) GetObjectClass(obj JObject) JClass {
	return NewJClass(GetObjectClass(env.Peer(), obj.Peer()))
}

func (env JNIEnv) IsInstanceOf(obj JObject, clazz JClass) bool {
	return GoBool(IsInstanceOf(env.Peer(), obj.Peer(), clazz.Peer()))
}

// jni.h:
//     typedef struct JavaVMOption {
//         char *optionString;
//         void *extraInfo;
//     } JavaVMOption;

// jni.h:
//     typedef struct JavaVMInitArgs {
//         jint version;
//
//         jint nOptions;
//         JavaVMOption *options;
//         jboolean ignoreUnrecognized;
//     } JavaVMInitArgs;

// jni.h:
//     typedef struct JavaVMAttachArgs {
//         jint version;
//
//         char *name;
//         jobject group;
//     } JavaVMAttachArgs;

// jni.h:
//     jint (JNICALL *DestroyJavaVM)(JavaVM *vm);

// jni.h:
//     jint (JNICALL *AttachCurrentThread)(JavaVM *vm, void **penv, void *args);

// jni.h:
//     jint (JNICALL *DetachCurrentThread)(JavaVM *vm);

// jni.h:
//     jint (JNICALL *GetEnv)(JavaVM *vm, void **penv, jint version);

// jni.h:
//     jint (JNICALL *AttachCurrentThreadAsDaemon)(JavaVM *vm, void **penv, void *args);

// jni.h:
//     _JNI_IMPORT_OR_EXPORT_ jint JNICALL JNI_GetDefaultJavaVMInitArgs(void *args);
// jni.h:
//     _JNI_IMPORT_OR_EXPORT_ jint JNICALL JNI_CreateJavaVM(JavaVM **pvm, void **penv, void *args);
// jni.h:
//     _JNI_IMPORT_OR_EXPORT_ jint JNICALL JNI_GetCreatedJavaVMs(JavaVM **, jsize, jsize *);

/* Defined by native libraries. */
// jni.h:
//     JNIEXPORT jint JNICALL JNI_OnLoad(JavaVM *vm, void *reserved);
// jni.h:
//     JNIEXPORT void JNICALL JNI_OnUnload(JavaVM *vm, void *reserved);

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

/* Support Functions - Primitive Conversion */

func JavaBoolean(z bool) C.jboolean {
	if z {
		return JNI_TRUE
	}
	return JNI_FALSE
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

func JavaInt(i int32) C.jint {
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

func GoBool(z C.jboolean) bool {
	return z != JNI_FALSE
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

func GoInt32(i C.jint) int32 {
	return int32(i)
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

/* Support Functions - String Conversion */

// For example, see:
//     src/net/cgo_unix.go
//     src/os/user/lookup_unix.go
func WithCString(s string, f func(*C.char)) {
	pc := C.CString(s)
	defer C.free(unsafe.Pointer(pc))
	f(pc)
}
