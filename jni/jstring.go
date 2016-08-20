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

/*
func JStringFromGoString(s string) (str JString, err error) {
	env := GetDefaultJNIEnv()
	return JStringFromStringE(s, env)
}
*/

func JStringFromGoStringE(s string, env JNIEnv) (str JString, err error) {
	return env.NewString(s)
}

// String implements the fmt.Stringer interface
/*
func (str JString) String() string {
	env := GetDefaultJNIEnv()
	return StringE(env)
}
*/

func (str JString) StringE(env JNIEnv) string {
	s, err := str.GoStringE(env)
	if err != nil {
		panic(err)
	}
	return s
}

/*
func (str JString) GoString() (s string, err error) {
	env := GetDefaultJNIEnv()
	return GoStringE(env)
}
*/

func (str JString) GoStringE(env JNIEnv) (s string, err error) {
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

/*
func (str JString) Len() int {
	env := GetDefaultJNIEnv()
	return str.LenE(env)
}
*/

func (str JString) LenE(env JNIEnv) int {
	return env.GetStringLength(str)
}

/*
func (str JString) WithChars(consumer func([]uint16, bool)) error {
	env := GetDefaultJNIEnv()
	return str.WithCharsE(consumer, env)
}
*/

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

/*
func (str JString) UTFLen() int {
	env := GetDefaultJNIEnv()
	return str.UTFLenE(env)
}
*/

func (str JString) UTFLenE(env JNIEnv) int {
	return env.GetStringUTFLength(str)
}

/*
func (str JString) WithUTFChars(consumer func(chars []byte, isCopy bool)) {
	env := GetDefaultJNIEnv()
	return str.WithUTFCharsE(consumer, env)
}
*/

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
