Getting Started 
## Setting up Go ##
For MAC 

Download installer from: https://dl.google.com/go/go1.10.3.darwin-amd64.tar.gz

check go installation directory

 /usr/local/go

## GO project structure ##
cd							## change directory to home folder

mkdir -p workspace/code/go  ## create go project directory

mkdir bin                   ## for building executables 

mkdir pkg                   ## for pakages 

mkdir src					## source directory for project

## setting up GOROOT GOBIN & GOPATH ##

Open bash_profile

vim ~/.bash_profile 

## Go workplace setup ## -- change user name acccordingly

export GOROOT="/usr/local/go"

export GOPATH="/Users/shubham/workspace/code/go"

export PATH="/Users/shubham/workspace/code/go/bin:$PATH"


## Build project ##

cd workspace/code/go/src

cd mycode

go install 

Executables will be created at bin 

## Run  ##
mycode



