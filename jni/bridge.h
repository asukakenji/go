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

// jni.h:
//     jint (JNICALL *Throw)(JNIEnv *env, jthrowable obj);
jint _GoJniThrow(JNIEnv *env, jthrowable obj);

// jni.h:
//     jint (JNICALL *ThrowNew)(JNIEnv *env, jclass clazz, const char *msg);
jint _GoJniThrowNew(JNIEnv *env, jclass clazz, const char *msg);

// jni.h:
//     jthrowable (JNICALL *ExceptionOccurred)(JNIEnv *env);
jthrowable _GoJniExceptionOccurred(JNIEnv *env);

// jni.h:
//     void (JNICALL *ExceptionDescribe)(JNIEnv *env);
void _GoJniExceptionDescribe(JNIEnv *env);

// jni.h:
//     void (JNICALL *ExceptionClear)(JNIEnv *env);
void _GoJniExceptionClear(JNIEnv *env);

// jni.h:
//     void (JNICALL *FatalError)(JNIEnv *env, const char *msg);
void _GoJniFatalError(JNIEnv *env, const char *msg);

// jni.h:
//     jint (JNICALL *PushLocalFrame)(JNIEnv *env, jint capacity);
jint _GoJniPushLocalFrame(JNIEnv *env, jint capacity);

// jni.h:
//     jobject (JNICALL *PopLocalFrame)(JNIEnv *env, jobject result);
jobject _GoJniPopLocalFrame(JNIEnv *env, jobject result);

// jni.h:
//     jobject (JNICALL *NewGlobalRef)(JNIEnv *env, jobject lobj);
jobject _GoJniNewGlobalRef(JNIEnv *env, jobject lobj);

// jni.h:
//     void (JNICALL *DeleteGlobalRef)(JNIEnv *env, jobject gref);
void _GoJniDeleteGlobalRef(JNIEnv *env, jobject gref);

// jni.h:
//     void (JNICALL *DeleteLocalRef)(JNIEnv *env, jobject obj);
void _GoJniDeleteLocalRef(JNIEnv *env, jobject obj);

// jni.h:
//     jboolean (JNICALL *IsSameObject)(JNIEnv *env, jobject obj1, jobject obj2);
jboolean _GoJniIsSameObject(JNIEnv *env, jobject obj1, jobject obj2);

// jni.h:
//     jobject (JNICALL *NewLocalRef)(JNIEnv *env, jobject ref);
jobject _GoJniNewLocalRef(JNIEnv *env, jobject ref);

// jni.h:
//     jint (JNICALL *EnsureLocalCapacity)(JNIEnv *env, jint capacity);
jint _GoJniEnsureLocalCapacity(JNIEnv *env, jint capacity);

// jni.h:
//     jobject (JNICALL *AllocObject)(JNIEnv *env, jclass clazz);
jobject _GoJniAllocObject(JNIEnv *env, jclass clazz);

// jni.h:
//     jobject (JNICALL *NewObject)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jobject (JNICALL *NewObjectV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jobject (JNICALL *NewObjectA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
jobject _GoJniNewObjectA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);

// jni.h:
//     jclass (JNICALL *GetObjectClass)(JNIEnv *env, jobject obj);
jclass _GoJniGetObjectClass(JNIEnv *env, jobject obj);

// jni.h:
//     jboolean (JNICALL *IsInstanceOf)(JNIEnv *env, jobject obj, jclass clazz);
jboolean _GoJniIsInstanceOf(JNIEnv *env, jobject obj, jclass clazz);




jmethodID _GoJniGetStaticMethodID(JNIEnv *env, jclass class, const char *name, const char *sig);

void _GoJniCallStaticVoidMethod(JNIEnv *env, jclass clazz, jmethodID methodID);

#endif // #ifndef _GO_JNI_BRIDGE_H_
