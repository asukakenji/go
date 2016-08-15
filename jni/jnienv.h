#ifndef _GO_JNI_JNIENV_H_
#define _GO_JNI_JNIENV_H_

#include <jni.h>

// jni.h:
//     jint (JNICALL *GetVersion)(JNIEnv *env);
jint _GoJniGetVersion(JNIEnv* env);

// jni.h:
//     jclass (JNICALL *DefineClass)(JNIEnv *env, const char *name, jobject loader, const jbyte *buf, jsize len);
jclass _GoJniDefineClass(JNIEnv* env, const char* name, jobject loader, const jbyte* buf, jsize len);

// jni.h:
//     jclass (JNICALL *FindClass)(JNIEnv *env, const char *name);
jclass _GoJniFindClass(JNIEnv* env, const char* name);

// jni.h:
//     jmethodID (JNICALL *FromReflectedMethod)(JNIEnv *env, jobject method);
jmethodID _GoJniFromReflectedMethod(JNIEnv* env, jobject method);

// jni.h:
//     jfieldID (JNICALL *FromReflectedField)(JNIEnv *env, jobject field);
jfieldID _GoJniFromReflectedField(JNIEnv* env, jobject field);

// jni.h:
//     jobject (JNICALL *ToReflectedMethod)(JNIEnv *env, jclass cls, jmethodID methodID, jboolean isStatic);
jobject _GoJniToReflectedMethod(JNIEnv* env, jclass cls, jmethodID methodID, jboolean isStatic);

// jni.h:
//     jclass (JNICALL *GetSuperclass)(JNIEnv *env, jclass sub);
jclass _GoJniGetSuperclass(JNIEnv* env, jclass sub);

// jni.h:
//     jboolean (JNICALL *IsAssignableFrom)(JNIEnv *env, jclass sub, jclass sup);
jboolean _GoJniIsAssignableFrom(JNIEnv* env, jclass sub, jclass sup);

// jni.h:
//     jobject (JNICALL *ToReflectedField)(JNIEnv *env, jclass cls, jfieldID fieldID, jboolean isStatic);
jobject _GoJniToReflectedField(JNIEnv* env, jclass cls, jfieldID fieldID, jboolean isStatic);

// jni.h:
//     jint (JNICALL *Throw)(JNIEnv *env, jthrowable obj);
jint _GoJniThrow(JNIEnv* env, jthrowable obj);

// jni.h:
//     jint (JNICALL *ThrowNew)(JNIEnv *env, jclass clazz, const char *msg);
jint _GoJniThrowNew(JNIEnv* env, jclass clazz, const char* msg);

// jni.h:
//     jthrowable (JNICALL *ExceptionOccurred)(JNIEnv *env);
jthrowable _GoJniExceptionOccurred(JNIEnv* env);

// jni.h:
//     void (JNICALL *ExceptionDescribe)(JNIEnv *env);
void _GoJniExceptionDescribe(JNIEnv* env);

// jni.h:
//     void (JNICALL *ExceptionClear)(JNIEnv *env);
void _GoJniExceptionClear(JNIEnv* env);

// jni.h:
//     void (JNICALL *FatalError)(JNIEnv *env, const char *msg);
void _GoJniFatalError(JNIEnv* env, const char* msg);

// jni.h:
//     jint (JNICALL *PushLocalFrame)(JNIEnv *env, jint capacity);
jint _GoJniPushLocalFrame(JNIEnv* env, jint capacity);

// jni.h:
//     jobject (JNICALL *PopLocalFrame)(JNIEnv *env, jobject result);
jobject _GoJniPopLocalFrame(JNIEnv* env, jobject result);

// jni.h:
//     jobject (JNICALL *NewGlobalRef)(JNIEnv *env, jobject lobj);
jobject _GoJniNewGlobalRef(JNIEnv* env, jobject lobj);

// jni.h:
//     void (JNICALL *DeleteGlobalRef)(JNIEnv *env, jobject gref);
void _GoJniDeleteGlobalRef(JNIEnv* env, jobject gref);

// jni.h:
//     void (JNICALL *DeleteLocalRef)(JNIEnv *env, jobject obj);
void _GoJniDeleteLocalRef(JNIEnv* env, jobject obj);

// jni.h:
//     jboolean (JNICALL *IsSameObject)(JNIEnv *env, jobject obj1, jobject obj2);
jboolean _GoJniIsSameObject(JNIEnv* env, jobject obj1, jobject obj2);

// jni.h:
//     jobject (JNICALL *NewLocalRef)(JNIEnv *env, jobject ref);
jobject _GoJniNewLocalRef(JNIEnv* env, jobject ref);

// jni.h:
//     jint (JNICALL *EnsureLocalCapacity)(JNIEnv *env, jint capacity);
jint _GoJniEnsureLocalCapacity(JNIEnv* env, jint capacity);

