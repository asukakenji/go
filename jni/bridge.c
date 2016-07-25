#include "bridge.h"

jmethodID GoGetStaticMethodID(JNIEnv *env, jclass class, const char *name, const char *sig)
{
	return (*env)->GetStaticMethodID(env, class, name, sig);
}

void GoCallStaticVoidMethod(JNIEnv *env, jclass clazz, jmethodID methodID)
{
	return (*env)->CallStaticVoidMethod(env, clazz, methodID);
}
