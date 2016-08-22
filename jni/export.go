package jni

/*
// free()
#include <stdlib.h>

#include <jni.h>
*/
import "C"

import (
	"errors"
	"fmt"
	"reflect"
	"unicode/utf16"
	"unicode/utf8"
	"unsafe"
)

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
type JObject uintptr

func GoJObject(peer C.jobject) JObject {
	return JObject(uintptr(unsafe.Pointer(peer)))
}

func (obj JObject) JObjectPeer() C.jobject {
	return C.jobject(unsafe.Pointer(uintptr(obj)))
}

func (obj JObject) Peer() C.jobject {
	return obj.JObjectPeer()
}

// jni.h:
//     typedef jobject jclass;
type JClass JObject

func GoJClass(peer C.jclass) JClass {
	return JClass(GoJObject(C.jobject(peer)))
}

func (clazz JClass) JObjectPeer() C.jobject {
	return JObject(clazz).JObjectPeer()
}

func (clazz JClass) Peer() C.jclass {
	return C.jclass(clazz.JObjectPeer())
}

// jni.h:
//     typedef jobject jthrowable;
type JThrowable JObject

func GoJThrowable(peer C.jthrowable) JThrowable {
	return JThrowable(GoJObject(C.jobject(peer)))
}

func (throw JThrowable) JObjectPeer() C.jobject {
	return JObject(throw).JObjectPeer()
}

func (throw JThrowable) Peer() C.jthrowable {
	return C.jthrowable(throw.JObjectPeer())
}

/* See "jthrowable.go" for related methods */

// jni.h:
//     typedef jobject jstring;
type JString JObject

func GoJString(peer C.jstring) JString {
	return JString(GoJObject(C.jobject(peer)))
}

func (str JString) JObjectPeer() C.jobject {
	return JObject(str).JObjectPeer()
}

func (str JString) Peer() C.jstring {
	return C.jstring(str.JObjectPeer())
}

/* See "jstring.go" for related methods */

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

func GoJWeak(peer C.jweak) JWeak {
	return JWeak(GoJObject(C.jobject(peer)))
}

func (weak JWeak) JObjectPeer() C.jobject {
	return JObject(weak).JObjectPeer()
}

