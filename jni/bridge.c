#include "jnienv.h"

jint _GoJniGetVersion(JNIEnv* env)
{
	return (*env)->GetVersion(env);
}

jclass _GoJniDefineClass(JNIEnv* env, const char* name, jobject loader, const jbyte* buf, jsize len)
{
	return (*env)->DefineClass(env, name, loader, buf, len);
}

jclass _GoJniFindClass(JNIEnv* env, const char* name)
{
	return (*env)->FindClass(env, name);
}

jmethodID _GoJniFromReflectedMethod(JNIEnv* env, jobject method)
{
	return (*env)->FromReflectedMethod(env, method);
}

jfieldID _GoJniFromReflectedField(JNIEnv* env, jobject field)
{
	return (*env)->FromReflectedField(env, field);
}

jobject _GoJniToReflectedMethod(JNIEnv* env, jclass cls, jmethodID methodID, jboolean isStatic)
{
	return (*env)->ToReflectedMethod(env, cls, methodID, isStatic);
}

jclass _GoJniGetSuperclass(JNIEnv* env, jclass sub)
{
	return (*env)->GetSuperclass(env, sub);
}

jboolean _GoJniIsAssignableFrom(JNIEnv* env, jclass sub, jclass sup)
{
	return (*env)->IsAssignableFrom(env, sub, sup);
}

jobject _GoJniToReflectedField(JNIEnv* env, jclass cls, jfieldID fieldID, jboolean isStatic)
{
	return (*env)->ToReflectedField(env, cls, fieldID, isStatic);
}

jint _GoJniThrow(JNIEnv* env, jthrowable obj)
{
	return (*env)->Throw(env, obj);
}

jint _GoJniThrowNew(JNIEnv* env, jclass clazz, const char* msg)
{
	return (*env)->ThrowNew(env, clazz, msg);
}

jthrowable _GoJniExceptionOccurred(JNIEnv* env)
{
	return (*env)->ExceptionOccurred(env);
}

void _GoJniExceptionDescribe(JNIEnv* env)
{
	return (*env)->ExceptionDescribe(env);
}

void _GoJniExceptionClear(JNIEnv* env)
{
	return (*env)->ExceptionClear(env);
}

void _GoJniFatalError(JNIEnv* env, const char* msg)
{
	return (*env)->FatalError(env, msg);
}

jint _GoJniPushLocalFrame(JNIEnv* env, jint capacity)
{
	return (*env)->PushLocalFrame(env, capacity);
}

jobject _GoJniPopLocalFrame(JNIEnv* env, jobject result)
{
	return (*env)->PopLocalFrame(env, result);
}

jobject _GoJniNewGlobalRef(JNIEnv* env, jobject lobj)
{
	return (*env)->NewGlobalRef(env, lobj);
}

void _GoJniDeleteGlobalRef(JNIEnv* env, jobject gref)
{
	return (*env)->DeleteGlobalRef(env, gref);
}

void _GoJniDeleteLocalRef(JNIEnv* env, jobject obj)
{
	return (*env)->DeleteLocalRef(env, obj);
}

jboolean _GoJniIsSameObject(JNIEnv* env, jobject obj1, jobject obj2)
{
	return (*env)->IsSameObject(env, obj1, obj2);
}

jobject _GoJniNewLocalRef(JNIEnv* env, jobject ref)
{
	return (*env)->NewLocalRef(env, ref);
}

jint _GoJniEnsureLocalCapacity(JNIEnv* env, jint capacity)
{
	return (*env)->EnsureLocalCapacity(env, capacity);
}

jobject _GoJniAllocObject(JNIEnv* env, jclass clazz)
{
	return (*env)->AllocObject(env, clazz);
}

jobject _GoJniNewObjectA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->NewObjectA(env, clazz, methodID, args);
}

jclass _GoJniGetObjectClass(JNIEnv* env, jobject obj)
{
	return (*env)->GetObjectClass(env, obj);
}

jboolean _GoJniIsInstanceOf(JNIEnv* env, jobject obj, jclass clazz)
{
	return (*env)->IsInstanceOf(env, obj, clazz);
}

