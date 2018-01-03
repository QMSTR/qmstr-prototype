#!/bin/bash

function quit(){
    curl http://localhost:9000/quit
}

function dump(){
    curl http://localhost:9000/dump
}

function printtargets(){
    curl http://localhost:9000/linkedtargets
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
    go get qmstr-prototype/qmstr/qmstr-master
    go install qmstr-prototype/qmstr/qmstr-master
    go install qmstr-prototype/qmstr/qmstr-wrapper
}
