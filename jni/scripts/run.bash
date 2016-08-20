#!/bin/bash

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

java -classpath "${DIR}/../demo" -Djava.library.path="${DIR}/../demo" Main