jmethodID _GoJniGetMethodID(JNIEnv* env, jclass clazz, const char* name, const char* sig)
{
	return (*env)->GetMethodID(env, clazz, name, sig);
}

jobject _GoJniCallObjectMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallObjectMethodA(env, obj, methodID, args);
}

jboolean _GoJniCallBooleanMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallBooleanMethodA(env, obj, methodID, args);
}

jbyte _GoJniCallByteMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallByteMethodA(env, obj, methodID, args);
}

jchar _GoJniCallCharMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallCharMethodA(env, obj, methodID, args);
}

jshort _GoJniCallShortMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallShortMethodA(env, obj, methodID, args);
}

jint _GoJniCallIntMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallIntMethodA(env, obj, methodID, args);
}

jlong _GoJniCallLongMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallLongMethodA(env, obj, methodID, args);
}

jfloat _GoJniCallFloatMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallFloatMethodA(env, obj, methodID, args);
}

jdouble _GoJniCallDoubleMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallDoubleMethodA(env, obj, methodID, args);
}

void _GoJniCallVoidMethodA(JNIEnv* env, jobject obj, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallVoidMethodA(env, obj, methodID, args);
}

jobject _GoJniCallNonvirtualObjectMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallNonvirtualObjectMethodA(env, obj, clazz, methodID, args);
}

jboolean _GoJniCallNonvirtualBooleanMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallNonvirtualBooleanMethodA(env, obj, clazz, methodID, args);
}

jbyte _GoJniCallNonvirtualByteMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallNonvirtualByteMethodA(env, obj, clazz, methodID, args);
}

jchar _GoJniCallNonvirtualCharMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallNonvirtualCharMethodA(env, obj, clazz, methodID, args);
}

jshort _GoJniCallNonvirtualShortMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallNonvirtualShortMethodA(env, obj, clazz, methodID, args);
}

jint _GoJniCallNonvirtualIntMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallNonvirtualIntMethodA(env, obj, clazz, methodID, args);
}

jlong _GoJniCallNonvirtualLongMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallNonvirtualLongMethodA(env, obj, clazz, methodID, args);
}

jfloat _GoJniCallNonvirtualFloatMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallNonvirtualFloatMethodA(env, obj, clazz, methodID, args);
}

jdouble _GoJniCallNonvirtualDoubleMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallNonvirtualDoubleMethodA(env, obj, clazz, methodID, args);
}

void _GoJniCallNonvirtualVoidMethodA(JNIEnv* env, jobject obj, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallNonvirtualVoidMethodA(env, obj, clazz, methodID, args);
}

jfieldID _GoJniGetFieldID(JNIEnv* env, jclass clazz, const char* name, const char* sig)
{
	return (*env)->GetFieldID(env, clazz, name, sig);
}

jobject _GoJniGetObjectField(JNIEnv* env, jobject obj, jfieldID fieldID)
{
	return (*env)->GetObjectField(env, obj, fieldID);
}

jboolean _GoJniGetBooleanField(JNIEnv* env, jobject obj, jfieldID fieldID)
{
	return (*env)->GetBooleanField(env, obj, fieldID);
}

jbyte _GoJniGetByteField(JNIEnv* env, jobject obj, jfieldID fieldID)
{
	return (*env)->GetByteField(env, obj, fieldID);
}

jchar _GoJniGetCharField(JNIEnv* env, jobject obj, jfieldID fieldID)
{
	return (*env)->GetCharField(env, obj, fieldID);
}

jshort _GoJniGetShortField(JNIEnv* env, jobject obj, jfieldID fieldID)
{
	return (*env)->GetShortField(env, obj, fieldID);
}

jint _GoJniGetIntField(JNIEnv* env, jobject obj, jfieldID fieldID)
{
	return (*env)->GetIntField(env, obj, fieldID);
}

jlong _GoJniGetLongField(JNIEnv* env, jobject obj, jfieldID fieldID)
{
	return (*env)->GetLongField(env, obj, fieldID);
}

