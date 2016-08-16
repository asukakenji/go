package main

// #include "jnienv.h"
import "C"

import (
	"unsafe"
)

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

// jni.h:
//     jint (JNICALL *Throw)(JNIEnv *env, jthrowable obj);
func Throw(env *C.JNIEnv, obj C.jthrowable) C.jint {
	return C._GoJniThrow(env, obj)
}

// jni.h:
//     jint (JNICALL *ThrowNew)(JNIEnv *env, jclass clazz, const char *msg);
func ThrowNew(env *C.JNIEnv, clazz C.jclass, msg *C.char) C.jint {
	return C._GoJniThrowNew(env, clazz, msg)
}

// jni.h:
//     jthrowable (JNICALL *ExceptionOccurred)(JNIEnv *env);
func ExceptionOccurred(env *C.JNIEnv) C.jthrowable {
	return C._GoJniExceptionOccurred(env)
}

// jni.h:
//     void (JNICALL *ExceptionDescribe)(JNIEnv *env);
func ExceptionDescribe(env *C.JNIEnv) {
	C._GoJniExceptionDescribe(env)
}

// jni.h:
//     void (JNICALL *ExceptionClear)(JNIEnv *env);
func ExceptionClear(env *C.JNIEnv) {
	C._GoJniExceptionClear(env)
}

// jni.h:
//     void (JNICALL *FatalError)(JNIEnv *env, const char *msg);
func FatalError(env *C.JNIEnv, msg *C.char) {
	C._GoJniFatalError(env, msg)
}

// jni.h:
//     jint (JNICALL *PushLocalFrame)(JNIEnv *env, jint capacity);
func PushLocalFrame(env *C.JNIEnv, capacity C.jint) C.jint {
	return C._GoJniPushLocalFrame(env, capacity)
}

// jni.h:
//     jobject (JNICALL *PopLocalFrame)(JNIEnv *env, jobject result);
func PopLocalFrame(env *C.JNIEnv, result C.jobject) C.jobject {
	return C._GoJniPopLocalFrame(env, result)
}

// jni.h:
//     jobject (JNICALL *NewGlobalRef)(JNIEnv *env, jobject lobj);
func NewGlobalRef(env *C.JNIEnv, lobj C.jobject) C.jobject {
	return C._GoJniNewGlobalRef(env, lobj)
}

// jni.h:
//     void (JNICALL *DeleteGlobalRef)(JNIEnv *env, jobject gref);
func DeleteGlobalRef(env *C.JNIEnv, gref C.jobject) {
	C._GoJniDeleteGlobalRef(env, gref)
}

// jni.h:
//     void (JNICALL *DeleteLocalRef)(JNIEnv *env, jobject obj);
func DeleteLocalRef(env *C.JNIEnv, obj C.jobject) {
	C._GoJniDeleteLocalRef(env, obj)
}

// jni.h:
//     jboolean (JNICALL *IsSameObject)(JNIEnv *env, jobject obj1, jobject obj2);
func IsSameObject(env *C.JNIEnv, obj1, obj2 C.jobject) C.jboolean {
	return C._GoJniIsSameObject(env, obj1, obj2)
}

// jni.h:
//     jobject (JNICALL *NewLocalRef)(JNIEnv *env, jobject ref);
func NewLocalRef(env *C.JNIEnv, ref C.jobject) C.jobject {
	return C._GoJniNewLocalRef(env, ref)
}

// jni.h:
//     jint (JNICALL *EnsureLocalCapacity)(JNIEnv *env, jint capacity);
func EnsureLocalCapacity(env *C.JNIEnv, capacity C.jint) C.jint {
	return C._GoJniEnsureLocalCapacity(env, capacity)
}

// jni.h:
//     jobject (JNICALL *AllocObject)(JNIEnv *env, jclass clazz);
func AllocObject(env *C.JNIEnv, clazz C.jclass) C.jobject {
	return C._GoJniAllocObject(env, clazz)
}

// jni.h:
//     jobject (JNICALL *NewObject)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jobject (JNICALL *NewObjectV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jobject (JNICALL *NewObjectA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
func NewObjectA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jobject {
	return C._GoJniNewObjectA(env, clazz, methodID, args)
}

// jni.h:
//     jclass (JNICALL *GetObjectClass)(JNIEnv *env, jobject obj);
func GetObjectClass(env *C.JNIEnv, obj C.jobject) C.jclass {
	return C._GoJniGetObjectClass(env, obj)
}

// jni.h:
//     jboolean (JNICALL *IsInstanceOf)(JNIEnv *env, jobject obj, jclass clazz);
func IsInstanceOf(env *C.JNIEnv, obj C.jobject, clazz C.jclass) C.jboolean {
	return C._GoJniIsInstanceOf(env, obj, clazz)
}

// jni.h:
//     jmethodID (JNICALL *GetMethodID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);
func GetMethodID(env *C.JNIEnv, clazz C.jclass, name, sig *C.char) C.jmethodID {
	return C._GoJniGetMethodID(env, clazz, name, sig)
}

// jni.h:
//     jobject (JNICALL *CallObjectMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jobject (JNICALL *CallObjectMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jobject (JNICALL *CallObjectMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);
func CallObjectMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jobject {
	return C._GoJniCallObjectMethodA(env, obj, methodID, args)
}

// jni.h:
//     jboolean (JNICALL *CallBooleanMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jboolean (JNICALL *CallBooleanMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jboolean (JNICALL *CallBooleanMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);
func CallBooleanMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jboolean {
	return C._GoJniCallBooleanMethodA(env, obj, methodID, args)
}

// jni.h:
//     jbyte (JNICALL *CallByteMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jbyte (JNICALL *CallByteMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jbyte (JNICALL *CallByteMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
func CallByteMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jbyte {
	return C._GoJniCallByteMethodA(env, obj, methodID, args)
}

