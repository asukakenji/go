#include "jvalue.h"

jvalue _GoJniJValueFromJBoolean(jboolean z)
{
	jvalue v = {.z = z};
	return v;
}

jvalue _GoJniJValueFromJByte(jbyte b)
{
	jvalue v = {.b = b};
	return v;
}

jvalue _GoJniJValueFromJChar(jchar c)
{
	jvalue v = {.c = c};
	return v;
}

jvalue _GoJniJValueFromJShort(jshort s)
{
	jvalue v = {.s = s};
	return v;
}

jvalue _GoJniJValueFromJInt(jint i)
{
	jvalue v = {.i = i};
	return v;
}

jvalue _GoJniJValueFromJLong(jlong j)
{
	jvalue v = {.j = j};
	return v;
}

jvalue _GoJniJValueFromJFloat(jfloat f)
{
	jvalue v = {.f = f};
	return v;
}

jvalue _GoJniJValueFromJDouble(jdouble d)
{
	jvalue v = {.d = d};
	return v;
}

jvalue _GoJniJValueFromJObject(jobject l)
{
	jvalue v = {.l = l};
	return v;
}

jboolean _GoJniJValueToJBoolean(jvalue v)
{
	return v.z;
}

jbyte _GoJniJValueToJByte(jvalue v)
{
	return v.b;
}

jchar _GoJniJValueToJChar(jvalue v)
{
	return v.c;
}

jshort _GoJniJValueToJShort(jvalue v)
{
	return v.s;
}

jint _GoJniJValueToJInt(jvalue v)
{
	return v.i;
}

jlong _GoJniJValueToJLong(jvalue v)
{
	return v.j;
}

jfloat _GoJniJValueToJFloat(jvalue v)
{
	return v.f;
}

jdouble _GoJniJValueToJDouble(jvalue v)
{
	return v.d;
}

jobject _GoJniJValueToJObject(jvalue v)
{
	return v.l;
}