jfloat _GoJniGetFloatField(JNIEnv* env, jobject obj, jfieldID fieldID)
{
	return (*env)->GetFloatField(env, obj, fieldID);
}

jdouble _GoJniGetDoubleField(JNIEnv* env, jobject obj, jfieldID fieldID)
{
	return (*env)->GetDoubleField(env, obj, fieldID);
}

void _GoJniSetObjectField(JNIEnv* env, jobject obj, jfieldID fieldID, jobject val)
{
	return (*env)->SetObjectField(env, obj, fieldID, val);
}

void _GoJniSetBooleanField(JNIEnv* env, jobject obj, jfieldID fieldID, jboolean val)
{
	return (*env)->SetBooleanField(env, obj, fieldID, val);
}

void _GoJniSetByteField(JNIEnv* env, jobject obj, jfieldID fieldID, jbyte val)
{
	return (*env)->SetByteField(env, obj, fieldID, val);
}

void _GoJniSetCharField(JNIEnv* env, jobject obj, jfieldID fieldID, jchar val)
{
	return (*env)->SetCharField(env, obj, fieldID, val);
}

void _GoJniSetShortField(JNIEnv* env, jobject obj, jfieldID fieldID, jshort val)
{
	return (*env)->SetShortField(env, obj, fieldID, val);
}

void _GoJniSetIntField(JNIEnv* env, jobject obj, jfieldID fieldID, jint val)
{
	return (*env)->SetIntField(env, obj, fieldID, val);
}

void _GoJniSetLongField(JNIEnv* env, jobject obj, jfieldID fieldID, jlong val)
{
	return (*env)->SetLongField(env, obj, fieldID, val);
}

void _GoJniSetFloatField(JNIEnv* env, jobject obj, jfieldID fieldID, jfloat val)
{
	return (*env)->SetFloatField(env, obj, fieldID, val);
}

void _GoJniSetDoubleField(JNIEnv* env, jobject obj, jfieldID fieldID, jdouble val)
{
	return (*env)->SetDoubleField(env, obj, fieldID, val);
}

jmethodID _GoJniGetStaticMethodID(JNIEnv* env, jclass clazz, const char* name, const char* sig)
{
	return (*env)->GetStaticMethodID(env, clazz, name, sig);
}

jobject _GoJniCallStaticObjectMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallStaticObjectMethodA(env, clazz, methodID, args);
}

jboolean _GoJniCallStaticBooleanMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallStaticBooleanMethodA(env, clazz, methodID, args);
}

jbyte _GoJniCallStaticByteMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallStaticByteMethodA(env, clazz, methodID, args);
}

jchar _GoJniCallStaticCharMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallStaticCharMethodA(env, clazz, methodID, args);
}

jshort _GoJniCallStaticShortMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallStaticShortMethodA(env, clazz, methodID, args);
}

jint _GoJniCallStaticIntMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallStaticIntMethodA(env, clazz, methodID, args);
}

jlong _GoJniCallStaticLongMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallStaticLongMethodA(env, clazz, methodID, args);
}

jfloat _GoJniCallStaticFloatMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallStaticFloatMethodA(env, clazz, methodID, args);
}

jdouble _GoJniCallStaticDoubleMethodA(JNIEnv* env, jclass clazz, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallStaticDoubleMethodA(env, clazz, methodID, args);
}

void _GoJniCallStaticVoidMethodA(JNIEnv* env, jclass cls, jmethodID methodID, const jvalue* args)
{
	return (*env)->CallStaticVoidMethodA(env, cls, methodID, args);
}

jfieldID _GoJniGetStaticFieldID(JNIEnv* env, jclass clazz, const char* name, const char* sig)
{
	return (*env)->GetStaticFieldID(env, clazz, name, sig);
}

jobject _GoJniGetStaticObjectField(JNIEnv* env, jclass clazz, jfieldID fieldID)
{
	return (*env)->GetStaticObjectField(env, clazz, fieldID);
}

jboolean _GoJniGetStaticBooleanField(JNIEnv* env, jclass clazz, jfieldID fieldID)
{
	return (*env)->GetStaticBooleanField(env, clazz, fieldID);
}

