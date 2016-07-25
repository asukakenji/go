#ifndef _BRIDGE_H_
#define _BRIDGE_H_

// JNIEnv, jclass, jmethodID
#include <jni.h>

jmethodID GoGetStaticMethodID(JNIEnv *env, jclass class, const char *name, const char *sig);

void GoCallStaticVoidMethod(JNIEnv *env, jclass clazz, jmethodID methodID);

#endif // #ifndef _BRIDGE_H_
