package main

// #include "jvalue.h"
import "C"

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

func JBooleanToJValue(z C.jboolean) C.jvalue {
	return C._GoJniJBooleanToJValue(z)
}

func JByteToJValue(b C.jbyte) C.jvalue {
	return C._GoJniJByteToJValue(b)
}

func JCharToJValue(c C.jchar) C.jvalue {
	return C._GoJniJCharToJValue(c)
}

func JShortToJValue(s C.jshort) C.jvalue {
	return C._GoJniJShortToJValue(s)
}

func JIntToJValue(i C.jint) C.jvalue {
	return C._GoJniJIntToJValue(i)
}

func JLongToJValue(j C.jlong) C.jvalue {
	return C._GoJniJLongToJValue(j)
}

func JFloatToJValue(f C.jfloat) C.jvalue {
	return C._GoJniJFloatToJValue(f)
}

func JDoubleToJValue(d C.jdouble) C.jvalue {
	return C._GoJniJDoubleToJValue(d)
}

func JObjectToJValue(l C.jobject) C.jvalue {
	return C._GoJniJObjectToJValue(l)
}

func (value JValue) Bool() bool {
	return GoBool(JValueToJBoolean(value.peer))
}

func (value JValue) Int8() int8 {
	return GoInt8(JValueToJByte(value.peer))
}

func (value JValue) Uint16() uint16 {
	return GoUint16(JValueToJChar(value.peer))
}

func (value JValue) Int16() int16 {
	return GoInt16(JValueToJShort(value.peer))
}

func (value JValue) Int32() int32 {
	return GoInt32(JValueToJInt(value.peer))
}

func (value JValue) Int64() int64 {
	return GoInt64(JValueToJLong(value.peer))
}

func (value JValue) Float32() float32 {
	return GoFloat32(JValueToJFloat(value.peer))
}

func (value JValue) Float64() float64 {
	return GoFloat64(JValueToJDouble(value.peer))
}

func (value JValue) JObject() JObject {
	return NewJObject(JValueToJObject(value.peer))
}

func JValueFromBool(z bool) JValue {
	return JValue{JBooleanToJValue(JavaBoolean(z))}
}

func JValueFromInt8(b int8) JValue {
	return JValue{JByteToJValue(JavaByte(b))}
}

func JValueFromUint16(c uint16) JValue {
	return JValue{JCharToJValue(JavaChar(c))}
}

func JValueFromInt16(s int16) JValue {
	return JValue{JShortToJValue(JavaShort(s))}
}

func JValueFromInt32(i int32) JValue {
	return JValue{JIntToJValue(JavaInt(i))}
}

func JValueFromInt64(j int64) JValue {
	return JValue{JLongToJValue(JavaLong(j))}
}

func JValueFromFloat32(f float32) JValue {
	return JValue{JFloatToJValue(JavaFloat(f))}
}

func JValueFromFloat64(d float64) JValue {
	return JValue{JDoubleToJValue(JavaDouble(d))}
}

func JValueFromJObject(l JObject) JValue {
	return JValue{JObjectToJValue(l.Peer())}
}