// jni.h:
//     jchar (JNICALL *CallCharMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jchar (JNICALL *CallCharMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jchar (JNICALL *CallCharMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
func CallCharMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jchar {
	return C._GoJniCallCharMethodA(env, obj, methodID, args)
}

// jni.h:
//     jshort (JNICALL *CallShortMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jshort (JNICALL *CallShortMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jshort (JNICALL *CallShortMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
func CallShortMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jshort {
	return C._GoJniCallShortMethodA(env, obj, methodID, args)
}

// jni.h:
//     jint (JNICALL *CallIntMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jint (JNICALL *CallIntMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jint (JNICALL *CallIntMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
func CallIntMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jint {
	return C._GoJniCallIntMethodA(env, obj, methodID, args)
}

// jni.h:
//     jlong (JNICALL *CallLongMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jlong (JNICALL *CallLongMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jlong (JNICALL *CallLongMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
func CallLongMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jlong {
	return C._GoJniCallLongMethodA(env, obj, methodID, args)
}

// jni.h:
//     jfloat (JNICALL *CallFloatMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jfloat (JNICALL *CallFloatMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jfloat (JNICALL *CallFloatMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
func CallFloatMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jfloat {
	return C._GoJniCallFloatMethodA(env, obj, methodID, args)
}

// jni.h:
//     jdouble (JNICALL *CallDoubleMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jdouble (JNICALL *CallDoubleMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jdouble (JNICALL *CallDoubleMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
func CallDoubleMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) C.jdouble {
	return C._GoJniCallDoubleMethodA(env, obj, methodID, args)
}

// jni.h:
//     void (JNICALL *CallVoidMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     void (JNICALL *CallVoidMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     void (JNICALL *CallVoidMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);
func CallVoidMethodA(env *C.JNIEnv, obj C.jobject, methodID C.jmethodID, args *C.jvalue) {
	C._GoJniCallVoidMethodA(env, obj, methodID, args)
}

// jni.h:
//     jobject (JNICALL *CallNonvirtualObjectMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jobject (JNICALL *CallNonvirtualObjectMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jobject (JNICALL *CallNonvirtualObjectMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);
func CallNonvirtualObjectMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jobject {
	return C._GoJniCallNonvirtualObjectMethodA(env, obj, clazz, methodID, args)
}

// jni.h:
//     jboolean (JNICALL *CallNonvirtualBooleanMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jboolean (JNICALL *CallNonvirtualBooleanMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jboolean (JNICALL *CallNonvirtualBooleanMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);
func CallNonvirtualBooleanMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jboolean {
	return C._GoJniCallNonvirtualBooleanMethodA(env, obj, clazz, methodID, args)
}

// jni.h:
//     jbyte (JNICALL *CallNonvirtualByteMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jbyte (JNICALL *CallNonvirtualByteMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jbyte (JNICALL *CallNonvirtualByteMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
func CallNonvirtualByteMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jbyte {
	return C._GoJniCallNonvirtualByteMethodA(env, obj, clazz, methodID, args)
}

// jni.h:
//     jchar (JNICALL *CallNonvirtualCharMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jchar (JNICALL *CallNonvirtualCharMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jchar (JNICALL *CallNonvirtualCharMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
func CallNonvirtualCharMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jchar {
	return C._GoJniCallNonvirtualCharMethodA(env, obj, clazz, methodID, args)
}

// jni.h:
//     jshort (JNICALL *CallNonvirtualShortMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jshort (JNICALL *CallNonvirtualShortMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jshort (JNICALL *CallNonvirtualShortMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
func CallNonvirtualShortMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jshort {
	return C._GoJniCallNonvirtualShortMethodA(env, obj, clazz, methodID, args)
}

// jni.h:
//     jint (JNICALL *CallNonvirtualIntMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jint (JNICALL *CallNonvirtualIntMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jint (JNICALL *CallNonvirtualIntMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
func CallNonvirtualIntMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jint {
	return C._GoJniCallNonvirtualIntMethodA(env, obj, clazz, methodID, args)
}

// jni.h:
//     jlong (JNICALL *CallNonvirtualLongMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jlong (JNICALL *CallNonvirtualLongMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jlong (JNICALL *CallNonvirtualLongMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
func CallNonvirtualLongMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jlong {
	return C._GoJniCallNonvirtualLongMethodA(env, obj, clazz, methodID, args)
}

// jni.h:
//     jfloat (JNICALL *CallNonvirtualFloatMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jfloat (JNICALL *CallNonvirtualFloatMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jfloat (JNICALL *CallNonvirtualFloatMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
func CallNonvirtualFloatMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jfloat {
	return C._GoJniCallNonvirtualFloatMethodA(env, obj, clazz, methodID, args)
}

// jni.h:
//     jdouble (JNICALL *CallNonvirtualDoubleMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jdouble (JNICALL *CallNonvirtualDoubleMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jdouble (JNICALL *CallNonvirtualDoubleMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
func CallNonvirtualDoubleMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jdouble {
	return C._GoJniCallNonvirtualDoubleMethodA(env, obj, clazz, methodID, args)
}

// jni.h:
//     void (JNICALL *CallNonvirtualVoidMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     void (JNICALL *CallNonvirtualVoidMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     void (JNICALL *CallNonvirtualVoidMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);
func CallNonvirtualVoidMethodA(env *C.JNIEnv, obj C.jobject, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) {
	C._GoJniCallNonvirtualVoidMethodA(env, obj, clazz, methodID, args)
}

// jni.h:
//     jfieldID (JNICALL *GetFieldID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);
func GetFieldID(env *C.JNIEnv, clazz C.jclass, name, sig *C.char) C.jfieldID {
	return C._GoJniGetFieldID(env, clazz, name, sig)
}

