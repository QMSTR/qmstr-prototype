#!/bin/bash
cd /home/vagrant
git clone https://git.savannah.gnu.org/git/bash.git
git clone https://github.com/dmgerman/ninka.git
(cd ninka/comments; make; sudo make install)
sudo rm /usr/local/man/man1
(cd ninka; perl Makefile.PL; make; sudo make install)
sudo apt install libio-captureoutput-perl

/vagrant/build-bash.sh
