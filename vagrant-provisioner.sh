#!/bin/sh
echo Installing required packages...

add-apt-repository ppa:gophers/archive
apt-get update
# Go binaries will be installed in /usr/lib/go-1.9/bin:
apt-get install -y golang-1.9-go autoconf

echo 'PATH=/usr/lib/go-1.9/bin:$PATH' >> /home/vagrant/.profile

mkdir -p /home/vagrant/go/{bin,pkg,src}
mkdir -p /home/vagrant/go/src/qmstr-prototype
ln -s /vagrant/qmstr /home/vagrant/go/src/qmstr-prototype/qmstr
chown -R vagrant:vagrant /home/vagrant/go

echo 'Run `. /vagrant/qmstr-env.sh` after compiling qmstr-compile'
