#/bin/bash

set -v on

if which govendor 2>/dev/null; then
 echo 'govendor exist'
else
 go get -u -v github.com/kardianos/govendor
 echo 'govendor does not exist'
fi


export GOOS=linux GOARCH=amd64
govendor build -o "${PWD##*/}-${GOOS}-${GOARCH}"
export GOOS=linux GOARCH=386 
govendor build -o "${PWD##*/}-${GOOS}-${GOARCH}"
export GOOS=linux GOARCH=arm64 
govendor build -o "${PWD##*/}-${GOOS}-${GOARCH}"
export GOOS=darwin GOARCH=amd64
govendor build -o "${PWD##*/}-${GOOS}-${GOARCH}"
export GOOS=windows GOARCH=amd64
govendor build -o "${PWD##*/}-${GOOS}-${GOARCH}.exe"