jbyte _GoJniGetStaticByteField(JNIEnv* env, jclass clazz, jfieldID fieldID)
{
	return (*env)->GetStaticByteField(env, clazz, fieldID);
}

jchar _GoJniGetStaticCharField(JNIEnv* env, jclass clazz, jfieldID fieldID)
{
	return (*env)->GetStaticCharField(env, clazz, fieldID);
}

jshort _GoJniGetStaticShortField(JNIEnv* env, jclass clazz, jfieldID fieldID)
{
	return (*env)->GetStaticShortField(env, clazz, fieldID);
}

jint _GoJniGetStaticIntField(JNIEnv* env, jclass clazz, jfieldID fieldID)
{
	return (*env)->GetStaticIntField(env, clazz, fieldID);
}

jlong _GoJniGetStaticLongField(JNIEnv* env, jclass clazz, jfieldID fieldID)
{
	return (*env)->GetStaticLongField(env, clazz, fieldID);
}

jfloat _GoJniGetStaticFloatField(JNIEnv* env, jclass clazz, jfieldID fieldID)
{
	return (*env)->GetStaticFloatField(env, clazz, fieldID);
}

jdouble _GoJniGetStaticDoubleField(JNIEnv* env, jclass clazz, jfieldID fieldID)
{
	return (*env)->GetStaticDoubleField(env, clazz, fieldID);
}

void _GoJniSetStaticObjectField(JNIEnv* env, jclass clazz, jfieldID fieldID, jobject value)
{
	return (*env)->SetStaticObjectField(env, clazz, fieldID, value);
}

void _GoJniSetStaticBooleanField(JNIEnv* env, jclass clazz, jfieldID fieldID, jboolean value)
{
	return (*env)->SetStaticBooleanField(env, clazz, fieldID, value);
}

void _GoJniSetStaticByteField(JNIEnv* env, jclass clazz, jfieldID fieldID, jbyte value)
{
	return (*env)->SetStaticByteField(env, clazz, fieldID, value);
}

void _GoJniSetStaticCharField(JNIEnv* env, jclass clazz, jfieldID fieldID, jchar value)
{
	return (*env)->SetStaticCharField(env, clazz, fieldID, value);
}

void _GoJniSetStaticShortField(JNIEnv* env, jclass clazz, jfieldID fieldID, jshort value)
{
	return (*env)->SetStaticShortField(env, clazz, fieldID, value);
}

void _GoJniSetStaticIntField(JNIEnv* env, jclass clazz, jfieldID fieldID, jint value)
{
	return (*env)->SetStaticIntField(env, clazz, fieldID, value);
}

void _GoJniSetStaticLongField(JNIEnv* env, jclass clazz, jfieldID fieldID, jlong value)
{
	return (*env)->SetStaticLongField(env, clazz, fieldID, value);
}

void _GoJniSetStaticFloatField(JNIEnv* env, jclass clazz, jfieldID fieldID, jfloat value)
{
	return (*env)->SetStaticFloatField(env, clazz, fieldID, value);
}

void _GoJniSetStaticDoubleField(JNIEnv* env, jclass clazz, jfieldID fieldID, jdouble value)
{
	return (*env)->SetStaticDoubleField(env, clazz, fieldID, value);
}

jstring _GoJniNewString(JNIEnv* env, const jchar* unicode, jsize len)
{
	return (*env)->NewString(env, unicode, len);
}

jsize _GoJniGetStringLength(JNIEnv* env, jstring str)
{
	return (*env)->GetStringLength(env, str);
}

const jchar* _GoJniGetStringChars(JNIEnv* env, jstring str, jboolean* isCopy)
{
	return (*env)->GetStringChars(env, str, isCopy);
}

void _GoJniReleaseStringChars(JNIEnv* env, jstring str, const jchar* chars)
{
	return (*env)->ReleaseStringChars(env, str, chars);
}

jstring _GoJniNewStringUTF(JNIEnv* env, const char* utf)
{
	return (*env)->NewStringUTF(env, utf);
}

