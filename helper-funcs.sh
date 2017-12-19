#!/bin/bash

function quit(){
    curl http://localhost:8080/quit
}

function dump(){
    curl http://localhost:8080/dump
}

function printtragets(){
    curl http://localhost:8080/linkedtargets
}