// jni.h:
//     jobject (JNICALL *GetObjectField)(JNIEnv *env, jobject obj, jfieldID fieldID);
func GetObjectField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jobject {
	return C._GoJniGetObjectField(env, obj, fieldID)
}

// jni.h:
//     jboolean (JNICALL *GetBooleanField)(JNIEnv *env, jobject obj, jfieldID fieldID);
func GetBooleanField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jboolean {
	return C._GoJniGetBooleanField(env, obj, fieldID)
}

// jni.h:
//     jbyte (JNICALL *GetByteField)(JNIEnv *env, jobject obj, jfieldID fieldID);
func GetByteField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jbyte {
	return C._GoJniGetByteField(env, obj, fieldID)
}

// jni.h:
//     jchar (JNICALL *GetCharField)(JNIEnv *env, jobject obj, jfieldID fieldID);
func GetCharField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jchar {
	return C._GoJniGetCharField(env, obj, fieldID)
}

// jni.h:
//     jshort (JNICALL *GetShortField)(JNIEnv *env, jobject obj, jfieldID fieldID);
func GetShortField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jshort {
	return C._GoJniGetShortField(env, obj, fieldID)
}

// jni.h:
//     jint (JNICALL *GetIntField)(JNIEnv *env, jobject obj, jfieldID fieldID);
func GetIntField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jint {
	return C._GoJniGetIntField(env, obj, fieldID)
}

// jni.h:
//     jlong (JNICALL *GetLongField)(JNIEnv *env, jobject obj, jfieldID fieldID);
func GetLongField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jlong {
	return C._GoJniGetLongField(env, obj, fieldID)
}

// jni.h:
//     jfloat (JNICALL *GetFloatField)(JNIEnv *env, jobject obj, jfieldID fieldID);
func GetFloatField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jfloat {
	return C._GoJniGetFloatField(env, obj, fieldID)
}

// jni.h:
//     jdouble (JNICALL *GetDoubleField)(JNIEnv *env, jobject obj, jfieldID fieldID);
func GetDoubleField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID) C.jdouble {
	return C._GoJniGetDoubleField(env, obj, fieldID)
}

// jni.h:
//     void (JNICALL *SetObjectField)(JNIEnv *env, jobject obj, jfieldID fieldID, jobject val);
func SetObjectField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jobject) {
	C._GoJniSetObjectField(env, obj, fieldID, val)
}

// jni.h:
//     void (JNICALL *SetBooleanField)(JNIEnv *env, jobject obj, jfieldID fieldID, jboolean val);
func SetBooleanField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jboolean) {
	C._GoJniSetBooleanField(env, obj, fieldID, val)
}

// jni.h:
//     void (JNICALL *SetByteField)(JNIEnv *env, jobject obj, jfieldID fieldID, jbyte val);
func SetByteField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jbyte) {
	C._GoJniSetByteField(env, obj, fieldID, val)
}

// jni.h:
//     void (JNICALL *SetCharField)(JNIEnv *env, jobject obj, jfieldID fieldID, jchar val);
func SetCharField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jchar) {
	C._GoJniSetCharField(env, obj, fieldID, val)
}

// jni.h:
//     void (JNICALL *SetShortField)(JNIEnv *env, jobject obj, jfieldID fieldID, jshort val);
func SetShortField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jshort) {
	C._GoJniSetShortField(env, obj, fieldID, val)
}

// jni.h:
//     void (JNICALL *SetIntField)(JNIEnv *env, jobject obj, jfieldID fieldID, jint val);
func SetIntField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jint) {
	C._GoJniSetIntField(env, obj, fieldID, val)
}

// jni.h:
//     void (JNICALL *SetLongField)(JNIEnv *env, jobject obj, jfieldID fieldID, jlong val);
func SetLongField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jlong) {
	C._GoJniSetLongField(env, obj, fieldID, val)
}

// jni.h:
//     void (JNICALL *SetFloatField)(JNIEnv *env, jobject obj, jfieldID fieldID, jfloat val);
func SetFloatField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jfloat) {
	C._GoJniSetFloatField(env, obj, fieldID, val)
}

// jni.h:
//     void (JNICALL *SetDoubleField)(JNIEnv *env, jobject obj, jfieldID fieldID, jdouble val);
func SetDoubleField(env *C.JNIEnv, obj C.jobject, fieldID C.jfieldID, val C.jdouble) {
	C._GoJniSetDoubleField(env, obj, fieldID, val)
}

// jni.h:
//     jmethodID (JNICALL *GetStaticMethodID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);
func GetStaticMethodID(env *C.JNIEnv, clazz C.jclass, name, sig *C.char) C.jmethodID {
	return C._GoJniGetStaticMethodID(env, clazz, name, sig)
}

// jni.h:
//     jobject (JNICALL *CallStaticObjectMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jobject (JNICALL *CallStaticObjectMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jobject (JNICALL *CallStaticObjectMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
func CallStaticObjectMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jobject {
	return C._GoJniCallStaticObjectMethodA(env, clazz, methodID, args)
}

// jni.h:
//     jboolean (JNICALL *CallStaticBooleanMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jboolean (JNICALL *CallStaticBooleanMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jboolean (JNICALL *CallStaticBooleanMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
func CallStaticBooleanMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jboolean {
	return C._GoJniCallStaticBooleanMethodA(env, clazz, methodID, args)
}

// jni.h:
//     jbyte (JNICALL *CallStaticByteMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jbyte (JNICALL *CallStaticByteMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jbyte (JNICALL *CallStaticByteMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
func CallStaticByteMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jbyte {
	return C._GoJniCallStaticByteMethodA(env, clazz, methodID, args)
}

