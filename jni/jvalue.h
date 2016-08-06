#ifndef _GO_JNI_JVALUE_H_
#define _GO_JNI_JVALUE_H_

#include <jni.h>

// jni.h
// typedef union jvalue {
//     jboolean z;
//     jbyte    b;
//     jchar    c;
//     jshort   s;
//     jint     i;
//     jlong    j;
//     jfloat   f;
//     jdouble  d;
//     jobject  l;
// } jvalue;

jboolean _GoJniJValueGetJBoolean(jvalue v);
jbyte _GoJniJValueGetJByte(jvalue v);
jchar _GoJniJValueGetJChar(jvalue v);
jshort _GoJniJValueGetJShort(jvalue v);
jint _GoJniJValueGetJInt(jvalue v);
jlong _GoJniJValueGetJLong(jvalue v);
jfloat _GoJniJValueGetJFloat(jvalue v);
jdouble _GoJniJValueGetJDouble(jvalue v);
jobject _GoJniJValueGetJObject(jvalue v);

jvalue _GoJniJValueFromJBoolean(jboolean z);
jvalue _GoJniJValueFromJByte(jbyte b);
jvalue _GoJniJValueFromJChar(jchar c);
jvalue _GoJniJValueFromJShort(jshort s);
jvalue _GoJniJValueFromJInt(jint i);
jvalue _GoJniJValueFromJLong(jlong j);
jvalue _GoJniJValueFromJFloat(jfloat f);
jvalue _GoJniJValueFromJDouble(jdouble d);
jvalue _GoJniJValueFromJObject(jobject l);

#endif // #ifndef _GO_JNI_JVALUE_H_
