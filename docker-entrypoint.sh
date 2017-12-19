#!/bin/bash
set -e

function build(){
    echo "Building project"
    cd "${QMSTR_BUILD_DIR}"
    if [ $QMSTR_DEBUG = "true" ]; then
        qmstr-master -v &
    else
        qmstr-master &
    fi
    export PATH=/qmstr-wrapper:$PATH
    export CC=/qmstr-wrapper/gcc
    export CXX=/qmstr-wrapper/g++
    exec "$@"
}

if [ "$1" = 'dev' ]; then
    if [ "$2" = 'build' ]; then
        echo "Building QMSTR"
        shift 2
        go get qmstr-prototype/qmstr/qmstr-master
        go install qmstr-prototype/qmstr/qmstr-master
        go install qmstr-prototype/qmstr/qmstr-wrapper
    fi
fi

if [ -n "$1" ]; then
    build "$@" 
fi
