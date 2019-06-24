#!/bin/bash

PASTA=bin/API_CPF
ARQUIVO=$PASTA/API_CPF

sudo apt install musl musl-dev musl-tools && echo -e "Instalado musl\n"

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
go build -v -o $ARQUIVO -a -ldflags '-extldflags "-static" -s'

cp -vr keys $PASTA

# # Caso não compile estáticamente, rode o comando abaixo no terminal e depois reexecute este script
# go build -v -a -ldflags '-extldflags "-static" -s'