func (weak JWeak) Peer() C.jweak {
	return C.jweak(weak.JObjectPeer())
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

/* See "jvalue.go" for related methods */

// jni.h:
//     struct _jfieldID;
//     typedef struct _jfieldID *jfieldID;
type JFieldID uintptr

func GoJFieldID(peer C.jfieldID) JFieldID {
	return JFieldID(uintptr(unsafe.Pointer(peer)))
}

func (fieldID JFieldID) Peer() C.jfieldID {
	return C.jfieldID(unsafe.Pointer(uintptr(fieldID)))
}

// jni.h:
//     struct _jmethodID;
//     typedef struct _jmethodID *jmethodID;
type JMethodID uintptr

func GoJMethodID(peer C.jmethodID) JMethodID {
	return JMethodID(uintptr(unsafe.Pointer(peer)))
}

func (methodID JMethodID) Peer() C.jmethodID {
	return C.jmethodID(unsafe.Pointer(uintptr(methodID)))
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
	// To prevent the values from not being defined in C side,
	// actual values are used.
	JNIInvalidRefType    = 0
	JNILocalRefType      = 1
	JNIGlobalRefType     = 2
	JNIWeakGlobalRefType = 3
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

var Errors = []error{
	errors.New("JNI_OK: Success"), // This should never be used. Use nil as err for success.
	errors.New("JNI_ERR: Unknown error"),
	errors.New("JNI_EDETACHED: Thread detached from the VM"),
	errors.New("JNI_EVERSION: JNI version error"),
	errors.New("JNI_ENOMEM: Not enough memory"),
	errors.New("JNI_EEXIST: VM already created"),
	errors.New("JNI_EINVAL: Invalid arguments"),
}

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
type JNIEnv uintptr

func GoJNIEnv(peer *C.JNIEnv) JNIEnv {
	return JNIEnv(uintptr(unsafe.Pointer(peer)))
}

func (env JNIEnv) Peer() *C.JNIEnv {
	return (*C.JNIEnv)(unsafe.Pointer(uintptr(env)))
}

/* See "jnienv_low_level.go" for related methods */

// jni.h:
//     /*
//      * JNI Invocation Interface.
//      */
//
//     struct JNIInvokeInterface_;
//     typedef const struct JNIInvokeInterface_ *JavaVM;
type JavaVM uintptr

func GoJavaVM(peer *C.JavaVM) JavaVM {
	return JavaVM(uintptr(unsafe.Pointer(peer)))
}

func (vm JavaVM) Peer() *C.JavaVM {
	return (*C.JavaVM)(unsafe.Pointer(uintptr(vm)))
}

////////////////////
// JNIEnv Methods //
////////////////////

func (env JNIEnv) GetVersion() (Version, error) {
	result := GoInt32(GetVersion(env.Peer()))
	if result < 0 {
		return 0, Errors[-result]
	}
	return Version(result), nil
}

func (env JNIEnv) DefineClass(name string, loader JObject, buf []byte, _len int) (JClass, error) {
	var result C.jclass
	WithCString(name, func(c_name *C.char) {
		header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
		c_buf := (*C.jbyte)(unsafe.Pointer(header.Data))
		c_len := JavaSize(header.Len)
		result = DefineClass(env.Peer(), c_name, loader.Peer(), c_buf, c_len)
	})
	if env.ExceptionCheck() {
		// TODO: Delete local reference (throwable)
		throwable := env.ExceptionOccurred()
		throwable = throwable
		// TODO: Implement this!
		var err error
		return GoJClass(nil), err
	}
	// TODO: Handle exceptions
	return GoJClass(result), nil
}

func (env JNIEnv) FindClass(name string) JClass {
	var result C.jclass
	WithCString(name, func(c_name *C.char) {
		result = FindClass(env.Peer(), c_name)
	})
	// TODO: Handle exceptions
	return GoJClass(result)
}

func (env JNIEnv) FromReflectedMethod(method JObject) JMethodID {
	return GoJMethodID(FromReflectedMethod(env.Peer(), method.Peer()))
}

func (env JNIEnv) FromReflectedField(field JObject) JFieldID {
	return GoJFieldID(FromReflectedField(env.Peer(), field.Peer()))
}

func (env JNIEnv) ToReflectedMethod(cls JClass, methodID JMethodID, isStatic bool) JObject {
	return GoJObject(ToReflectedMethod(env.Peer(), cls.Peer(), methodID.Peer(), JavaBoolean(isStatic)))
}

func (env JNIEnv) GetSuperclass(sub JClass) JClass {
	return GoJClass(GetSuperclass(env.Peer(), sub.Peer()))
}

func (env JNIEnv) IsAssignableFrom(sub, sup JClass) bool {
	return GoBool(IsAssignableFrom(env.Peer(), sub.Peer(), sup.Peer()))
}

func (env JNIEnv) ToReflectedField(cls JClass, fieldID JFieldID, isStatic bool) JObject {
	return GoJObject(ToReflectedField(env.Peer(), cls.Peer(), fieldID.Peer(), JavaBoolean(isStatic)))
}

func (env JNIEnv) Throw(obj JThrowable) int {
	return GoInt(Throw(env.Peer(), obj.Peer()))
}

func (env JNIEnv) ThrowNew(clazz JClass, msg string) int {
	var result C.jint
	WithCString(msg, func(c_msg *C.char) {
		result = ThrowNew(env.Peer(), clazz.Peer(), c_msg)
	})
	return GoInt(result)
}

func (env JNIEnv) ExceptionOccurred() JThrowable {
	return GoJThrowable(ExceptionOccurred(env.Peer()))
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

func (env JNIEnv) PushLocalFrame(capacity int) int {
	return GoInt(PushLocalFrame(env.Peer(), C.jint(capacity)))
}

func (env JNIEnv) PopLocalFrame(result JObject) JObject {
	return GoJObject(PopLocalFrame(env.Peer(), result.Peer()))
}

func (env JNIEnv) NewGlobalRef(lobj JObject) JObject {
	return GoJObject(NewGlobalRef(env.Peer(), lobj.Peer()))
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
	return GoJObject(NewLocalRef(env.Peer(), ref.Peer()))
}

func (env JNIEnv) EnsureLocalCapacity(capacity int) int {
	return GoInt(EnsureLocalCapacity(env.Peer(), C.jint(capacity)))
}

func (env JNIEnv) AllocObject(clazz JClass) JObject {
	return GoJObject(AllocObject(env.Peer(), clazz.Peer()))
}

func (env JNIEnv) NewObject(clazz JClass, methodID JMethodID, args ...interface{}) JObject {
	// TODO: Implement this!
	panic("Not Yet Implemented")
}

func (env JNIEnv) GetObjectClass(obj JObject) JClass {
	return GoJClass(GetObjectClass(env.Peer(), obj.Peer()))
}

func (env JNIEnv) IsInstanceOf(obj JObject, clazz JClass) bool {
	return GoBool(IsInstanceOf(env.Peer(), obj.Peer(), clazz.Peer()))
}

func (env JNIEnv) GetMethodID(clazz JClass, name, sig string) JMethodID {
	var result C.jmethodID
	WithCString(name, func(c_name *C.char) {
		WithCString(sig, func(c_sig *C.char) {
			result = GetMethodID(env.Peer(), clazz.Peer(), c_name, c_sig)
		})
	})
	return GoJMethodID(result)
}

func (env JNIEnv) CallObjectMethod(obj JObject, methodID JMethodID, args ...interface{}) JObject {
	argsCopy := make([]C.jvalue, len(args))
	for i, v := range args {
		value := reflect.ValueOf(v)
		kind := value.Kind()
		switch kind {
		case reflect.Bool:
			argsCopy[i] = JValueFromJBoolean(JavaBoolean(value.Bool()))
		case reflect.Int:
			argsCopy[i] = JValueFromJInt(JavaInt(int(value.Int())))
		case reflect.Int8:
			argsCopy[i] = JValueFromJByte(JavaByte(int8(value.Int())))
		case reflect.Int16:
			argsCopy[i] = JValueFromJShort(JavaShort(int16(value.Int())))
		case reflect.Int32:
			argsCopy[i] = JValueFromJInt(JavaIntRaw(int32(value.Int())))
		case reflect.Int64:
			argsCopy[i] = JValueFromJLong(JavaLong(int64(value.Int())))
		case reflect.Uint16:
			argsCopy[i] = JValueFromJChar(JavaChar(uint16(value.Uint())))
		case reflect.Float32:
			argsCopy[i] = JValueFromJFloat(JavaFloat(float32(value.Float())))
		case reflect.Float64:
			argsCopy[i] = JValueFromJDouble(JavaDouble(float64(value.Float())))
		case reflect.Array:
			panic("Not Yet Implemented")
		case reflect.Interface:
			panic("Not Yet Implemented")
		case reflect.Map:
			panic("Not Yet Implemented")
		case reflect.Slice:
			panic("Not Yet Implemented")
		case reflect.String:
			panic("Not Yet Implemented")
		default:
			panic("Not Yet Implemented")
		}
	}
	header := (*reflect.SliceHeader)(unsafe.Pointer(&argsCopy))
	c_argsCopy := (*C.jvalue)(unsafe.Pointer(header.Data))
	result := CallObjectMethodA(env.Peer(), obj.Peer(), methodID.Peer(), c_argsCopy)
	// TODO: Exception Handling
	return GoJObject(result)
}

// NewString is a middle-level API.
// Use JStringFromGoString{,E} from jstring.go for general purposes.
func (env JNIEnv) NewString(s string) (str JString, err error) {
	// Encoding to UTF-16
	runes := []rune(s)
	unicode := utf16.Encode(runes)

	// Slice to pointer conversion
	header := (*reflect.SliceHeader)(unsafe.Pointer(&unicode))
	c_unicode := (*C.jchar)(unsafe.Pointer(header.Data))
	c_len := C.jsize(header.Len)

	// Call and wrap
	result := NewString(env.Peer(), c_unicode, c_len)
	str = GoJString(result)

	// NilReturnValueError
	if result == nil {
		err = fmt.Errorf("The string cannot be constructed")
	}

	// TODO: Deal with OutOfMemoryError
	return
}

// GetStringLength is a middle-level API.
// Use JString.Len{,E} from jstring.go for general purposes.
func (env JNIEnv) GetStringLength(str JString) int {
	return GoIntFromSize(GetStringLength(env.Peer(), str.Peer()))
}

// GetStringChars is a middle-level API.
// Use JString.WithChars{,E} from jstring.go for general purposes.
func (env JNIEnv) GetStringChars(str JString) (chars []uint16, isCopy bool, err error) {
	var c_isCopy C.jboolean

	// GetStringLength()
	_len := env.GetStringLength(str)

	// GetStringChars()
	c_chars := GetStringChars(env.Peer(), str.Peer(), &c_isCopy)

	// Pointer to slice conversion
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(c_chars)),
		Len:  _len,
		Cap:  _len,
	}
	chars = *(*[]uint16)(unsafe.Pointer(&header))

	// Boolean conversion
	isCopy = GoBool(c_isCopy)

	// NilReturnValueError
	if chars == nil {
		err = fmt.Errorf("The operation failed")
	}
	return
}

// ReleaseStringChars is a middle-level API.
// Use JString.WithChars{,E} from jstring.go for general purposes.
func (env JNIEnv) ReleaseStringChars(str JString, chars []uint16) {
	// Slice to pointer conversion
	header := (*reflect.SliceHeader)(unsafe.Pointer(&chars))
	c_chars := (*C.jchar)(unsafe.Pointer(header.Data))

	// Clear header data (TODO: is this needed?)
	header.Data = uintptr(0)
	header.Len = 0
	header.Cap = 0

	ReleaseStringChars(env.Peer(), str.Peer(), c_chars)
}

// NewStringUTF is a middle-level API.
// In JNI, this function allows the caller to provide a "modified UTF-8" buffer,
// instead of a UTF-16 buffer as in NewString.
// However, this makes no difference in this wrapper,
// since the parameter type is a Go string.
// As a result, this duplicats the functionality of NewString.
// This function exists for completeness only.
// Use JStringFromGoString{,E} from jstring.go for general purposes.
func (env JNIEnv) NewStringUTF(s string) (str JString, err error) {
	// Encoding to "modified UTF-8"
	// len(s) is a good estimation of the length of utf.
	// If there is no embedded null characters, len(s) the actual length of utf.
	utf := make([]byte, 0, len(s))
	utfIndex := 0
	for sIndex, sLen, size := 0, len(s), 0; sIndex < sLen; sIndex += size {
		r, size := utf8.DecodeRuneInString(s[sIndex:])
		if r != 0 {
			utf = append(utf, s[sIndex:sIndex+size]...)
			utfIndex += size
		} else {
			// Embedded null character for "modified UTF-8"
			utf = append(utf, 0xC0, 0x80)
			utfIndex += 2
		}
	}

	// Slice to pointer conversion
	header := (*reflect.SliceHeader)(unsafe.Pointer(&utf))
	c_utf := (*C.char)(unsafe.Pointer(header.Data))

	// Call and wrap
	result := NewStringUTF(env.Peer(), c_utf)
	str = GoJString(result)

	// NilReturnValueError
	if result == nil {
		err = fmt.Errorf("The string cannot be constructed")
	}

	// TODO: Deal with OutOfMemoryError
	return
}

// GetStringUTFLength is a middle-level API.
// Use JString.UTFLen{,E} from jstring.go for general purposes.
func (env JNIEnv) GetStringUTFLength(str JString) int {
	return GoIntFromSize(GetStringUTFLength(env.Peer(), str.Peer()))
}

// GetStringUTFChars is a middle-level API.
// Use JString.WithUTFChars{,E} from jstring.go for general purposes.
func (env JNIEnv) GetStringUTFChars(str JString) (chars []byte, isCopy bool, err error) {
	var c_isCopy C.jboolean

	// GetStringUTFLength()
	_len := env.GetStringUTFLength(str)

	// GetStringUTFChars()
	c_chars := GetStringUTFChars(env.Peer(), str.Peer(), &c_isCopy)

	// Pointer to slice conversion
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(c_chars)),
		Len:  _len,
		Cap:  _len,
	}
	chars = *((*[]byte)(unsafe.Pointer(&header)))

	// Boolean conversion
	isCopy = GoBool(c_isCopy)

	// NilReturnValueError
	if chars == nil {
		err = fmt.Errorf("The operation failed")
	}
	return
}

