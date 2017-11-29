#!/bin/sh

mkdir -p /vagrant/bin
ln -s /vagrant/qmstr/qmstr-compile/qmstr-compile /vagrant/bin/gcc
export PATH=/vagrant/bin:$PATH

echo "Everything is ready!!"
echo "You can now run make"