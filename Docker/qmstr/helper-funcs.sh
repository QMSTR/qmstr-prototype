#!/bin/bash

function quit(){
    curl http://localhost:9000/quit
}

function dump(){
    curl http://localhost:9000/dump
}

function printtragets(){
    curl http://localhost:9000/linkedtargets
}
