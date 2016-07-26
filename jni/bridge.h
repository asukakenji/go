#ifndef _GO_JNI_BRIDGE_H_
#define _GO_JNI_BRIDGE_H_

#include <jni.h>

// jni.h:
//     jint (JNICALL *GetVersion)(JNIEnv *env);
jint _GoJniGetVersion(JNIEnv *env);

// jni.h:
//     jclass (JNICALL *DefineClass)(JNIEnv *env, const char *name, jobject loader, const jbyte *buf, jsize len);
jclass _GoJniDefineClass(JNIEnv *env, const char *name, jobject loader, const jbyte *buf, jsize len);

// jni.h:
//     jclass (JNICALL *FindClass)(JNIEnv *env, const char *name);
jclass _GoJniFindClass(JNIEnv *env, const char *name);

// jni.h:
//     jmethodID (JNICALL *FromReflectedMethod)(JNIEnv *env, jobject method);
jmethodID _GoJniFromReflectedMethod(JNIEnv *env, jobject method);

// jni.h:
//     jfieldID (JNICALL *FromReflectedField)(JNIEnv *env, jobject field);
jfieldID _GoJniFromReflectedField(JNIEnv *env, jobject field);

// jni.h:
//     jobject (JNICALL *ToReflectedMethod)(JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic);
jobject _GoJniToReflectedMethod(JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic);

// jni.h:
//     jclass (JNICALL *GetSuperclass)(JNIEnv *env, jclass sub);
jclass _GoJniGetSuperclass(JNIEnv *env, jclass sub);

// jni.h:
//     jboolean (JNICALL *IsAssignableFrom)(JNIEnv *env, jclass sub, jclass sup);
jboolean _GoJniIsAssignableFrom(JNIEnv *env, jclass sub, jclass sup);

// jni.h:
//     jobject (JNICALL *ToReflectedField)(JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic);
jobject _GoJniToReflectedField(JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic);


jmethodID _GoJniGetStaticMethodID(JNIEnv *env, jclass class, const char *name, const char *sig);

void _GoJniCallStaticVoidMethod(JNIEnv *env, jclass clazz, jmethodID methodID);

#endif // #ifndef _GO_JNI_BRIDGE_H_
