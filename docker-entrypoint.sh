#!/bin/bash
set -e

function build(){
    echo "Building project"
    cd "${QMSTR_BUILD_DIR}"
    qmstr-master &
    PATH=/qmstr-wrapper:$PATH
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