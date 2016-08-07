#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

cat ${JAVA_HOME}/include/jni.h | awk -f ${DIR}/gen_jnienv_h.awk > ${DIR}/../jnienv.h
cat ${DIR}/../jnienv.h | awk -f ${DIR}/gen_jnienv_c.awk > ${DIR}/../jnienv.c

cat ${JAVA_HOME}/include/jni.h | awk -f ${DIR}/gen_jvalue_h.awk > ${DIR}/../jvalue.h
cat ${JAVA_HOME}/include/jni.h | awk -f ${DIR}/gen_jvalue_c.awk > ${DIR}/../jvalue.c
cat ${JAVA_HOME}/include/jni.h | awk -f ${DIR}/gen_jvalue_go.awk > ${DIR}/../jvalue.go
