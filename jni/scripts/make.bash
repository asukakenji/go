#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

#go help buildmode
#
# The 'go build' and 'go install' commands take a -buildmode argument which
# indicates which kind of object file is to be built. Currently supported values
# are:
# 
# 	-buildmode=archive
# 		Build the listed non-main packages into .a files. Packages named
# 		main are ignored.
# 
# 	-buildmode=c-archive
# 		Build the listed main package, plus all packages it imports,
# 		into a C archive file. The only callable symbols will be those
# 		functions exported using a cgo //export comment. Requires
# 		exactly one main package to be listed.
# 
# 	-buildmode=c-shared
# 		Build the listed main packages, plus all packages that they
# 		import, into C shared libraries. The only callable symbols will
# 		be those functions exported using a cgo //export comment.
# 		Non-main packages are ignored.
# 
# 	-buildmode=default
# 		Listed main packages are built into executables and listed
# 		non-main packages are built into .a files (the default
# 		behavior).
# 
# 	-buildmode=shared
# 		Combine all the listed non-main packages into a single shared
# 		library that will be used when building with the -linkshared
# 		option. Packages named main are ignored.
# 
# 	-buildmode=exe
# 		Build the listed main packages and everything they import into
# 		executables. Packages not named main are ignored.

javac -sourcepath "${DIR}" -d "${DIR}" "${DIR}/Main.java"
javah -classpath "${DIR}" -d "${DIR}/.." Main

# This does NOT work because *.c files have to be built.
# C sources cannot be built via listing as arguments.
# Therefore, a package is needed.
#CGO_CFLAGS="-I${JAVA_HOME}/include -I${JAVA_HOME}/include/darwin" go build -o libgojni.jnilib -buildmode=c-shared export.go

# Normal:
#go fmt github.com/asukakenji/go/jni
CGO_CFLAGS="-I${JAVA_HOME}/include -I${JAVA_HOME}/include/darwin" go build -o libgojni.jnilib -buildmode=c-shared github.com/asukakenji/go/jni

# For Debugging Linker Errors:
#CGO_CFLAGS="-I${JAVA_HOME}/include -I${JAVA_HOME}/include/darwin" go build -x -ldflags "-extldflags -v" -o libgojni.jnilib -buildmode=c-shared github.com/asukakenji/go/jni

# For Viewing Temporary Files:
#CGO_CFLAGS="-I${JAVA_HOME}/include -I${JAVA_HOME}/include/darwin" go build -o libgojni.jnilib -work -buildmode=c-shared github.com/asukakenji/go/jni
