# SemverUtils

Utility for Semantic Version control. Its purpose to is read semantic release styled commit messages and current tags to calculate the next version number to use when releasing a package of a Git controlled project.

This is an initial release. Command line options are likely to change. 


# Download

```
git clone https://github.com/marshyon/semverUtils.git
```

# Build

The following is a typical installation for a Linux OS:

```
cd semverUtils/cmd/version
go build -o semverutils
sudo cp semverutils /usr/local/bin
```
# Prerequisites

Semverutils uses the command line Git executable so this will need to be installed prior to using ths application.

For Linux operating systems such as Debian / Ubuntu this is typically already installed or will need to be installed with:

```
sudo apt install git
```

Windows installations will need to have git bash installed from :

https://gitforwindows.org/

Use the `Download` link and follow instructions to install for Windows

# Usage

Change to the directory of a Git controlled project

run

```
semverutils
```

which outputs ( to stdout ) the next version to use when publishing a new release. 

# Tests

The following from the root of the application directory will run unit tests:

```
go test -v ./...
```
The following from the root directory of the application directory will run BDD tests:

```
godog
```

# RELEASES

testing ....
feature ....