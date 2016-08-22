package jni

/*
#include <stdlib.h>
*/
import "C"

import (
	"unicode/utf16"
)

// JStringFromGoString creates a JString from a Go string.
// It is the high-level API for the NewString JNI function.
// errors: NilReturnValueError, OutOfMemoryError
func JStringFromGoString(s string) (str JString, err error) {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return JStringFromGoStringE(s, env)
}

// JStringFromGoStringE is like JStringFromGoString, but it accepts a custom JNIEnv.
// It calls the underlying functions with the provided JNIEnv.
// errors: NilReturnValueError, OutOfMemoryError
func JStringFromGoStringE(s string, env JNIEnv) (str JString, err error) {
	return env.NewString(s)
}

// String creates a Go string from a JString.
// It implements the fmt.Stringer interface.
// It calls StringRaw for the actual work.
// If the returned "err" field is not nil, it panics.
func (str JString) String() string {
	s, err := str.StringRaw()
	if err != nil {
		panic(err)
	}
	return s
}

// StringE is like String, but it accepts a custom JNIEnv.
// It calls the underlying functions with the provided JNIEnv.
func (str JString) StringE(env JNIEnv) string {
	s, err := str.StringRawE(env)
	if err != nil {
		panic(err)
	}
	return s
}

// StringRaw creates a Go string from a JString.
// errors:
func (str JString) StringRaw() (s string, err error) {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return str.StringRawE(env)
}

// StringRawE is like StringRaw, but it accepts a custom JNIEnv.
// It calls the underlying functions with the provided JNIEnv.
// errors:
func (str JString) StringRawE(env JNIEnv) (s string, err error) {
	var result string
	consumer := func(chars []uint16, _ bool) {
		runes := utf16.Decode(chars)
		result = string(runes)
	}
	if err := str.WithUTF16CharsE(consumer, env); err != nil {
		return "", err
	}
	return result, nil
}

// UTF16Len returns the length of the JString in number of UTF-16 characters.
// It is the high-level API for the GetStringLength JNI function.
func (str JString) UTF16Len() int {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return str.UTF16LenE(env)
}

// UTF16LenE is like UTF16Len, but it accepts a custom JNIEnv.
// It calls the underlying functions with the provided JNIEnv.
func (str JString) UTF16LenE(env JNIEnv) int {
	return env.GetStringLength(str)
}

// WithUTF16Chars lets the caller to work with the UTF-16 characters contained in the JString.
// It is the high-level API for the GetStringChars and ReleaseStringChars JNI functions.
// In the callback function (consumer),
// "chars" is the slice containing the UTF-16 characters,
// and "isCopy" indicates whether a copy of the C array is made by the JVM,
// or the underlying C array of the Java String is returned directly.
// "chars" points directly to whatever value the underlying call returns,
// so if "isCopy" is false, the caller MUST NOT modify any element in "chars".
// errors: NilReturnValueError
func (str JString) WithUTF16Chars(consumer func(chars []uint16, isCopy bool)) error {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return str.WithUTF16CharsE(consumer, env)
}

// WithUTF16CharsE is like WithUTF16Chars, but it accepts a custom JNIEnv.
// It calls the underlying functions with the provided JNIEnv.
// errors: NilReturnValueError
func (str JString) WithUTF16CharsE(consumer func(chars []uint16, isCopy bool), env JNIEnv) error {
	chars, isCopy, err := env.GetStringChars(str)
	if err != nil {
		return err
	}
	defer env.ReleaseStringChars(str, chars)
	consumer(chars, isCopy)
	return nil
}

// ModifiedUTF8Len returns the length of the JString in number of "modified UTF-8" characters.
// It is the high-level API for the GetStringUTFLength JNI function.
func (str JString) ModifiedUTF8Len() int {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return str.ModifiedUTF8LenE(env)
}

// ModifiedUTF8LenE is like ModifiedUTF8Len, but it accepts a custom JNIEnv.
// It calls the underlying functions with the provided JNIEnv.
func (str JString) ModifiedUTF8LenE(env JNIEnv) int {
	return env.GetStringUTFLength(str)
}

// WithModifiedUTF8Bytes lets the caller to work with the "modified UTF-8" characters contained in the JString.
// It is the high-level API for the GetStringUTFChars and ReleaseStringUTFChars JNI functions.
// In the callback function (consumer),
// "chars" is the slice containing the "modified UTF-8" characters,
// and "isCopy" indicates whether a copy of the C array is made by the JVM,
// or the underlying C array of the Java String is returned directly.
// "chars" points directly to whatever value the underlying call returns,
// so if "isCopy" is false, the caller MUST NOT modify any element in "chars".
// errors: NilReturnValueError
func (str JString) WithModifiedUTF8Bytes(consumer func(bytes []byte, isCopy bool)) error {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return str.WithModifiedUTF8BytesE(consumer, env)
}

// WithModifiedUTF8BytesE is like WithModifiedUTF8Bytes, but it accepts a custom JNIEnv.
// It calls the underlying functions with the provided JNIEnv.
// errors: NilReturnValueError
func (str JString) WithModifiedUTF8BytesE(consumer func(bytes []byte, isCopy bool), env JNIEnv) error {
	// GetStringUTFChars
	bytes, isCopy, err := env.GetStringUTFChars(str)
	if err != nil {
		return err
	}
	defer env.ReleaseStringUTFChars(str, bytes)
	consumer(bytes, isCopy)
	return nil
}

// TODO: Doc
// errors: StringIndexOutOfBoundsException
func (str JString) UTF16Region(start, length int, buf []uint16) error {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return str.UTF16RegionE(start, length, buf, env)
}

// TODO: Doc
// errors: StringIndexOutOfBoundsException
func (str JString) UTF16RegionE(start, length int, buf []uint16, env JNIEnv) error {
	return env.GetStringRegion(str, start, length, buf)
}

// TODO: Doc
// errors: StringIndexOutOfBoundsException
func (str JString) ModifiedUTF8Region(start, length int, buf []byte) error {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return str.ModifiedUTF8RegionE(start, length, buf, env)
}

// TODO: Doc
// errors: StringIndexOutOfBoundsException
func (str JString) ModifiedUTF8RegionE(start, length int, buf []byte, env JNIEnv) error {
	return env.GetStringUTFRegion(str, start, length, buf)
}

// TODO: Doc
// errors: NilReturnValueError
func (str JString) WithUTF16CharsCritical(consumer func(chars []uint16, isCopy bool)) error {
	env := GoJNIEnv(GetDefaultJNIEnv())
	return str.WithUTF16CharsCriticalE(consumer, env)
}

// TODO: Doc
// errors: NilReturnValueError
func (str JString) WithUTF16CharsCriticalE(consumer func(chars []uint16, isCopy bool), env JNIEnv) error {
	chars, isCopy, err := env.GetStringCritical(str)
	if err != nil {
		return err
	}
	defer env.ReleaseStringCritical(str, chars)
	consumer(chars, isCopy)
	return nil
}
