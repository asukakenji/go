#ifndef _GO_JNI_JVALUE_H_
#define _GO_JNI_JVALUE_H_

#include <jni.h>

// jni.h:
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

jvalue _GoJniJValueFromJBoolean(jboolean z);
jvalue _GoJniJValueFromJByte(jbyte b);
jvalue _GoJniJValueFromJChar(jchar c);
jvalue _GoJniJValueFromJShort(jshort s);
jvalue _GoJniJValueFromJInt(jint i);
jvalue _GoJniJValueFromJLong(jlong j);
jvalue _GoJniJValueFromJFloat(jfloat f);
jvalue _GoJniJValueFromJDouble(jdouble d);
jvalue _GoJniJValueFromJObject(jobject l);

jboolean _GoJniJValueToJBoolean(jvalue v);
jbyte _GoJniJValueToJByte(jvalue v);
jchar _GoJniJValueToJChar(jvalue v);
jshort _GoJniJValueToJShort(jvalue v);
jint _GoJniJValueToJInt(jvalue v);
jlong _GoJniJValueToJLong(jvalue v);
jfloat _GoJniJValueToJFloat(jvalue v);
jdouble _GoJniJValueToJDouble(jvalue v);
jobject _GoJniJValueToJObject(jvalue v);

#endif // #ifndef _GO_JNI_JVALUE_H_