jsize _GoJniGetStringUTFLength(JNIEnv* env, jstring str)
{
	return (*env)->GetStringUTFLength(env, str);
}

const char* _GoJniGetStringUTFChars(JNIEnv* env, jstring str, jboolean* isCopy)
{
	return (*env)->GetStringUTFChars(env, str, isCopy);
}

void _GoJniReleaseStringUTFChars(JNIEnv* env, jstring str, const char* chars)
{
	return (*env)->ReleaseStringUTFChars(env, str, chars);
}

jsize _GoJniGetArrayLength(JNIEnv* env, jarray array)
{
	return (*env)->GetArrayLength(env, array);
}

jobjectArray _GoJniNewObjectArray(JNIEnv* env, jsize len, jclass clazz, jobject init)
{
	return (*env)->NewObjectArray(env, len, clazz, init);
}

jobject _GoJniGetObjectArrayElement(JNIEnv* env, jobjectArray array, jsize index)
{
	return (*env)->GetObjectArrayElement(env, array, index);
}

void _GoJniSetObjectArrayElement(JNIEnv* env, jobjectArray array, jsize index, jobject val)
{
	return (*env)->SetObjectArrayElement(env, array, index, val);
}

jbooleanArray _GoJniNewBooleanArray(JNIEnv* env, jsize len)
{
	return (*env)->NewBooleanArray(env, len);
}

jbyteArray _GoJniNewByteArray(JNIEnv* env, jsize len)
{
	return (*env)->NewByteArray(env, len);
}

jcharArray _GoJniNewCharArray(JNIEnv* env, jsize len)
{
	return (*env)->NewCharArray(env, len);
}

jshortArray _GoJniNewShortArray(JNIEnv* env, jsize len)
{
	return (*env)->NewShortArray(env, len);
}

jintArray _GoJniNewIntArray(JNIEnv* env, jsize len)
{
	return (*env)->NewIntArray(env, len);
}

jlongArray _GoJniNewLongArray(JNIEnv* env, jsize len)
{
	return (*env)->NewLongArray(env, len);
}

jfloatArray _GoJniNewFloatArray(JNIEnv* env, jsize len)
{
	return (*env)->NewFloatArray(env, len);
}

jdoubleArray _GoJniNewDoubleArray(JNIEnv* env, jsize len)
{
	return (*env)->NewDoubleArray(env, len);
}

jboolean* _GoJniGetBooleanArrayElements(JNIEnv* env, jbooleanArray array, jboolean* isCopy)
{
	return (*env)->GetBooleanArrayElements(env, array, isCopy);
}

jbyte* _GoJniGetByteArrayElements(JNIEnv* env, jbyteArray array, jboolean* isCopy)
{
	return (*env)->GetByteArrayElements(env, array, isCopy);
}

jchar* _GoJniGetCharArrayElements(JNIEnv* env, jcharArray array, jboolean* isCopy)
{
	return (*env)->GetCharArrayElements(env, array, isCopy);
}

jshort* _GoJniGetShortArrayElements(JNIEnv* env, jshortArray array, jboolean* isCopy)
{
	return (*env)->GetShortArrayElements(env, array, isCopy);
}

jint* _GoJniGetIntArrayElements(JNIEnv* env, jintArray array, jboolean* isCopy)
{
	return (*env)->GetIntArrayElements(env, array, isCopy);
}

jlong* _GoJniGetLongArrayElements(JNIEnv* env, jlongArray array, jboolean* isCopy)
{
	return (*env)->GetLongArrayElements(env, array, isCopy);
}

jfloat* _GoJniGetFloatArrayElements(JNIEnv* env, jfloatArray array, jboolean* isCopy)
{
	return (*env)->GetFloatArrayElements(env, array, isCopy);
}

jdouble* _GoJniGetDoubleArrayElements(JNIEnv* env, jdoubleArray array, jboolean* isCopy)
{
	return (*env)->GetDoubleArrayElements(env, array, isCopy);
}

void _GoJniReleaseBooleanArrayElements(JNIEnv* env, jbooleanArray array, jboolean* elems, jint mode)
{
	return (*env)->ReleaseBooleanArrayElements(env, array, elems, mode);
}

