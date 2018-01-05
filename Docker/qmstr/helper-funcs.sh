#!/bin/bash

function multiecho(){
    str=$1
    num=$2
    v=$(printf "%-${num}s" "$str")
    echo "${v// /${str}}"
}

function printheader(){
    str=$1
    echo
    echo "$str"
    multiecho "=" ${#str}
}

function quit(){
    curl http://localhost:9000/quit
}

function dump(){
    curl http://localhost:9000/dump
}

function printtargets(){
    curl http://localhost:9000/linkedtargets
}

function report(){
    curl http://localhost:9000/report?id=$1
}

function build_cmake(){
    cd /build
    rm -fr build
    mkdir build
    cd build
    cmake ..
    make
}

function build_qmstr(){
    PROGS="qmstr-prototype/qmstr/qmstr-master qmstr-prototype/qmstr/qmstr-wrapper"
    for p in $PROGS; do
        printheader "Installing dependencies for $p"
        go get -t $p
        printheader "Building $p"
        go build -v $p
        printheader "Installing $p"
        go install -v $p
    done
}

function run_unittests() {
    printheader "Running unit tests"
    for p in $(ls /qmstr/qmstr); do
        printheader "Testing package $p"
        go get -t qmstr-prototype/qmstr/$p
        go test -i qmstr-prototype/qmstr/$p
        go test -v qmstr-prototype/qmstr/$p
    done
}
