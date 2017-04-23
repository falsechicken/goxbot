#!/bin/bash

# SHOULD FETCH AND COMPILE GOXBOT ON UBUNTU BASED MACHINES
# PROVIDE A PATH EX: ./UBUNTU_BUILD.SH /some/path


export GOPATH="$1"

sudo apt-get install golang git -y

go get -u github.com/falsechicken/goxbot
cd "$GOPATH/src/github.com/falsechicken/goxbot/"
go get ./...
go install github.com/falsechicken/goxbot/.

echo "GoXBot binary installed to: $GOPATH/bin"