// ReleaseStringUTFChars is a middle-level API.
// Use JString.WithUTFChars{,E} from jstring.go for general purposes.
func (env JNIEnv) ReleaseStringUTFChars(str JString, chars []byte) {
	// Slice to pointer conversion
	header := (*reflect.SliceHeader)(unsafe.Pointer(&chars))
	c_chars := (*C.char)(unsafe.Pointer(header.Data))

	// Clear header data (TODO: is this needed?)
	header.Data = uintptr(0)
	header.Len = 0
	header.Cap = 0

	ReleaseStringUTFChars(env.Peer(), str.Peer(), c_chars)
}

// ReleaseStringUTFChars is a middle-level API.
func (env JNIEnv) GetStringRegion(str JString, start, _len int, buf []uint16) error {
	// Slice to pointer conversion
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	c_buf := (*C.jchar)(unsafe.Pointer(header.Data))

	GetStringRegion(env.Peer(), str.Peer(), JavaSize(start), JavaSize(_len), c_buf)

	// TODO: Deal with StringIndexOutOfBoundsException
	return nil
}

// ReleaseStringUTFChars is a middle-level API.
func (env JNIEnv) GetStringUTFRegion(str JString, start, _len int, buf []byte) error {
	// Slice to pointer conversion
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	c_buf := (*C.char)(unsafe.Pointer(header.Data))

	GetStringUTFRegion(env.Peer(), str.Peer(), JavaSize(start), JavaSize(_len), c_buf)

	// TODO: Deal with StringIndexOutOfBoundsException
	return nil
}

