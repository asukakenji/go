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

jint _GoJniThrow(JNIEnv *env, jthrowable obj)
{
	return (*env)->Throw(env, obj);
}

jint _GoJniThrowNew(JNIEnv *env, jclass clazz, const char *msg)
{
	return (*env)->ThrowNew(env, clazz, msg);
}

jthrowable _GoJniExceptionOccurred(JNIEnv *env)
{
	return (*env)->ExceptionOccurred(env);
}

void _GoJniExceptionDescribe(JNIEnv *env)
{
	return (*env)->ExceptionDescribe(env);
}

void _GoJniExceptionClear(JNIEnv *env)
{
	return (*env)->ExceptionClear(env);
}

void _GoJniFatalError(JNIEnv *env, const char *msg)
{
	return (*env)->FatalError(env, msg);
}

jint _GoJniPushLocalFrame(JNIEnv *env, jint capacity)
{
	return (*env)->PushLocalFrame(env, capacity);
}

jobject _GoJniPopLocalFrame(JNIEnv *env, jobject result)
{
	return (*env)->PopLocalFrame(env, result);
}

jobject _GoJniNewGlobalRef(JNIEnv *env, jobject lobj)
{
	return (*env)->NewGlobalRef(env, lobj);
}

void _GoJniDeleteGlobalRef(JNIEnv *env, jobject gref)
{
	return (*env)->DeleteGlobalRef(env, gref);
}

void _GoJniDeleteLocalRef(JNIEnv *env, jobject obj)
{
	return (*env)->DeleteLocalRef(env, obj);
}

jboolean _GoJniIsSameObject(JNIEnv *env, jobject obj1, jobject obj2)
{
	return (*env)->IsSameObject(env, obj1, obj2);
}

jobject _GoJniNewLocalRef(JNIEnv *env, jobject ref)
{
	return (*env)->NewLocalRef(env, ref);
}

jint _GoJniEnsureLocalCapacity(JNIEnv *env, jint capacity)
{
	return (*env)->EnsureLocalCapacity(env, capacity);
}

jobject _GoJniAllocObject(JNIEnv *env, jclass clazz)
{
	return (*env)->AllocObject(env, clazz);
}

jobject _GoJniNewObjectA(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args)
{
	return (*env)->NewObjectA(env, clazz, methodID, args);
}

jclass _GoJniGetObjectClass(JNIEnv *env, jobject obj)
{
	return (*env)->GetObjectClass(env, obj);
}

jboolean _GoJniIsInstanceOf(JNIEnv *env, jobject obj, jclass clazz)
{
	return (*env)->IsInstanceOf(env, obj, clazz);
}





jmethodID _GoJniGetStaticMethodID(JNIEnv *env, jclass class, const char *name, const char *sig)
{
	return (*env)->GetStaticMethodID(env, class, name, sig);
}

void _GoJniCallStaticVoidMethod(JNIEnv *env, jclass clazz, jmethodID methodID)
{
	return (*env)->CallStaticVoidMethod(env, clazz, methodID);
}
