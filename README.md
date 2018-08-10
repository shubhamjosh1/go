Golang BiolerPlate for getting Started

## Setting up Go ##

For MAC 
Download installer from: https://dl.google.com/go/go1.10.3.darwin-amd64.tar.gz

check go installation directory : /usr/local/go


## setting up GOROOT GOBIN & GOPATH ##
Open bash_profile
vim ~/.bash_profile 

## Go workplace setup ##
export GOROOT="/usr/local/go"
export GOPATH="/Users/shubham/workspace/code/go"
export PATH="/Users/shubham/workspace/code/go/bin:$PATH"


## GO project structure ##
cd 
mkdir -p workspace/code/go
mkdir bin
mkdir pkg
mkdir src


## Build project ##
cd workspace/code/go/src
cd mycode
go install 

Executables will be created at bin 

## Run  ##
mycode



