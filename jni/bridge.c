#include "bridge.h"

jint _GoJniGetVersion(JNIEnv *env)
{
	return (*env)->GetVersion(env);
}

jclass _GoJniDefineClass(JNIEnv *env, const char *name, jobject loader, const jbyte *buf, jsize len)
{
	return (*env)->DefineClass(env, name, loader, buf, len);
}

jclass _GoJniFindClass(JNIEnv *env, const char *name)
{
	return (*env)->FindClass(env, name);
}

jmethodID _GoJniFromReflectedMethod(JNIEnv *env, jobject method)
{
	return (*env)->FromReflectedMethod(env, method);
}

jfieldID _GoJniFromReflectedField(JNIEnv *env, jobject field)
{
	return (*env)->FromReflectedField(env, field);
}

jobject _GoJniToReflectedMethod(JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic)
{
	return (*env)->ToReflectedMethod(env, cls, methodID, isStatic);
}

jclass _GoJniGetSuperclass(JNIEnv *env, jclass sub)
{
	return (*env)->GetSuperclass(env, sub);
}

jboolean _GoJniIsAssignableFrom(JNIEnv *env, jclass sub, jclass sup)
{
	return (*env)->IsAssignableFrom(env, sub, sup);
}

jobject _GoJniToReflectedField(JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic)
{
	return (*env)->ToReflectedField(env, cls, fieldID, isStatic);
}

jmethodID _GoJniGetStaticMethodID(JNIEnv *env, jclass class, const char *name, const char *sig)
{
	return (*env)->GetStaticMethodID(env, class, name, sig);
}

void _GoJniCallStaticVoidMethod(JNIEnv *env, jclass clazz, jmethodID methodID)
{
	return (*env)->CallStaticVoidMethod(env, clazz, methodID);
}
