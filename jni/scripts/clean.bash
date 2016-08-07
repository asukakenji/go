#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

rm -f "${DIR}/../libtest.jnilib" "${DIR}/../libtest.h" "${DIR}/../Main.h" "${DIR}/Main.class"
