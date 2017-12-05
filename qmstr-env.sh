#!/bin/sh

rm -fr /vagrant/bin
mkdir -p /vagrant/bin
ln -s /home/vagrant/go/bin/qmstr-wrapper /vagrant/bin/gcc
export PATH=/vagrant/bin:$PATH

echo "Everything is ready!!"
echo "You can now run make"
