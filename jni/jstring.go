package jni

/*
#include <stdlib.h>
*/
import "C"

import (
	"reflect"
	"unicode/utf16"
	"unsafe"
)

// JStringFromGoString() creates a JString from a Go string.
// It is the high-level API for the NewString() JNI function.
func JStringFromGoString(s string) (str JString, err error) {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return JStringFromGoStringE(s, env)
}

// JStringFromGoStringE() is like JStringFromGoString().
// It calls the underlying functions with a the provided JNIEnv.
func JStringFromGoStringE(s string, env JNIEnv) (str JString, err error) {
	return env.NewString(s)
}

// String() creates a Go string from a JString.
// It implements the fmt.Stringer interface.
// It calls StringRaw() for the actual work.
// If the returned "err" field is not nil, it panics.
func (str JString) String() string {
	s, err := str.StringRaw()
	if err != nil {
		panic(err)
	}
	return s
}

// StringE() is like String().
// It calls the underlying functions with a the provided JNIEnv.
func (str JString) StringE(env JNIEnv) string {
	s, err := str.StringRawE(env)
	if err != nil {
		panic(err)
	}
	return s
}

// StringRaw() creates a Go string from a JString.
func (str JString) StringRaw() (s string, err error) {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return str.StringRawE(env)
}

// StringRawE() is like StringRaw().
// It calls the underlying functions with a the provided JNIEnv.
func (str JString) StringRawE(env JNIEnv) (s string, err error) {
	var result string
	consumer := func(chars []uint16, _ bool) {
		runes := utf16.Decode(chars)
		result = string(runes)
	}
	if err := str.WithCharsE(consumer, env); err != nil {
		return "", err
	}
	return result, nil
}

// Len() returns the length of the JString in number of UTF-16 characters.
// It is the high-level API for the GetStringLength() JNI function.
func (str JString) Len() int {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return str.LenE(env)
}

// LenE() is like Len().
// It calls the underlying functions with a the provided JNIEnv.
func (str JString) LenE(env JNIEnv) int {
	return env.GetStringLength(str)
}

// WithChars() lets the caller to work with the UTF-16 characters contained in the JString.
// It is the high-level API for the GetStringChars() and ReleaseStringChars() JNI functions.
// In the callback function (consumer),
// "chars" is the slice containing the UTF-16 characters,
// and "isCopy" indicates whether a copy of the C array is made by the JVM,
// or the underlying C array of the Java String is returned directly.
// "chars" points directly to whatever value the underlying call returns,
// so if "isCopy" is false, the caller MUST NOT modify any element in "chars".
func (str JString) WithChars(consumer func(chars []uint16, isCopy bool)) error {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return str.WithCharsE(consumer, env)
}

// WithCharsE() is like WithChars().
// It calls the underlying functions with a the provided JNIEnv.
func (str JString) WithCharsE(consumer func(chars []uint16, isCopy bool), env JNIEnv) error {
	// GetStringChars()
	chars, isCopy, err := env.GetStringChars(str)
	if err != nil {
		return err
	}
	defer env.ReleaseStringChars(str, chars)

	// GetStringLength()
	_len := str.LenE(env)

	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(chars)),
		Len: _len,
		Cap: _len,
	}
	consumer(*((*[]uint16)(unsafe.Pointer(&header))), isCopy)
	return nil
}

// UTFLen() returns the length of the JString in number of "modified UTF-8" characters.
// It is the high-level API for the GetStringUTFLength() JNI function.
func (str JString) UTFLen() int {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return str.UTFLenE(env)
}

// UTFLenE() is like UTFLen().
// It calls the underlying functions with a the provided JNIEnv.
func (str JString) UTFLenE(env JNIEnv) int {
	return env.GetStringUTFLength(str)
}

// WithUTFChars() lets the caller to work with the "modified UTF-8" characters contained in the JString.
// It is the high-level API for the GetStringUTFChars() and ReleaseStringUTFChars() JNI functions.
// In the callback function (consumer),
// "chars" is the slice containing the "modified UTF-8" characters,
// and "isCopy" indicates whether a copy of the C array is made by the JVM,
// or the underlying C array of the Java String is returned directly.
// "chars" points directly to whatever value the underlying call returns,
// so if "isCopy" is false, the caller MUST NOT modify any element in "chars".
func (str JString) WithUTFChars(consumer func(chars []byte, isCopy bool)) error {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return str.WithUTFCharsE(consumer, env)
}

// WithUTFCharsE() is like WithUTFChars().
// It calls the underlying functions with a the provided JNIEnv.
func (str JString) WithUTFCharsE(consumer func(chars []byte, isCopy bool), env JNIEnv) error {
	// GetStringUTFChars()
	chars, isCopy, err := env.GetStringUTFChars(str)
	if err != nil {
		return err
	}
	defer env.ReleaseStringUTFChars(str, chars)

	// GetStringUTFLength()
	_len := str.UTFLenE(env)

	header := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(chars)),
		Len: _len,
		Cap: _len,
	}
	consumer(*((*[]byte)(unsafe.Pointer(&header))), isCopy)
	return nil
}

// TODO: GetStringRegion
// TODO: GetStringUTFRegion
// TODO: GetStringCritical
// TODO: ReleaseStringCritical
