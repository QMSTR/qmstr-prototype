#!/bin/sh
echo Installing required packages...

add-apt-repository ppa:gophers/archive
apt-get update
# Go binaries will be installed in /usr/lib/go-1.9/bin:
apt-get install -y golang-1.9-go
