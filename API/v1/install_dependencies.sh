#!/bin/bash

install(){
    echo "$1:"
    go get -v $1
}

echo -e "Instalando dependências\n"

install github.com/gorilla/mux
install github.com/lib/pq
install github.com/auth0/go-jwt-middleware
