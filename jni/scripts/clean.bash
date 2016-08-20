#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

rm -f "${DIR}/../demo/libgojni.jnilib" "${DIR}/../demo/libgojni.h" "${DIR}/../demo/Main.h" "${DIR}/../demo/Main.class"