// jni.h:
//     jchar (JNICALL *CallStaticCharMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jchar (JNICALL *CallStaticCharMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jchar (JNICALL *CallStaticCharMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
func CallStaticCharMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jchar {
	return C._GoJniCallStaticCharMethodA(env, clazz, methodID, args)
}

// jni.h:
//     jshort (JNICALL *CallStaticShortMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jshort (JNICALL *CallStaticShortMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jshort (JNICALL *CallStaticShortMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
func CallStaticShortMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jshort {
	return C._GoJniCallStaticShortMethodA(env, clazz, methodID, args)
}

// jni.h:
//     jint (JNICALL *CallStaticIntMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jint (JNICALL *CallStaticIntMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jint (JNICALL *CallStaticIntMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
func CallStaticIntMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jint {
	return C._GoJniCallStaticIntMethodA(env, clazz, methodID, args)
}

// jni.h:
//     jlong (JNICALL *CallStaticLongMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jlong (JNICALL *CallStaticLongMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jlong (JNICALL *CallStaticLongMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
func CallStaticLongMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jlong {
	return C._GoJniCallStaticLongMethodA(env, clazz, methodID, args)
}

// jni.h:
//     jfloat (JNICALL *CallStaticFloatMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jfloat (JNICALL *CallStaticFloatMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jfloat (JNICALL *CallStaticFloatMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
func CallStaticFloatMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jfloat {
	return C._GoJniCallStaticFloatMethodA(env, clazz, methodID, args)
}

// jni.h:
//     jdouble (JNICALL *CallStaticDoubleMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jdouble (JNICALL *CallStaticDoubleMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jdouble (JNICALL *CallStaticDoubleMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
func CallStaticDoubleMethodA(env *C.JNIEnv, clazz C.jclass, methodID C.jmethodID, args *C.jvalue) C.jdouble {
	return C._GoJniCallStaticDoubleMethodA(env, clazz, methodID, args)
}

// jni.h:
//     void (JNICALL *CallStaticVoidMethod)(JNIEnv *env, jclass cls, jmethodID methodID, ...);
//     void (JNICALL *CallStaticVoidMethodV)(JNIEnv *env, jclass cls, jmethodID methodID, va_list args);
//     void (JNICALL *CallStaticVoidMethodA)(JNIEnv *env, jclass cls, jmethodID methodID, const jvalue * args);
func CallStaticVoidMethodA(env *C.JNIEnv, cls C.jclass, methodID C.jmethodID, args *C.jvalue) {
	C._GoJniCallStaticVoidMethodA(env, cls, methodID, args)
}

// jni.h:
//     jfieldID (JNICALL *GetStaticFieldID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);
func GetStaticFieldID(env *C.JNIEnv, clazz C.jclass, name, sig *C.char) C.jfieldID {
	return C._GoJniGetStaticFieldID(env, clazz, name, sig)
}

// jni.h:
//     jobject (JNICALL *GetStaticObjectField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
func GetStaticObjectField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jobject {
	return C._GoJniGetStaticObjectField(env, clazz, fieldID)
}

// jni.h:
//     jboolean (JNICALL *GetStaticBooleanField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
func GetStaticBooleanField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jboolean {
	return C._GoJniGetStaticBooleanField(env, clazz, fieldID)
}

// jni.h:
//     jbyte (JNICALL *GetStaticByteField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
func GetStaticByteField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jbyte {
	return C._GoJniGetStaticByteField(env, clazz, fieldID)
}

// jni.h:
//     jchar (JNICALL *GetStaticCharField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
func GetStaticCharField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jchar {
	return C._GoJniGetStaticCharField(env, clazz, fieldID)
}

// jni.h:
//     jshort (JNICALL *GetStaticShortField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
func GetStaticShortField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jshort {
	return C._GoJniGetStaticShortField(env, clazz, fieldID)
}

// jni.h:
//     jint (JNICALL *GetStaticIntField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
func GetStaticIntField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jint {
	return C._GoJniGetStaticIntField(env, clazz, fieldID)
}

// jni.h:
//     jlong (JNICALL *GetStaticLongField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
func GetStaticLongField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jlong {
	return C._GoJniGetStaticLongField(env, clazz, fieldID)
}

// jni.h:
//     jfloat (JNICALL *GetStaticFloatField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
func GetStaticFloatField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jfloat {
	return C._GoJniGetStaticFloatField(env, clazz, fieldID)
}

// jni.h:
//     jdouble (JNICALL *GetStaticDoubleField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
func GetStaticDoubleField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID) C.jdouble {
	return C._GoJniGetStaticDoubleField(env, clazz, fieldID)
}

// jni.h:
//     void (JNICALL *SetStaticObjectField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jobject value);
func SetStaticObjectField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jobject) {
	C._GoJniSetStaticObjectField(env, clazz, fieldID, value)
}

// jni.h:
//     void (JNICALL *SetStaticBooleanField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jboolean value);
func SetStaticBooleanField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jboolean) {
	C._GoJniSetStaticBooleanField(env, clazz, fieldID, value)
}

// jni.h:
//     void (JNICALL *SetStaticByteField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jbyte value);
func SetStaticByteField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jbyte) {
	C._GoJniSetStaticByteField(env, clazz, fieldID, value)
}

// jni.h:
//     void (JNICALL *SetStaticCharField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jchar value);
func SetStaticCharField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jchar) {
	C._GoJniSetStaticCharField(env, clazz, fieldID, value)
}

// jni.h:
//     void (JNICALL *SetStaticShortField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jshort value);
func SetStaticShortField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jshort) {
	C._GoJniSetStaticShortField(env, clazz, fieldID, value)
}

// jni.h:
//     void (JNICALL *SetStaticIntField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jint value);
func SetStaticIntField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jint) {
	C._GoJniSetStaticIntField(env, clazz, fieldID, value)
}

