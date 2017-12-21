#!/bin/bash
# build qmstr
go get qmstr-prototype/qmstr/qmstr-master
go install qmstr-prototype/qmstr/qmstr-master
go install qmstr-prototype/qmstr/qmstr-wrapper

# configure bash:
cd ~/bash
./configure
make clean
# configure PATH for QMSTR
if [[ $PATH != *"vagrant/bin"* ]];then
    export PATH=/vagrant/bin:$PATH
fi
# run the build
~/go/bin/qmstr-master &
make -j5

curl http://localhost:9000/report?id=bash

curl http://localhost:9000/quit

