#!/bin/sh
cd /home/vagrant
git clone https://git.savannah.gnu.org/git/bash.git
git clone https://github.com/dmgerman/ninka.git
(cd ninka/comments; make; sudo make install)
cpan App::cpanminus
export PATH=/home/vagrant/perl5/bin:$PATH
sudo cpanm IO::CaptureOutput Spreadsheet::WriteExcel Test::Pod Test::Strict DBD::SQLite DBI

# configure bash:
# ...
# configure PATH for QMSTR
# ...
# run the build
# ...