// jni.h:
//     jobject (JNICALL *AllocObject)(JNIEnv *env, jclass clazz);
jobject _GoJniAllocObject(JNIEnv* env, jclass clazz);

// jni.h:
//     jobject (JNICALL *NewObject)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jobject (JNICALL *NewObjectV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jobject (JNICALL *NewObjectA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
jobject _GoJniNewObjectA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jclass (JNICALL *GetObjectClass)(JNIEnv *env, jobject obj);
jclass _GoJniGetObjectClass(JNIEnv* env, jobject obj);

// jni.h:
//     jboolean (JNICALL *IsInstanceOf)(JNIEnv *env, jobject obj, jclass clazz);
jboolean _GoJniIsInstanceOf(JNIEnv* env, jobject obj, jclass clazz);

// jni.h:
//     jmethodID (JNICALL *GetMethodID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);
jmethodID _GoJniGetMethodID(JNIEnv* env, jclass clazz, const char* name, const char* sig);

// jni.h:
//     jobject (JNICALL *CallObjectMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jobject (JNICALL *CallObjectMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jobject (JNICALL *CallObjectMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);
jobject _GoJniCallObjectMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args);

// jni.h:
//     jboolean (JNICALL *CallBooleanMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jboolean (JNICALL *CallBooleanMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jboolean (JNICALL *CallBooleanMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);
jboolean _GoJniCallBooleanMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args);

// jni.h:
//     jbyte (JNICALL *CallByteMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jbyte (JNICALL *CallByteMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jbyte (JNICALL *CallByteMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
jbyte _GoJniCallByteMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args);

// jni.h:
//     jchar (JNICALL *CallCharMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jchar (JNICALL *CallCharMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jchar (JNICALL *CallCharMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
jchar _GoJniCallCharMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args);

// jni.h:
//     jshort (JNICALL *CallShortMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jshort (JNICALL *CallShortMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jshort (JNICALL *CallShortMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
jshort _GoJniCallShortMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args);

// jni.h:
//     jint (JNICALL *CallIntMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jint (JNICALL *CallIntMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jint (JNICALL *CallIntMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
jint _GoJniCallIntMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args);

// jni.h:
//     jlong (JNICALL *CallLongMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jlong (JNICALL *CallLongMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jlong (JNICALL *CallLongMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
jlong _GoJniCallLongMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args);

// jni.h:
//     jfloat (JNICALL *CallFloatMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jfloat (JNICALL *CallFloatMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jfloat (JNICALL *CallFloatMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
jfloat _GoJniCallFloatMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args);

// jni.h:
//     jdouble (JNICALL *CallDoubleMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     jdouble (JNICALL *CallDoubleMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     jdouble (JNICALL *CallDoubleMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue *args);
jdouble _GoJniCallDoubleMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args);

// jni.h:
//     void (JNICALL *CallVoidMethod)(JNIEnv *env, jobject obj, jmethodID methodID, ...);
//     void (JNICALL *CallVoidMethodV)(JNIEnv *env, jobject obj, jmethodID methodID, va_list args);
//     void (JNICALL *CallVoidMethodA)(JNIEnv *env, jobject obj, jmethodID methodID, const jvalue * args);
void _GoJniCallVoidMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args);