void _GoJniReleaseByteArrayElements(JNIEnv* env, jbyteArray array, jbyte* elems, jint mode)
{
	return (*env)->ReleaseByteArrayElements(env, array, elems, mode);
}

void _GoJniReleaseCharArrayElements(JNIEnv* env, jcharArray array, jchar* elems, jint mode)
{
	return (*env)->ReleaseCharArrayElements(env, array, elems, mode);
}

void _GoJniReleaseShortArrayElements(JNIEnv* env, jshortArray array, jshort* elems, jint mode)
{
	return (*env)->ReleaseShortArrayElements(env, array, elems, mode);
}

void _GoJniReleaseIntArrayElements(JNIEnv* env, jintArray array, jint* elems, jint mode)
{
	return (*env)->ReleaseIntArrayElements(env, array, elems, mode);
}

void _GoJniReleaseLongArrayElements(JNIEnv* env, jlongArray array, jlong* elems, jint mode)
{
	return (*env)->ReleaseLongArrayElements(env, array, elems, mode);
}

void _GoJniReleaseFloatArrayElements(JNIEnv* env, jfloatArray array, jfloat* elems, jint mode)
{
	return (*env)->ReleaseFloatArrayElements(env, array, elems, mode);
}

void _GoJniReleaseDoubleArrayElements(JNIEnv* env, jdoubleArray array, jdouble* elems, jint mode)
{
	return (*env)->ReleaseDoubleArrayElements(env, array, elems, mode);
}

void _GoJniGetBooleanArrayRegion(JNIEnv* env, jbooleanArray array, jsize start, jsize l, jboolean* buf)
{
	return (*env)->GetBooleanArrayRegion(env, array, start, l, buf);
}

void _GoJniGetByteArrayRegion(JNIEnv* env, jbyteArray array, jsize start, jsize len, jbyte* buf)
{
	return (*env)->GetByteArrayRegion(env, array, start, len, buf);
}

void _GoJniGetCharArrayRegion(JNIEnv* env, jcharArray array, jsize start, jsize len, jchar* buf)
{
	return (*env)->GetCharArrayRegion(env, array, start, len, buf);
}

void _GoJniGetShortArrayRegion(JNIEnv* env, jshortArray array, jsize start, jsize len, jshort* buf)
{
	return (*env)->GetShortArrayRegion(env, array, start, len, buf);
}

void _GoJniGetIntArrayRegion(JNIEnv* env, jintArray array, jsize start, jsize len, jint* buf)
{
	return (*env)->GetIntArrayRegion(env, array, start, len, buf);
}

void _GoJniGetLongArrayRegion(JNIEnv* env, jlongArray array, jsize start, jsize len, jlong* buf)
{
	return (*env)->GetLongArrayRegion(env, array, start, len, buf);
}

void _GoJniGetFloatArrayRegion(JNIEnv* env, jfloatArray array, jsize start, jsize len, jfloat* buf)
{
	return (*env)->GetFloatArrayRegion(env, array, start, len, buf);
}

void _GoJniGetDoubleArrayRegion(JNIEnv* env, jdoubleArray array, jsize start, jsize len, jdouble* buf)
{
	return (*env)->GetDoubleArrayRegion(env, array, start, len, buf);
}

void _GoJniSetBooleanArrayRegion(JNIEnv* env, jbooleanArray array, jsize start, jsize l, const jboolean* buf)
{
	return (*env)->SetBooleanArrayRegion(env, array, start, l, buf);
}

void _GoJniSetByteArrayRegion(JNIEnv* env, jbyteArray array, jsize start, jsize len, const jbyte* buf)
{
	return (*env)->SetByteArrayRegion(env, array, start, len, buf);
}

void _GoJniSetCharArrayRegion(JNIEnv* env, jcharArray array, jsize start, jsize len, const jchar* buf)
{
	return (*env)->SetCharArrayRegion(env, array, start, len, buf);
}

void _GoJniSetShortArrayRegion(JNIEnv* env, jshortArray array, jsize start, jsize len, const jshort* buf)
{
	return (*env)->SetShortArrayRegion(env, array, start, len, buf);
}