// ReleaseStringUTFChars is a middle-level API.
func (env JNIEnv) GetStringCritical(str JString) (cstring []uint16, isCopy bool, err error) {
	var c_isCopy C.jboolean
	c_cstring := GetStringCritical(env.Peer(), str.Peer(), &c_isCopy)

	// GetStringUTFLength()
	_len := env.GetStringUTFLength(str)

	// Pointer to slice conversion
	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(c_cstring)),
		Len:  _len,
		Cap:  _len,
	}
	cstring = *((*[]uint16)(unsafe.Pointer(&header)))

	isCopy = GoBool(c_isCopy)
	return
}

// ReleaseStringUTFChars is a middle-level API.
func (env JNIEnv) ReleaseStringCritical(str JString, cstring []uint16) {
	// Slice to pointer conversion
	header := (*reflect.SliceHeader)(unsafe.Pointer(&cstring))
	c_cstring := (*C.jchar)(unsafe.Pointer(header.Data))

	// Clear header data (TODO: is this needed?)
	header.Data = uintptr(0)
	header.Len = 0
	header.Cap = 0

	ReleaseStringCritical(env.Peer(), str.Peer(), c_cstring)
}

func (env JNIEnv) ExceptionCheck() bool {
	return GoBool(ExceptionCheck(env.Peer()))
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
type Version int32

const (
	// To prevent the values from not being defined in C side,
	// actual values are used.
	JNI_VERSION_1_1 Version = iota + 0x00010001
	JNI_VERSION_1_2
	_
	JNI_VERSION_1_4
	_
	JNI_VERSION_1_6
	_
	JNI_VERSION_1_8
)

//go:generate stringer -type=Version
