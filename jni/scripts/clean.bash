#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

rm -f "${DIR}/../libgojni.jnilib" "${DIR}/../libgojni.h" "${DIR}/../Main.h" "${DIR}/Main.class"
