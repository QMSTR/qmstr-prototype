#!/bin/bash
set -e

function is_debug(){
    if [ ! -z "$QMSTR_DEBUG" ]; then
        true
    else
        false
    fi
}

function build(){
    echo "Building project"
    cd "${QMSTR_BUILD_DIR}"
    if is_debug; then
        qmstr-master -v &
    else
        qmstr-master &
    fi
    export PATH=/qmstr-wrapper:$PATH
    export CC=/qmstr-wrapper/gcc
    export CXX=/qmstr-wrapper/g++
    export CMAKE_LINKER=gcc
    exec "$@"
}

if [ "$1" = 'dev' ]; then
    if [ "$2" = 'build' ]; then
        echo "Building QMSTR"
        shift 2
        source /helper-funcs.sh
        build_qmstr
    fi
fi

if [ -n "$1" ]; then
    if is_debug; then
        echo "source /helper-funcs.sh" > ~/.bashrc
    fi
    build "$@" 
fi
