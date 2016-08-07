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

jboolean _GoJniJValueToJBoolean(jvalue v);
jbyte _GoJniJValueToJByte(jvalue v);
jchar _GoJniJValueToJChar(jvalue v);
jshort _GoJniJValueToJShort(jvalue v);
jint _GoJniJValueToJInt(jvalue v);
jlong _GoJniJValueToJLong(jvalue v);
jfloat _GoJniJValueToJFloat(jvalue v);
jdouble _GoJniJValueToJDouble(jvalue v);
jobject _GoJniJValueToJObject(jvalue v);

jvalue _GoJniJBooleanToJValue(jboolean z);
jvalue _GoJniJByteToJValue(jbyte b);
jvalue _GoJniJCharToJValue(jchar c);
jvalue _GoJniJShortToJValue(jshort s);
jvalue _GoJniJIntToJValue(jint i);
jvalue _GoJniJLongToJValue(jlong j);
jvalue _GoJniJFloatToJValue(jfloat f);
jvalue _GoJniJDoubleToJValue(jdouble d);
jvalue _GoJniJObjectToJValue(jobject l);

#endif // #ifndef _GO_JNI_JVALUE_H_
