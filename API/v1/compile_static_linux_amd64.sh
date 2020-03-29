#!/bin/bash

BASE=$(dirname $0)
PASTA=$BASE/bin/API_CPF
ARQUIVO=$PASTA/API_CPF

sudo apt install musl musl-dev musl-tools tree && echo -e "Instalado musl\n"

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export CXX_FOR_TARGET=musl-gcc
export CC_FOR_TARGET=musl-gcc
export CC=musl-gcc
export CXX=musl-gcc

if [ -d $PASTA ]; then
    rm -r $PASTA && echo -e "Limpado $PASTA\n"
fi

mkdir -p $PASTA 2> /dev/null

echo -e "Compilando para $GOOS:$GOARCH\n"
go build -v -o $ARQUIVO -a -installsuffix cgo -ldflags '-extldflags "-static" -s' $BASE

cp -vr $BASE/keys $PASTA
chmod -R 777 $PASTA

echo
file $ARQUIVO && ldd $ARQUIVO

echo -e "\nArquivos em $PASTA:"
tree -D $PASTA
