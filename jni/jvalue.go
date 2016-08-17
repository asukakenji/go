package main

/*
// See src/os/user/lookup_unix.go for usage of "static"

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

// jvalue From Any

static jvalue _GoJniJValueFromJBoolean(jboolean z)
{
	jvalue v = {.z = z};
	return v;
}

static jvalue _GoJniJValueFromJByte(jbyte b)
{
	jvalue v = {.b = b};
	return v;
}

static jvalue _GoJniJValueFromJChar(jchar c)
{
	jvalue v = {.c = c};
	return v;
}

static jvalue _GoJniJValueFromJShort(jshort s)
{
	jvalue v = {.s = s};
	return v;
}

static jvalue _GoJniJValueFromJInt(jint i)
{
	jvalue v = {.i = i};
	return v;
}

static jvalue _GoJniJValueFromJLong(jlong j)
{
	jvalue v = {.j = j};
	return v;
}

static jvalue _GoJniJValueFromJFloat(jfloat f)
{
	jvalue v = {.f = f};
	return v;
}

static jvalue _GoJniJValueFromJDouble(jdouble d)
{
	jvalue v = {.d = d};
	return v;
}

static jvalue _GoJniJValueFromJObject(jobject l)
{
	jvalue v = {.l = l};
	return v;
}

// jvalue To Any

static jboolean _GoJniJValueToJBoolean(jvalue v)
{
	return v.z;
}

static jbyte _GoJniJValueToJByte(jvalue v)
{
	return v.b;
}

static jchar _GoJniJValueToJChar(jvalue v)
{
	return v.c;
}

static jshort _GoJniJValueToJShort(jvalue v)
{
	return v.s;
}

static jint _GoJniJValueToJInt(jvalue v)
{
	return v.i;
}

static jlong _GoJniJValueToJLong(jvalue v)
{
	return v.j;
}

static jfloat _GoJniJValueToJFloat(jvalue v)
{
	return v.f;
}

static jdouble _GoJniJValueToJDouble(jvalue v)
{
	return v.d;
}

static jobject _GoJniJValueToJObject(jvalue v)
{
	return v.l;
}
*/
import "C"

// jvalue From Any

func JValueFromJBoolean(z C.jboolean) C.jvalue {
	return C._GoJniJValueFromJBoolean(z)
}

func JValueFromJByte(b C.jbyte) C.jvalue {
	return C._GoJniJValueFromJByte(b)
}

func JValueFromJChar(c C.jchar) C.jvalue {
	return C._GoJniJValueFromJChar(c)
}

func JValueFromJShort(s C.jshort) C.jvalue {
	return C._GoJniJValueFromJShort(s)
}

func JValueFromJInt(i C.jint) C.jvalue {
	return C._GoJniJValueFromJInt(i)
}

func JValueFromJLong(j C.jlong) C.jvalue {
	return C._GoJniJValueFromJLong(j)
}

func JValueFromJFloat(f C.jfloat) C.jvalue {
	return C._GoJniJValueFromJFloat(f)
}

func JValueFromJDouble(d C.jdouble) C.jvalue {
	return C._GoJniJValueFromJDouble(d)
}

func JValueFromJObject(l C.jobject) C.jvalue {
	return C._GoJniJValueFromJObject(l)
}

// jvalue To Any

func JValueToJBoolean(v C.jvalue) C.jboolean {
	return C._GoJniJValueToJBoolean(v)
}

func JValueToJByte(v C.jvalue) C.jbyte {
	return C._GoJniJValueToJByte(v)
}

func JValueToJChar(v C.jvalue) C.jchar {
	return C._GoJniJValueToJChar(v)
}

func JValueToJShort(v C.jvalue) C.jshort {
	return C._GoJniJValueToJShort(v)
}

func JValueToJInt(v C.jvalue) C.jint {
	return C._GoJniJValueToJInt(v)
}

func JValueToJLong(v C.jvalue) C.jlong {
	return C._GoJniJValueToJLong(v)
}

func JValueToJFloat(v C.jvalue) C.jfloat {
	return C._GoJniJValueToJFloat(v)
}

func JValueToJDouble(v C.jvalue) C.jdouble {
	return C._GoJniJValueToJDouble(v)
}

func JValueToJObject(v C.jvalue) C.jobject {
	return C._GoJniJValueToJObject(v)
}