// jni.h:
//     jobject (JNICALL *CallNonvirtualObjectMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jobject (JNICALL *CallNonvirtualObjectMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jobject (JNICALL *CallNonvirtualObjectMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);
jobject _GoJniCallNonvirtualObjectMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jboolean (JNICALL *CallNonvirtualBooleanMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jboolean (JNICALL *CallNonvirtualBooleanMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jboolean (JNICALL *CallNonvirtualBooleanMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);
jboolean _GoJniCallNonvirtualBooleanMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jbyte (JNICALL *CallNonvirtualByteMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jbyte (JNICALL *CallNonvirtualByteMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jbyte (JNICALL *CallNonvirtualByteMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
jbyte _GoJniCallNonvirtualByteMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jchar (JNICALL *CallNonvirtualCharMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jchar (JNICALL *CallNonvirtualCharMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jchar (JNICALL *CallNonvirtualCharMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
jchar _GoJniCallNonvirtualCharMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jshort (JNICALL *CallNonvirtualShortMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jshort (JNICALL *CallNonvirtualShortMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jshort (JNICALL *CallNonvirtualShortMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
jshort _GoJniCallNonvirtualShortMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jint (JNICALL *CallNonvirtualIntMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jint (JNICALL *CallNonvirtualIntMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jint (JNICALL *CallNonvirtualIntMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
jint _GoJniCallNonvirtualIntMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jlong (JNICALL *CallNonvirtualLongMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jlong (JNICALL *CallNonvirtualLongMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jlong (JNICALL *CallNonvirtualLongMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
jlong _GoJniCallNonvirtualLongMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jfloat (JNICALL *CallNonvirtualFloatMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jfloat (JNICALL *CallNonvirtualFloatMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jfloat (JNICALL *CallNonvirtualFloatMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
jfloat _GoJniCallNonvirtualFloatMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jdouble (JNICALL *CallNonvirtualDoubleMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     jdouble (JNICALL *CallNonvirtualDoubleMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     jdouble (JNICALL *CallNonvirtualDoubleMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue *args);
jdouble _GoJniCallNonvirtualDoubleMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     void (JNICALL *CallNonvirtualVoidMethod)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, ...);
//     void (JNICALL *CallNonvirtualVoidMethodV)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, va_list args);
//     void (JNICALL *CallNonvirtualVoidMethodA)(JNIEnv *env, jobject obj, jclass clazz, jmethodID methodID, const jvalue * args);
void _GoJniCallNonvirtualVoidMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jfieldID (JNICALL *GetFieldID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);
jfieldID _GoJniGetFieldID(JNIEnv* env, jclass clazz, const char* name, const char* sig);

// jni.h:
//     jobject (JNICALL *GetObjectField)(JNIEnv *env, jobject obj, jfieldID fieldID);
jobject _GoJniGetObjectField(JNIEnv* env, jobject obj, jfieldID fieldID);

// jni.h:
//     jboolean (JNICALL *GetBooleanField)(JNIEnv *env, jobject obj, jfieldID fieldID);
jboolean _GoJniGetBooleanField(JNIEnv* env, jobject obj, jfieldID fieldID);

// jni.h:
//     jbyte (JNICALL *GetByteField)(JNIEnv *env, jobject obj, jfieldID fieldID);
jbyte _GoJniGetByteField(JNIEnv* env, jobject obj, jfieldID fieldID);

// jni.h:
//     jchar (JNICALL *GetCharField)(JNIEnv *env, jobject obj, jfieldID fieldID);
jchar _GoJniGetCharField(JNIEnv* env, jobject obj, jfieldID fieldID);

// jni.h:
//     jshort (JNICALL *GetShortField)(JNIEnv *env, jobject obj, jfieldID fieldID);
jshort _GoJniGetShortField(JNIEnv* env, jobject obj, jfieldID fieldID);

// jni.h:
//     jint (JNICALL *GetIntField)(JNIEnv *env, jobject obj, jfieldID fieldID);
jint _GoJniGetIntField(JNIEnv* env, jobject obj, jfieldID fieldID);

// jni.h:
//     jlong (JNICALL *GetLongField)(JNIEnv *env, jobject obj, jfieldID fieldID);
jlong _GoJniGetLongField(JNIEnv* env, jobject obj, jfieldID fieldID);

// jni.h:
//     jfloat (JNICALL *GetFloatField)(JNIEnv *env, jobject obj, jfieldID fieldID);
jfloat _GoJniGetFloatField(JNIEnv* env, jobject obj, jfieldID fieldID);

// jni.h:
//     jdouble (JNICALL *GetDoubleField)(JNIEnv *env, jobject obj, jfieldID fieldID);
jdouble _GoJniGetDoubleField(JNIEnv* env, jobject obj, jfieldID fieldID);

// jni.h:
//     void (JNICALL *SetObjectField)(JNIEnv *env, jobject obj, jfieldID fieldID, jobject val);
void _GoJniSetObjectField(JNIEnv* env, jobject obj, jfieldID fieldID, jobject val);

// jni.h:
//     void (JNICALL *SetBooleanField)(JNIEnv *env, jobject obj, jfieldID fieldID, jboolean val);
void _GoJniSetBooleanField(JNIEnv* env, jobject obj, jfieldID fieldID, jboolean val);

// jni.h:
//     void (JNICALL *SetByteField)(JNIEnv *env, jobject obj, jfieldID fieldID, jbyte val);
void _GoJniSetByteField(JNIEnv* env, jobject obj, jfieldID fieldID, jbyte val);

// jni.h:
//     void (JNICALL *SetCharField)(JNIEnv *env, jobject obj, jfieldID fieldID, jchar val);
void _GoJniSetCharField(JNIEnv* env, jobject obj, jfieldID fieldID, jchar val);

// jni.h:
//     void (JNICALL *SetShortField)(JNIEnv *env, jobject obj, jfieldID fieldID, jshort val);
void _GoJniSetShortField(JNIEnv* env, jobject obj, jfieldID fieldID, jshort val);

// jni.h:
//     void (JNICALL *SetIntField)(JNIEnv *env, jobject obj, jfieldID fieldID, jint val);
void _GoJniSetIntField(JNIEnv* env, jobject obj, jfieldID fieldID, jint val);

// jni.h:
//     void (JNICALL *SetLongField)(JNIEnv *env, jobject obj, jfieldID fieldID, jlong val);
void _GoJniSetLongField(JNIEnv* env, jobject obj, jfieldID fieldID, jlong val);

// jni.h:
//     void (JNICALL *SetFloatField)(JNIEnv *env, jobject obj, jfieldID fieldID, jfloat val);
void _GoJniSetFloatField(JNIEnv* env, jobject obj, jfieldID fieldID, jfloat val);

// jni.h:
//     void (JNICALL *SetDoubleField)(JNIEnv *env, jobject obj, jfieldID fieldID, jdouble val);
void _GoJniSetDoubleField(JNIEnv* env, jobject obj, jfieldID fieldID, jdouble val);

// jni.h:
//     jmethodID (JNICALL *GetStaticMethodID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);
jmethodID _GoJniGetStaticMethodID(JNIEnv* env, jclass clazz, const char* name, const char* sig);

// jni.h:
//     jobject (JNICALL *CallStaticObjectMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jobject (JNICALL *CallStaticObjectMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jobject (JNICALL *CallStaticObjectMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
jobject _GoJniCallStaticObjectMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jboolean (JNICALL *CallStaticBooleanMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jboolean (JNICALL *CallStaticBooleanMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jboolean (JNICALL *CallStaticBooleanMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
jboolean _GoJniCallStaticBooleanMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jbyte (JNICALL *CallStaticByteMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jbyte (JNICALL *CallStaticByteMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jbyte (JNICALL *CallStaticByteMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
jbyte _GoJniCallStaticByteMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jchar (JNICALL *CallStaticCharMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jchar (JNICALL *CallStaticCharMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jchar (JNICALL *CallStaticCharMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
jchar _GoJniCallStaticCharMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jshort (JNICALL *CallStaticShortMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jshort (JNICALL *CallStaticShortMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jshort (JNICALL *CallStaticShortMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
jshort _GoJniCallStaticShortMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jint (JNICALL *CallStaticIntMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jint (JNICALL *CallStaticIntMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jint (JNICALL *CallStaticIntMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
jint _GoJniCallStaticIntMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jlong (JNICALL *CallStaticLongMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jlong (JNICALL *CallStaticLongMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jlong (JNICALL *CallStaticLongMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
jlong _GoJniCallStaticLongMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jfloat (JNICALL *CallStaticFloatMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jfloat (JNICALL *CallStaticFloatMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jfloat (JNICALL *CallStaticFloatMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
jfloat _GoJniCallStaticFloatMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     jdouble (JNICALL *CallStaticDoubleMethod)(JNIEnv *env, jclass clazz, jmethodID methodID, ...);
//     jdouble (JNICALL *CallStaticDoubleMethodV)(JNIEnv *env, jclass clazz, jmethodID methodID, va_list args);
//     jdouble (JNICALL *CallStaticDoubleMethodA)(JNIEnv *env, jclass clazz, jmethodID methodID, const jvalue *args);
jdouble _GoJniCallStaticDoubleMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args);

// jni.h:
//     void (JNICALL *CallStaticVoidMethod)(JNIEnv *env, jclass cls, jmethodID methodID, ...);
//     void (JNICALL *CallStaticVoidMethodV)(JNIEnv *env, jclass cls, jmethodID methodID, va_list args);
//     void (JNICALL *CallStaticVoidMethodA)(JNIEnv *env, jclass cls, jmethodID methodID, const jvalue * args);
void _GoJniCallStaticVoidMethodA(JNIEnv* env, jclass cls, jmethodID methodID, const jvalue* args);

// jni.h:
//     jfieldID (JNICALL *GetStaticFieldID)(JNIEnv *env, jclass clazz, const char *name, const char *sig);
jfieldID _GoJniGetStaticFieldID(JNIEnv* env, jclass clazz, const char* name, const char* sig);

// jni.h:
//     jobject (JNICALL *GetStaticObjectField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
jobject _GoJniGetStaticObjectField(JNIEnv* env, jclass clazz, jfieldID fieldID);

// jni.h:
//     jboolean (JNICALL *GetStaticBooleanField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
jboolean _GoJniGetStaticBooleanField(JNIEnv* env, jclass clazz, jfieldID fieldID);

// jni.h:
//     jbyte (JNICALL *GetStaticByteField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
jbyte _GoJniGetStaticByteField(JNIEnv* env, jclass clazz, jfieldID fieldID);

// jni.h:
//     jchar (JNICALL *GetStaticCharField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
jchar _GoJniGetStaticCharField(JNIEnv* env, jclass clazz, jfieldID fieldID);

// jni.h:
//     jshort (JNICALL *GetStaticShortField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
jshort _GoJniGetStaticShortField(JNIEnv* env, jclass clazz, jfieldID fieldID);

// jni.h:
//     jint (JNICALL *GetStaticIntField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
jint _GoJniGetStaticIntField(JNIEnv* env, jclass clazz, jfieldID fieldID);

// jni.h:
//     jlong (JNICALL *GetStaticLongField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
jlong _GoJniGetStaticLongField(JNIEnv* env, jclass clazz, jfieldID fieldID);

// jni.h:
//     jfloat (JNICALL *GetStaticFloatField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
jfloat _GoJniGetStaticFloatField(JNIEnv* env, jclass clazz, jfieldID fieldID);

// jni.h:
//     jdouble (JNICALL *GetStaticDoubleField)(JNIEnv *env, jclass clazz, jfieldID fieldID);
jdouble _GoJniGetStaticDoubleField(JNIEnv* env, jclass clazz, jfieldID fieldID);

// jni.h:
//     void (JNICALL *SetStaticObjectField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jobject value);
void _GoJniSetStaticObjectField(JNIEnv* env, jclass clazz, jfieldID fieldID, jobject value);

// jni.h:
//     void (JNICALL *SetStaticBooleanField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jboolean value);
void _GoJniSetStaticBooleanField(JNIEnv* env, jclass clazz, jfieldID fieldID, jboolean value);

// jni.h:
//     void (JNICALL *SetStaticByteField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jbyte value);
void _GoJniSetStaticByteField(JNIEnv* env, jclass clazz, jfieldID fieldID, jbyte value);

// jni.h:
//     void (JNICALL *SetStaticCharField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jchar value);
void _GoJniSetStaticCharField(JNIEnv* env, jclass clazz, jfieldID fieldID, jchar value);

// jni.h:
//     void (JNICALL *SetStaticShortField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jshort value);
void _GoJniSetStaticShortField(JNIEnv* env, jclass clazz, jfieldID fieldID, jshort value);

// jni.h:
//     void (JNICALL *SetStaticIntField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jint value);
void _GoJniSetStaticIntField(JNIEnv* env, jclass clazz, jfieldID fieldID, jint value);

// jni.h:
//     void (JNICALL *SetStaticLongField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jlong value);
void _GoJniSetStaticLongField(JNIEnv* env, jclass clazz, jfieldID fieldID, jlong value);

// jni.h:
//     void (JNICALL *SetStaticFloatField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jfloat value);
void _GoJniSetStaticFloatField(JNIEnv* env, jclass clazz, jfieldID fieldID, jfloat value);

// jni.h:
//     void (JNICALL *SetStaticDoubleField)(JNIEnv *env, jclass clazz, jfieldID fieldID, jdouble value);
void _GoJniSetStaticDoubleField(JNIEnv* env, jclass clazz, jfieldID fieldID, jdouble value);

// jni.h:
//     jstring (JNICALL *NewString)(JNIEnv *env, const jchar *unicode, jsize len);
jstring _GoJniNewString(JNIEnv* env, const jchar* unicode, jsize len);

// jni.h:
//     jsize (JNICALL *GetStringLength)(JNIEnv *env, jstring str);
jsize _GoJniGetStringLength(JNIEnv* env, jstring str);

// jni.h:
//     const jchar *(JNICALL *GetStringChars)(JNIEnv *env, jstring str, jboolean *isCopy);
const jchar* _GoJniGetStringChars(JNIEnv* env, jstring str, jboolean* isCopy);

// jni.h:
//     void (JNICALL *ReleaseStringChars)(JNIEnv *env, jstring str, const jchar *chars);
void _GoJniReleaseStringChars(JNIEnv* env, jstring str, const jchar* chars);

// jni.h:
//     jstring (JNICALL *NewStringUTF)(JNIEnv *env, const char *utf);
jstring _GoJniNewStringUTF(JNIEnv* env, const char* utf);

// jni.h:
//     jsize (JNICALL *GetStringUTFLength)(JNIEnv *env, jstring str);
jsize _GoJniGetStringUTFLength(JNIEnv* env, jstring str);

// jni.h:
//     const char* (JNICALL *GetStringUTFChars)(JNIEnv *env, jstring str, jboolean *isCopy);
const char* _GoJniGetStringUTFChars(JNIEnv* env, jstring str, jboolean* isCopy);

// jni.h:
//     void (JNICALL *ReleaseStringUTFChars)(JNIEnv *env, jstring str, const char* chars);
void _GoJniReleaseStringUTFChars(JNIEnv* env, jstring str, const char* chars);

// jni.h:
//     jsize (JNICALL *GetArrayLength)(JNIEnv *env, jarray array);
jsize _GoJniGetArrayLength(JNIEnv* env, jarray array);

// jni.h:
//     jobjectArray (JNICALL *NewObjectArray)(JNIEnv *env, jsize len, jclass clazz, jobject init);
jobjectArray _GoJniNewObjectArray(JNIEnv* env, jsize len, jclass clazz, jobject init);

// jni.h:
//     jobject (JNICALL *GetObjectArrayElement)(JNIEnv *env, jobjectArray array, jsize index);
jobject _GoJniGetObjectArrayElement(JNIEnv* env, jobjectArray array, jsize index);

// jni.h:
//     void (JNICALL *SetObjectArrayElement)(JNIEnv *env, jobjectArray array, jsize index, jobject val);
void _GoJniSetObjectArrayElement(JNIEnv* env, jobjectArray array, jsize index, jobject val);

// jni.h:
//     jbooleanArray (JNICALL *NewBooleanArray)(JNIEnv *env, jsize len);
jbooleanArray _GoJniNewBooleanArray(JNIEnv* env, jsize len);

// jni.h:
//     jbyteArray (JNICALL *NewByteArray)(JNIEnv *env, jsize len);
jbyteArray _GoJniNewByteArray(JNIEnv* env, jsize len);

// jni.h:
//     jcharArray (JNICALL *NewCharArray)(JNIEnv *env, jsize len);
jcharArray _GoJniNewCharArray(JNIEnv* env, jsize len);

// jni.h:
//     jshortArray (JNICALL *NewShortArray)(JNIEnv *env, jsize len);
jshortArray _GoJniNewShortArray(JNIEnv* env, jsize len);

// jni.h:
//     jintArray (JNICALL *NewIntArray)(JNIEnv *env, jsize len);
jintArray _GoJniNewIntArray(JNIEnv* env, jsize len);

// jni.h:
//     jlongArray (JNICALL *NewLongArray)(JNIEnv *env, jsize len);
jlongArray _GoJniNewLongArray(JNIEnv* env, jsize len);

// jni.h:
//     jfloatArray (JNICALL *NewFloatArray)(JNIEnv *env, jsize len);
jfloatArray _GoJniNewFloatArray(JNIEnv* env, jsize len);

// jni.h:
//     jdoubleArray (JNICALL *NewDoubleArray)(JNIEnv *env, jsize len);
jdoubleArray _GoJniNewDoubleArray(JNIEnv* env, jsize len);

// jni.h:
//     jboolean * (JNICALL *GetBooleanArrayElements)(JNIEnv *env, jbooleanArray array, jboolean *isCopy);
jboolean* _GoJniGetBooleanArrayElements(JNIEnv* env, jbooleanArray array, jboolean* isCopy);

// jni.h:
//     jbyte * (JNICALL *GetByteArrayElements)(JNIEnv *env, jbyteArray array, jboolean *isCopy);
jbyte* _GoJniGetByteArrayElements(JNIEnv* env, jbyteArray array, jboolean* isCopy);

// jni.h:
//     jchar * (JNICALL *GetCharArrayElements)(JNIEnv *env, jcharArray array, jboolean *isCopy);
jchar* _GoJniGetCharArrayElements(JNIEnv* env, jcharArray array, jboolean* isCopy);

// jni.h:
//     jshort * (JNICALL *GetShortArrayElements)(JNIEnv *env, jshortArray array, jboolean *isCopy);
jshort* _GoJniGetShortArrayElements(JNIEnv* env, jshortArray array, jboolean* isCopy);

// jni.h:
//     jint * (JNICALL *GetIntArrayElements)(JNIEnv *env, jintArray array, jboolean *isCopy);
jint* _GoJniGetIntArrayElements(JNIEnv* env, jintArray array, jboolean* isCopy);

// jni.h:
//     jlong * (JNICALL *GetLongArrayElements)(JNIEnv *env, jlongArray array, jboolean *isCopy);
jlong* _GoJniGetLongArrayElements(JNIEnv* env, jlongArray array, jboolean* isCopy);

// jni.h:
//     jfloat * (JNICALL *GetFloatArrayElements)(JNIEnv *env, jfloatArray array, jboolean *isCopy);
jfloat* _GoJniGetFloatArrayElements(JNIEnv* env, jfloatArray array, jboolean* isCopy);

// jni.h:
//     jdouble * (JNICALL *GetDoubleArrayElements)(JNIEnv *env, jdoubleArray array, jboolean *isCopy);
jdouble* _GoJniGetDoubleArrayElements(JNIEnv* env, jdoubleArray array, jboolean* isCopy);

// jni.h:
//     void (JNICALL *ReleaseBooleanArrayElements)(JNIEnv *env, jbooleanArray array, jboolean *elems, jint mode);
void _GoJniReleaseBooleanArrayElements(JNIEnv* env, jbooleanArray array, jboolean* elems, jint mode);

// jni.h:
//     void (JNICALL *ReleaseByteArrayElements)(JNIEnv *env, jbyteArray array, jbyte *elems, jint mode);
void _GoJniReleaseByteArrayElements(JNIEnv* env, jbyteArray array, jbyte* elems, jint mode);

// jni.h:
//     void (JNICALL *ReleaseCharArrayElements)(JNIEnv *env, jcharArray array, jchar *elems, jint mode);
void _GoJniReleaseCharArrayElements(JNIEnv* env, jcharArray array, jchar* elems, jint mode);

// jni.h:
//     void (JNICALL *ReleaseShortArrayElements)(JNIEnv *env, jshortArray array, jshort *elems, jint mode);
void _GoJniReleaseShortArrayElements(JNIEnv* env, jshortArray array, jshort* elems, jint mode);

// jni.h:
//     void (JNICALL *ReleaseIntArrayElements)(JNIEnv *env, jintArray array, jint *elems, jint mode);
void _GoJniReleaseIntArrayElements(JNIEnv* env, jintArray array, jint* elems, jint mode);

// jni.h:
//     void (JNICALL *ReleaseLongArrayElements)(JNIEnv *env, jlongArray array, jlong *elems, jint mode);
void _GoJniReleaseLongArrayElements(JNIEnv* env, jlongArray array, jlong* elems, jint mode);

// jni.h:
//     void (JNICALL *ReleaseFloatArrayElements)(JNIEnv *env, jfloatArray array, jfloat *elems, jint mode);
void _GoJniReleaseFloatArrayElements(JNIEnv* env, jfloatArray array, jfloat* elems, jint mode);

// jni.h:
//     void (JNICALL *ReleaseDoubleArrayElements)(JNIEnv *env, jdoubleArray array, jdouble *elems, jint mode);
void _GoJniReleaseDoubleArrayElements(JNIEnv* env, jdoubleArray array, jdouble* elems, jint mode);

// jni.h:
//     void (JNICALL *GetBooleanArrayRegion)(JNIEnv *env, jbooleanArray array, jsize start, jsize l, jboolean *buf);
void _GoJniGetBooleanArrayRegion(JNIEnv* env, jbooleanArray array, jsize start, jsize l, jboolean* buf);

// jni.h:
//     void (JNICALL *GetByteArrayRegion)(JNIEnv *env, jbyteArray array, jsize start, jsize len, jbyte *buf);
void _GoJniGetByteArrayRegion(JNIEnv* env, jbyteArray array, jsize start, jsize len, jbyte* buf);

// jni.h:
//     void (JNICALL *GetCharArrayRegion)(JNIEnv *env, jcharArray array, jsize start, jsize len, jchar *buf);
void _GoJniGetCharArrayRegion(JNIEnv* env, jcharArray array, jsize start, jsize len, jchar* buf);

// jni.h:
//     void (JNICALL *GetShortArrayRegion)(JNIEnv *env, jshortArray array, jsize start, jsize len, jshort *buf);
void _GoJniGetShortArrayRegion(JNIEnv* env, jshortArray array, jsize start, jsize len, jshort* buf);

// jni.h:
//     void (JNICALL *GetIntArrayRegion)(JNIEnv *env, jintArray array, jsize start, jsize len, jint *buf);
void _GoJniGetIntArrayRegion(JNIEnv* env, jintArray array, jsize start, jsize len, jint* buf);

// jni.h:
//     void (JNICALL *GetLongArrayRegion)(JNIEnv *env, jlongArray array, jsize start, jsize len, jlong *buf);
void _GoJniGetLongArrayRegion(JNIEnv* env, jlongArray array, jsize start, jsize len, jlong* buf);

// jni.h:
//     void (JNICALL *GetFloatArrayRegion)(JNIEnv *env, jfloatArray array, jsize start, jsize len, jfloat *buf);
void _GoJniGetFloatArrayRegion(JNIEnv* env, jfloatArray array, jsize start, jsize len, jfloat* buf);

// jni.h:
//     void (JNICALL *GetDoubleArrayRegion)(JNIEnv *env, jdoubleArray array, jsize start, jsize len, jdouble *buf);
void _GoJniGetDoubleArrayRegion(JNIEnv* env, jdoubleArray array, jsize start, jsize len, jdouble* buf);

// jni.h:
//     void (JNICALL *SetBooleanArrayRegion)(JNIEnv *env, jbooleanArray array, jsize start, jsize l, const jboolean *buf);
void _GoJniSetBooleanArrayRegion(JNIEnv* env, jbooleanArray array, jsize start, jsize l, const jboolean* buf);

// jni.h:
//     void (JNICALL *SetByteArrayRegion)(JNIEnv *env, jbyteArray array, jsize start, jsize len, const jbyte *buf);
void _GoJniSetByteArrayRegion(JNIEnv* env, jbyteArray array, jsize start, jsize len, const jbyte* buf);

// jni.h:
//     void (JNICALL *SetCharArrayRegion)(JNIEnv *env, jcharArray array, jsize start, jsize len, const jchar *buf);
void _GoJniSetCharArrayRegion(JNIEnv* env, jcharArray array, jsize start, jsize len, const jchar* buf);

// jni.h:
//     void (JNICALL *SetShortArrayRegion)(JNIEnv *env, jshortArray array, jsize start, jsize len, const jshort *buf);
void _GoJniSetShortArrayRegion(JNIEnv* env, jshortArray array, jsize start, jsize len, const jshort* buf);

// jni.h:
//     void (JNICALL *SetIntArrayRegion)(JNIEnv *env, jintArray array, jsize start, jsize len, const jint *buf);
void _GoJniSetIntArrayRegion(JNIEnv* env, jintArray array, jsize start, jsize len, const jint* buf);

// jni.h:
//     void (JNICALL *SetLongArrayRegion)(JNIEnv *env, jlongArray array, jsize start, jsize len, const jlong *buf);
void _GoJniSetLongArrayRegion(JNIEnv* env, jlongArray array, jsize start, jsize len, const jlong* buf);

// jni.h:
//     void (JNICALL *SetFloatArrayRegion)(JNIEnv *env, jfloatArray array, jsize start, jsize len, const jfloat *buf);
void _GoJniSetFloatArrayRegion(JNIEnv* env, jfloatArray array, jsize start, jsize len, const jfloat* buf);

// jni.h:
//     void (JNICALL *SetDoubleArrayRegion)(JNIEnv *env, jdoubleArray array, jsize start, jsize len, const jdouble *buf);
void _GoJniSetDoubleArrayRegion(JNIEnv* env, jdoubleArray array, jsize start, jsize len, const jdouble* buf);

// jni.h:
//     jint (JNICALL *RegisterNatives)(JNIEnv *env, jclass clazz, const JNINativeMethod *methods, jint nMethods);
jint _GoJniRegisterNatives(JNIEnv* env, jclass clazz, const JNINativeMethod* methods, jint nMethods);

// jni.h:
//     jint (JNICALL *UnregisterNatives)(JNIEnv *env, jclass clazz);
jint _GoJniUnregisterNatives(JNIEnv* env, jclass clazz);

// jni.h:
//     jint (JNICALL *MonitorEnter)(JNIEnv *env, jobject obj);
jint _GoJniMonitorEnter(JNIEnv* env, jobject obj);

// jni.h:
//     jint (JNICALL *MonitorExit)(JNIEnv *env, jobject obj);
jint _GoJniMonitorExit(JNIEnv* env, jobject obj);

// jni.h:
//     jint (JNICALL *GetJavaVM)(JNIEnv *env, JavaVM **vm);
jint _GoJniGetJavaVM(JNIEnv* env, JavaVM** vm);

// jni.h:
//     void (JNICALL *GetStringRegion)(JNIEnv *env, jstring str, jsize start, jsize len, jchar *buf);
void _GoJniGetStringRegion(JNIEnv* env, jstring str, jsize start, jsize len, jchar* buf);

// jni.h:
//     void (JNICALL *GetStringUTFRegion)(JNIEnv *env, jstring str, jsize start, jsize len, char *buf);
void _GoJniGetStringUTFRegion(JNIEnv* env, jstring str, jsize start, jsize len, char* buf);

// jni.h:
//     void * (JNICALL *GetPrimitiveArrayCritical)(JNIEnv *env, jarray array, jboolean *isCopy);
void* _GoJniGetPrimitiveArrayCritical(JNIEnv* env, jarray array, jboolean* isCopy);

// jni.h:
//     void (JNICALL *ReleasePrimitiveArrayCritical)(JNIEnv *env, jarray array, void *carray, jint mode);
void _GoJniReleasePrimitiveArrayCritical(JNIEnv* env, jarray array, void* carray, jint mode);

// jni.h:
//     const jchar * (JNICALL *GetStringCritical)(JNIEnv *env, jstring string, jboolean *isCopy);
const jchar* _GoJniGetStringCritical(JNIEnv* env, jstring string, jboolean* isCopy);

// jni.h:
//     void (JNICALL *ReleaseStringCritical)(JNIEnv *env, jstring string, const jchar *cstring);
void _GoJniReleaseStringCritical(JNIEnv* env, jstring string, const jchar* cstring);

// jni.h:
//     jweak (JNICALL *NewWeakGlobalRef)(JNIEnv *env, jobject obj);
jweak _GoJniNewWeakGlobalRef(JNIEnv* env, jobject obj);

// jni.h:
//     void (JNICALL *DeleteWeakGlobalRef)(JNIEnv *env, jweak ref);
void _GoJniDeleteWeakGlobalRef(JNIEnv* env, jweak ref);

// jni.h:
//     jboolean (JNICALL *ExceptionCheck)(JNIEnv *env);
jboolean _GoJniExceptionCheck(JNIEnv* env);

// jni.h:
//     jobject (JNICALL *NewDirectByteBuffer)(JNIEnv* env, void* address, jlong capacity);
jobject _GoJniNewDirectByteBuffer(JNIEnv* env, void* address, jlong capacity);

// jni.h:
//     void* (JNICALL *GetDirectBufferAddress)(JNIEnv* env, jobject buf);
void* _GoJniGetDirectBufferAddress(JNIEnv* env, jobject buf);

// jni.h:
//     jlong (JNICALL *GetDirectBufferCapacity)(JNIEnv* env, jobject buf);
jlong _GoJniGetDirectBufferCapacity(JNIEnv* env, jobject buf);

// jni.h:
//     jobjectRefType (JNICALL *GetObjectRefType)(JNIEnv* env, jobject obj);
jobjectRefType _GoJniGetObjectRefType(JNIEnv* env, jobject obj);

#endif // #ifndef _GO_JNI_JNIENV_H_
