#include "jvalue.h"

jboolean _GoJniJValueGetJBoolean(jvalue v)
{
	return v.z;
}

jbyte _GoJniJValueGetJByte(jvalue v)
{
	return v.b;
}

jchar _GoJniJValueGetJChar(jvalue v)
{
	return v.c;
}

jshort _GoJniJValueGetJShort(jvalue v)
{
	return v.s;
}

jint _GoJniJValueGetJInt(jvalue v)
{
	return v.i;
}

jlong _GoJniJValueGetJLong(jvalue v)
{
	return v.j;
}

jfloat _GoJniJValueGetJFloat(jvalue v)
{
	return v.f;
}

jdouble _GoJniJValueGetJDouble(jvalue v)
{
	return v.d;
}

jobject _GoJniJValueGetJObject(jvalue v)
{
	return v.l;
}

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
