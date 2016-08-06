#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

cat ${JAVA_HOME}/include/jni.h | awk -f ${DIR}/gen_jnienv_h.awk > ${DIR}/../jnienv.h
cat ${DIR}/../jnienv.h | awk -f ${DIR}/gen_jnienv_c.awk > ${DIR}/../jnienv.c