// jni.h:
//     void (JNICALL *SetStaticLongField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jlong value);
func SetStaticLongField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jlong) {
	C._GoJniSetStaticLongField(env, clazz, fieldID, value)
}

// jni.h:
//     void (JNICALL *SetStaticFloatField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jfloat value);
func SetStaticFloatField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jfloat) {
	C._GoJniSetStaticFloatField(env, clazz, fieldID, value)
}

// jni.h:
//     void (JNICALL *SetStaticDoubleField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jdouble value);
func SetStaticDoubleField(env *C.JNIEnv, clazz C.jclass, fieldID C.jfieldID, value C.jdouble) {
	C._GoJniSetStaticDoubleField(env, clazz, fieldID, value)
}

// jni.h:
//     jstring (JNICALL *NewString)(JNIEnv *env, const jchar *unicode, jsize len);
func NewString(env *C.JNIEnv, unicode *C.jchar, _len C.jsize) C.jstring {
	return C._GoJniNewString(env, unicode, _len)
}

// jni.h:
//     jsize (JNICALL *GetStringLength)(JNIEnv *env, jstring str);
func GetStringLength(env *C.JNIEnv, str C.jstring) C.jsize {
	return C._GoJniGetStringLength(env, str)
}

// jni.h:
//     const jchar *(JNICALL *GetStringChars)(JNIEnv *env, jstring str, jboolean *isCopy);
func GetStringChars(env *C.JNIEnv, str C.jstring, isCopy *C.jboolean) *C.jchar {
	return C._GoJniGetStringChars(env, str, isCopy)
}

// jni.h:
//     void (JNICALL *ReleaseStringChars)(JNIEnv *env, jstring str, const jchar *chars);
func ReleaseStringChars(env *C.JNIEnv, str C.jstring, chars *C.jchar) {
	C._GoJniReleaseStringChars(env, str, chars)
}

// jni.h:
//     jstring (JNICALL *NewStringUTF)(JNIEnv *env, const char *utf);
func NewStringUTF(env *C.JNIEnv, utf *C.char) C.jstring {
	return C._GoJniNewStringUTF(env, utf)
}

// jni.h:
//     jsize (JNICALL *GetStringUTFLength)(JNIEnv *env, jstring str);
func GetStringUTFLength(env *C.JNIEnv, str C.jstring) C.jsize {
	return C._GoJniGetStringUTFLength(env, str)
}

// jni.h:
//     const char* (JNICALL *GetStringUTFChars)(JNIEnv *env, jstring str, jboolean *isCopy);
func GetStringUTFChars(env *C.JNIEnv, str C.jstring, isCopy *C.jboolean) *C.char {
	return C._GoJniGetStringUTFChars(env, str, isCopy)
}

// jni.h:
//     void (JNICALL *ReleaseStringUTFChars)(JNIEnv *env, jstring str, const char* chars);
func ReleaseStringUTFChars(env *C.JNIEnv, str C.jstring, chars *C.char) {
	C._GoJniReleaseStringUTFChars(env, str, chars)
}

// jni.h:
//     jsize (JNICALL *GetArrayLength)(JNIEnv *env, jarray array);
func GetArrayLength(env *C.JNIEnv, array C.jarray) C.jsize {
	return C._GoJniGetArrayLength(env, array)
}

// jni.h:
//     jobjectArray (JNICALL *NewObjectArray)(JNIEnv *env, jsize len, jclass clazz, jobject init);
func NewObjectArray(env *C.JNIEnv, _len C.jsize, clazz C.jclass, init C.jobject) C.jobjectArray {
	return C._GoJniNewObjectArray(env, _len, clazz, init)
}

// jni.h:
//     jobject (JNICALL *GetObjectArrayElement)(JNIEnv *env, jobjectArray array, jsize index);
func GetObjectArrayElement(env *C.JNIEnv, array C.jobjectArray, index C.jsize) C.jobject {
	return C._GoJniGetObjectArrayElement(env, array, index)
}

// jni.h:
//     void (JNICALL *SetObjectArrayElement)(JNIEnv *env, jobjectArray array, jsize index, jobject val);
func SetObjectArrayElement(env *C.JNIEnv, array C.jobjectArray, index C.jsize, val C.jobject) {
	C._GoJniSetObjectArrayElement(env, array, index, val)
}

// jni.h:
//     jbooleanArray (JNICALL *NewBooleanArray)(JNIEnv *env, jsize len);
func NewBooleanArray(env *C.JNIEnv, _len C.jsize) C.jbooleanArray {
	return C._GoJniNewBooleanArray(env, _len)
}

// jni.h:
//     jbyteArray (JNICALL *NewByteArray)(JNIEnv *env, jsize len);
func NewByteArray(env *C.JNIEnv, _len C.jsize) C.jbyteArray {
	return C._GoJniNewByteArray(env, _len)
}

// jni.h:
//     jcharArray (JNICALL *NewCharArray)(JNIEnv *env, jsize len);
func NewCharArray(env *C.JNIEnv, _len C.jsize) C.jcharArray {
	return C._GoJniNewCharArray(env, _len)
}

// jni.h:
//     jshortArray (JNICALL *NewShortArray)(JNIEnv *env, jsize len);
func NewShortArray(env *C.JNIEnv, _len C.jsize) C.jshortArray {
	return C._GoJniNewShortArray(env, _len)
}

// jni.h:
//     jintArray (JNICALL *NewIntArray)(JNIEnv *env, jsize len);
func NewIntArray(env *C.JNIEnv, _len C.jsize) C.jintArray {
	return C._GoJniNewIntArray(env, _len)
}

// jni.h:
//     jlongArray (JNICALL *NewLongArray)(JNIEnv *env, jsize len);
func NewLongArray(env *C.JNIEnv, _len C.jsize) C.jlongArray {
	return C._GoJniNewLongArray(env, _len)
}

