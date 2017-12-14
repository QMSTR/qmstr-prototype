# qmstr-prototype
The Quartermaster prototype.

Quartermaster is a suite of command line tools and build system plugins that instruments software builds to create
FLOSS compliance documentation and support compliance decision making. It executes as part of a software build process
to generate reports about the analysed product.

### Test Quartermaster

In our first demo, Quartermaster builds bash, retrieves license
information using Ninka, and uses the identified licenses to generate
a simple SPDX file.

You can test Quartermaster and see how it works
by following the steps bellow:

1. Check out a clone of this repo to a location of your choice.
2. For historical reasons we are using the virtual machine tool
   Vagrant, so you need to type in your command line: `vagrant up
   --provider virtualbox` , to bring up the test VM.
3. Then run `vagrant ssh` to enter the VM
3. Change your location to /vagrant by typing `cd /vagrant`.
4. Start the demo by running our demo setup `./vagrant-demo-setup.sh`.

You will see a lot of things going on in your terminal. What happens is that Quartermaster hooks into the existing
build phase and takes information about the target and the sources that were used to build bash. In the end of the analysis you will
see an SPDX output like this:

`SPDXVersion: SPDX-2.0`

`DataLicense: CCO-1.0`

`PackageName:  bash`

`PackageLicenseDeclared: NONE AND GPLv3+ AND GPLv3+ BisonException`

* SPDXVersion: The version of the spec used, normally "SPDX-2.0".
* DataLicense: The license for the license data itself, normally this
  is "CC0-1.0" (note that this is not the license for the software or
  data being packaged).
* PackageName: The full name of the package.
* PackageLicenseDeclared: The licenses identified in the source files.

The prototype does not yet understand all of the source file licenses
in bash. Some of the files compiled are generated during the
build. Support for code generators will be added to Quartermaster in
one of the upcoming sprints.