void _GoJniSetIntArrayRegion(JNIEnv* env, jintArray array, jsize start, jsize len, const jint* buf)
{
	return (*env)->SetIntArrayRegion(env, array, start, len, buf);
}

void _GoJniSetLongArrayRegion(JNIEnv* env, jlongArray array, jsize start, jsize len, const jlong* buf)
{
	return (*env)->SetLongArrayRegion(env, array, start, len, buf);
}

void _GoJniSetFloatArrayRegion(JNIEnv* env, jfloatArray array, jsize start, jsize len, const jfloat* buf)
{
	return (*env)->SetFloatArrayRegion(env, array, start, len, buf);
}

void _GoJniSetDoubleArrayRegion(JNIEnv* env, jdoubleArray array, jsize start, jsize len, const jdouble* buf)
{
	return (*env)->SetDoubleArrayRegion(env, array, start, len, buf);
}

jint _GoJniRegisterNatives(JNIEnv* env, jclass clazz, const JNINativeMethod* methods, jint nMethods)
{
	return (*env)->RegisterNatives(env, clazz, methods, nMethods);
}

jint _GoJniUnregisterNatives(JNIEnv* env, jclass clazz)
{
	return (*env)->UnregisterNatives(env, clazz);
}

jint _GoJniMonitorEnter(JNIEnv* env, jobject obj)
{
	return (*env)->MonitorEnter(env, obj);
}

jint _GoJniMonitorExit(JNIEnv* env, jobject obj)
{
	return (*env)->MonitorExit(env, obj);
}

jint _GoJniGetJavaVM(JNIEnv* env, JavaVM* * vm)
{
	return (*env)->GetJavaVM(env, vm);
}

void _GoJniGetStringRegion(JNIEnv* env, jstring str, jsize start, jsize len, jchar* buf)
{
	return (*env)->GetStringRegion(env, str, start, len, buf);
}

void _GoJniGetStringUTFRegion(JNIEnv* env, jstring str, jsize start, jsize len, char* buf)
{
	return (*env)->GetStringUTFRegion(env, str, start, len, buf);
}

void* _GoJniGetPrimitiveArrayCritical(JNIEnv* env, jarray array, jboolean* isCopy)
{
	return (*env)->GetPrimitiveArrayCritical(env, array, isCopy);
}

void _GoJniReleasePrimitiveArrayCritical(JNIEnv* env, jarray array, void* carray, jint mode)
{
	return (*env)->ReleasePrimitiveArrayCritical(env, array, carray, mode);
}

const jchar* _GoJniGetStringCritical(JNIEnv* env, jstring string, jboolean* isCopy)
{
	return (*env)->GetStringCritical(env, string, isCopy);
}

void _GoJniReleaseStringCritical(JNIEnv* env, jstring string, const jchar* cstring)
{
	return (*env)->ReleaseStringCritical(env, string, cstring);
}

jweak _GoJniNewWeakGlobalRef(JNIEnv* env, jobject obj)
{
	return (*env)->NewWeakGlobalRef(env, obj);
}

void _GoJniDeleteWeakGlobalRef(JNIEnv* env, jweak ref)
{
	return (*env)->DeleteWeakGlobalRef(env, ref);
}

jboolean _GoJniExceptionCheck(JNIEnv* env)
{
	return (*env)->ExceptionCheck(env);
}

jobject _GoJniNewDirectByteBuffer(JNIEnv* env, void* address, jlong capacity)
{
	return (*env)->NewDirectByteBuffer(env, address, capacity);
}

void* _GoJniGetDirectBufferAddress(JNIEnv* env, jobject buf)
{
	return (*env)->GetDirectBufferAddress(env, buf);
}

jlong _GoJniGetDirectBufferCapacity(JNIEnv* env, jobject buf)
{
	return (*env)->GetDirectBufferCapacity(env, buf);
}

jobjectRefType _GoJniGetObjectRefType(JNIEnv* env, jobject obj)
{
	return (*env)->GetObjectRefType(env, obj);
}