// jni.h:
//     jfloatArray (JNICALL *NewFloatArray)(JNIEnv *env, jsize len);
func NewFloatArray(env *C.JNIEnv, _len C.jsize) C.jfloatArray {
	return C._GoJniNewFloatArray(env, _len)
}

// jni.h:
//     jdoubleArray (JNICALL *NewDoubleArray)(JNIEnv *env, jsize len);
func NewDoubleArray(env *C.JNIEnv, _len C.jsize) C.jdoubleArray {
	return C._GoJniNewDoubleArray(env, _len)
}

// jni.h:
//     jboolean * (JNICALL *GetBooleanArrayElements)(JNIEnv *env, jbooleanArray array, jboolean *isCopy);
func GetBooleanArrayElements(env *C.JNIEnv, array C.jbooleanArray, isCopy *C.jboolean) *C.jboolean {
	return C._GoJniGetBooleanArrayElements(env, array, isCopy)
}

// jni.h:
//     jbyte * (JNICALL *GetByteArrayElements)(JNIEnv *env, jbyteArray array, jboolean *isCopy);
func GetByteArrayElements(env *C.JNIEnv, array C.jbyteArray, isCopy *C.jboolean) *C.jbyte {
	return C._GoJniGetByteArrayElements(env, array, isCopy)
}

// jni.h:
//     jchar * (JNICALL *GetCharArrayElements)(JNIEnv *env, jcharArray array, jboolean *isCopy);
func GetCharArrayElements(env *C.JNIEnv, array C.jcharArray, isCopy *C.jboolean) *C.jchar {
	return C._GoJniGetCharArrayElements(env, array, isCopy)
}

// jni.h:
//     jshort * (JNICALL *GetShortArrayElements)(JNIEnv *env, jshortArray array, jboolean *isCopy);
func GetShortArrayElements(env *C.JNIEnv, array C.jshortArray, isCopy *C.jboolean) *C.jshort {
	return C._GoJniGetShortArrayElements(env, array, isCopy)
}

// jni.h:
//     jint * (JNICALL *GetIntArrayElements)(JNIEnv *env, jintArray array, jboolean *isCopy);
func GetIntArrayElements(env *C.JNIEnv, array C.jintArray, isCopy *C.jboolean) *C.jint {
	return C._GoJniGetIntArrayElements(env, array, isCopy)
}

// jni.h:
//     jlong * (JNICALL *GetLongArrayElements)(JNIEnv *env, jlongArray array, jboolean *isCopy);
func GetLongArrayElements(env *C.JNIEnv, array C.jlongArray, isCopy *C.jboolean) *C.jlong {
	return C._GoJniGetLongArrayElements(env, array, isCopy)
}

// jni.h:
//     jfloat * (JNICALL *GetFloatArrayElements)(JNIEnv *env, jfloatArray array, jboolean *isCopy);
func GetFloatArrayElements(env *C.JNIEnv, array C.jfloatArray, isCopy *C.jboolean) *C.jfloat {
	return C._GoJniGetFloatArrayElements(env, array, isCopy)
}

// jni.h:
//     jdouble * (JNICALL *GetDoubleArrayElements)(JNIEnv *env, jdoubleArray array, jboolean *isCopy);
func GetDoubleArrayElements(env *C.JNIEnv, array C.jdoubleArray, isCopy *C.jboolean) *C.jdouble {
	return C._GoJniGetDoubleArrayElements(env, array, isCopy)
}

// jni.h:
//     void (JNICALL *ReleaseBooleanArrayElements)(JNIEnv *env, jbooleanArray array, jboolean *elems, jint mode);
func ReleaseBooleanArrayElements(env *C.JNIEnv, array C.jbooleanArray, elems *C.jboolean, mode C.jint) {
	C._GoJniReleaseBooleanArrayElements(env, array, elems, mode)
}

// jni.h:
//     void (JNICALL *ReleaseByteArrayElements)(JNIEnv *env, jbyteArray array, jbyte *elems, jint mode);
func ReleaseByteArrayElements(env *C.JNIEnv, array C.jbyteArray, elems *C.jbyte, mode C.jint) {
	C._GoJniReleaseByteArrayElements(env, array, elems, mode)
}

// jni.h:
//     void (JNICALL *ReleaseCharArrayElements)(JNIEnv *env, jcharArray array, jchar *elems, jint mode);
func ReleaseCharArrayElements(env *C.JNIEnv, array C.jcharArray, elems *C.jchar, mode C.jint) {
	C._GoJniReleaseCharArrayElements(env, array, elems, mode)
}

// jni.h:
//     void (JNICALL *ReleaseShortArrayElements)(JNIEnv *env, jshortArray array, jshort *elems, jint mode);
func ReleaseShortArrayElements(env *C.JNIEnv, array C.jshortArray, elems *C.jshort, mode C.jint) {
	C._GoJniReleaseShortArrayElements(env, array, elems, mode)
}

// jni.h:
//     void (JNICALL *ReleaseIntArrayElements)(JNIEnv *env, jintArray array, jint *elems, jint mode);
func ReleaseIntArrayElements(env *C.JNIEnv, array C.jintArray, elems *C.jint, mode C.jint) {
	C._GoJniReleaseIntArrayElements(env, array, elems, mode)
}

// jni.h:
//     void (JNICALL *ReleaseLongArrayElements)(JNIEnv *env, jlongArray array, jlong *elems, jint mode);
func ReleaseLongArrayElements(env *C.JNIEnv, array C.jlongArray, elems *C.jlong, mode C.jint) {
	C._GoJniReleaseLongArrayElements(env, array, elems, mode)
}

