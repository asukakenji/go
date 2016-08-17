package main

// #include "jvalue.h"
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