// jni.h:
//     void (JNICALL *ReleaseFloatArrayElements)(JNIEnv *env, jfloatArray array, jfloat *elems, jint mode);
func ReleaseFloatArrayElements(env *C.JNIEnv, array C.jfloatArray, elems *C.jfloat, mode C.jint) {
	C._GoJniReleaseFloatArrayElements(env, array, elems, mode)
}

// jni.h:
//     void (JNICALL *ReleaseDoubleArrayElements)(JNIEnv *env, jdoubleArray array, jdouble *elems, jint mode);
func ReleaseDoubleArrayElements(env *C.JNIEnv, array C.jdoubleArray, elems *C.jdouble, mode C.jint) {
	C._GoJniReleaseDoubleArrayElements(env, array, elems, mode)
}

// jni.h:
//     void (JNICALL *GetBooleanArrayRegion)(JNIEnv *env, jbooleanArray array, jsize start, jsize l, jboolean *buf);
func GetBooleanArrayRegion(env *C.JNIEnv, array C.jbooleanArray, start, l C.jsize, buf *C.jboolean) {
	C._GoJniGetBooleanArrayRegion(env, array, start, l, buf)
}

// jni.h:
//     void (JNICALL *GetByteArrayRegion)(JNIEnv *env, jbyteArray array, jsize start, jsize len, jbyte *buf);
func GetByteArrayRegion(env *C.JNIEnv, array C.jbyteArray, start, _len C.jsize, buf *C.jbyte) {
	C._GoJniGetByteArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     void (JNICALL *GetCharArrayRegion)(JNIEnv *env, jcharArray array, jsize start, jsize len, jchar *buf);
func GetCharArrayRegion(env *C.JNIEnv, array C.jcharArray, start, _len C.jsize, buf *C.jchar) {
	C._GoJniGetCharArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     void (JNICALL *GetShortArrayRegion)(JNIEnv *env, jshortArray array, jsize start, jsize len, jshort *buf);
func GetShortArrayRegion(env *C.JNIEnv, array C.jshortArray, start, _len C.jsize, buf *C.jshort) {
	C._GoJniGetShortArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     void (JNICALL *GetIntArrayRegion)(JNIEnv *env, jintArray array, jsize start, jsize len, jint *buf);
func GetIntArrayRegion(env *C.JNIEnv, array C.jintArray, start, _len C.jsize, buf *C.jint) {
	C._GoJniGetIntArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     void (JNICALL *GetLongArrayRegion)(JNIEnv *env, jlongArray array, jsize start, jsize len, jlong *buf);
func GetLongArrayRegion(env *C.JNIEnv, array C.jlongArray, start, _len C.jsize, buf *C.jlong) {
	C._GoJniGetLongArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     void (JNICALL *GetFloatArrayRegion)(JNIEnv *env, jfloatArray array, jsize start, jsize len, jfloat *buf);
func GetFloatArrayRegion(env *C.JNIEnv, array C.jfloatArray, start, _len C.jsize, buf *C.jfloat) {
	C._GoJniGetFloatArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     void (JNICALL *GetDoubleArrayRegion)(JNIEnv *env, jdoubleArray array, jsize start, jsize len, jdouble *buf);
func GetDoubleArrayRegion(env *C.JNIEnv, array C.jdoubleArray, start, _len C.jsize, buf *C.jdouble) {
	C._GoJniGetDoubleArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     void (JNICALL *SetBooleanArrayRegion)(JNIEnv *env, jbooleanArray array, jsize start, jsize l, const jboolean *buf);
func SetBooleanArrayRegion(env *C.JNIEnv, array C.jbooleanArray, start, l C.jsize, buf *C.jboolean) {
	C._GoJniSetBooleanArrayRegion(env, array, start, l, buf)
}

// jni.h:
//     void (JNICALL *SetByteArrayRegion)(JNIEnv *env, jbyteArray array, jsize start, jsize len, const jbyte *buf);
func SetByteArrayRegion(env *C.JNIEnv, array C.jbyteArray, start, _len C.jsize, buf *C.jbyte) {
	C._GoJniSetByteArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     void (JNICALL *SetCharArrayRegion)(JNIEnv *env, jcharArray array, jsize start, jsize len, const jchar *buf);
func SetCharArrayRegion(env *C.JNIEnv, array C.jcharArray, start, _len C.jsize, buf *C.jchar) {
	C._GoJniSetCharArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     void (JNICALL *SetShortArrayRegion)(JNIEnv *env, jshortArray array, jsize start, jsize len, const jshort *buf);
func SetShortArrayRegion(env *C.JNIEnv, array C.jshortArray, start, _len C.jsize, buf *C.jshort) {
	C._GoJniSetShortArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     void (JNICALL *SetIntArrayRegion)(JNIEnv *env, jintArray array, jsize start, jsize len, const jint *buf);
func SetIntArrayRegion(env *C.JNIEnv, array C.jintArray, start, _len C.jsize, buf *C.jint) {
	C._GoJniSetIntArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     void (JNICALL *SetLongArrayRegion)(JNIEnv *env, jlongArray array, jsize start, jsize len, const jlong *buf);
func SetLongArrayRegion(env *C.JNIEnv, array C.jlongArray, start, _len C.jsize, buf *C.jlong) {
	C._GoJniSetLongArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     void (JNICALL *SetFloatArrayRegion)(JNIEnv *env, jfloatArray array, jsize start, jsize len, const jfloat *buf);
func SetFloatArrayRegion(env *C.JNIEnv, array C.jfloatArray, start, _len C.jsize, buf *C.jfloat) {
	C._GoJniSetFloatArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     void (JNICALL *SetDoubleArrayRegion)(JNIEnv *env, jdoubleArray array, jsize start, jsize len, const jdouble *buf);
func SetDoubleArrayRegion(env *C.JNIEnv, array C.jdoubleArray, start, _len C.jsize, buf *C.jdouble) {
	C._GoJniSetDoubleArrayRegion(env, array, start, _len, buf)
}

// jni.h:
//     jint (JNICALL *RegisterNatives)(JNIEnv *env, jclass clazz, const JNINativeMethod *methods, jint nMethods);
func RegisterNatives(env *C.JNIEnv, clazz C.jclass, methods *C.JNINativeMethod, nMethods C.jint) C.jint {
	return C._GoJniRegisterNatives(env, clazz, methods, nMethods)
}

// jni.h:
//     jint (JNICALL *UnregisterNatives)(JNIEnv *env, jclass clazz);
func UnregisterNatives(env *C.JNIEnv, clazz C.jclass) C.jint {
	return C._GoJniUnregisterNatives(env, clazz)
}

// jni.h:
//     jint (JNICALL *MonitorEnter)(JNIEnv *env, jobject obj);
func MonitorEnter(env *C.JNIEnv, obj C.jobject) C.jint {
	return C._GoJniMonitorEnter(env, obj)
}

// jni.h:
//     jint (JNICALL *MonitorExit)(JNIEnv *env, jobject obj);
func MonitorExit(env *C.JNIEnv, obj C.jobject) C.jint {
	return C._GoJniMonitorExit(env, obj)
}

// jni.h:
//     jint (JNICALL *GetJavaVM)(JNIEnv *env, JavaVM **vm);
func GetJavaVM(env *C.JNIEnv, vm **C.JavaVM) C.jint {
	return C._GoJniGetJavaVM(env, vm)
}

// jni.h:
//     void (JNICALL *GetStringRegion)(JNIEnv *env, jstring str, jsize start, jsize len, jchar *buf);
func GetStringRegion(env *C.JNIEnv, str C.jstring, start, _len C.jsize, buf *C.jchar) {
	C._GoJniGetStringRegion(env, str, start, _len, buf)
}

// jni.h:
//     void (JNICALL *GetStringUTFRegion)(JNIEnv *env, jstring str, jsize start, jsize len, char *buf);
func GetStringUTFRegion(env *C.JNIEnv, str C.jstring, start, _len C.jsize, buf *C.char) {
	C._GoJniGetStringUTFRegion(env, str, start, _len, buf)
}

// jni.h:
//     void * (JNICALL *GetPrimitiveArrayCritical)(JNIEnv *env, jarray array, jboolean *isCopy);
func GetPrimitiveArrayCritical(env *C.JNIEnv, array C.jarray, isCopy *C.jboolean) unsafe.Pointer {
	return C._GoJniGetPrimitiveArrayCritical(env, array, isCopy)
}

// jni.h:
//     void (JNICALL *ReleasePrimitiveArrayCritical)(JNIEnv *env, jarray array, void *carray, jint mode);
func ReleasePrimitiveArrayCritical(env *C.JNIEnv, array C.jarray, carray unsafe.Pointer, mode C.jint) {
	C._GoJniReleasePrimitiveArrayCritical(env, array, carray, mode)
}

// jni.h:
//     const jchar * (JNICALL *GetStringCritical)(JNIEnv *env, jstring string, jboolean *isCopy);
func GetStringCritical(env *C.JNIEnv, string C.jstring, isCopy *C.jboolean) *C.jchar {
	return C._GoJniGetStringCritical(env, string, isCopy)
}

// jni.h:
//     void (JNICALL *ReleaseStringCritical)(JNIEnv *env, jstring string, const jchar *cstring);
func ReleaseStringCritical(env *C.JNIEnv, string C.jstring, cstring *C.jchar) {
	C._GoJniReleaseStringCritical(env, string, cstring)
}

// jni.h:
//     jweak (JNICALL *NewWeakGlobalRef)(JNIEnv *env, jobject obj);
func NewWeakGlobalRef(env *C.JNIEnv, obj C.jobject) C.jweak {
	return C._GoJniNewWeakGlobalRef(env, obj)
}

// jni.h:
//     void (JNICALL *DeleteWeakGlobalRef)(JNIEnv *env, jweak ref);
func DeleteWeakGlobalRef(env *C.JNIEnv, ref C.jweak) {
	C._GoJniDeleteWeakGlobalRef(env, ref)
}

// jni.h:
//     jboolean (JNICALL *ExceptionCheck)(JNIEnv *env);
func ExceptionCheck(env *C.JNIEnv) C.jboolean {
	return C._GoJniExceptionCheck(env)
}

// jni.h:
//     jobject (JNICALL *NewDirectByteBuffer)(JNIEnv* env, void* address, jlong capacity);
func NewDirectByteBuffer(env *C.JNIEnv, address unsafe.Pointer, capacity C.jlong) C.jobject {
	return C._GoJniNewDirectByteBuffer(env, address, capacity)
}

// jni.h:
//     void* (JNICALL *GetDirectBufferAddress)(JNIEnv* env, jobject buf);
func GetDirectBufferAddress(env *C.JNIEnv, buf C.jobject) unsafe.Pointer {
	return C._GoJniGetDirectBufferAddress(env, buf)
}

// jni.h:
//     jlong (JNICALL *GetDirectBufferCapacity)(JNIEnv* env, jobject buf);
func GetDirectBufferCapacity(env *C.JNIEnv, buf C.jobject) C.jlong {
	return C._GoJniGetDirectBufferCapacity(env, buf)
}

// jni.h:
//     jobjectRefType (JNICALL *GetObjectRefType)(JNIEnv* env, jobject obj);
func GetObjectRefType(env *C.JNIEnv, obj C.jobject) C.jobjectRefType {
	return C._GoJniGetObjectRefType(env, obj)
}
