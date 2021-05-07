#!/bin/bash

install(){
    echo "$1:"
    go get -v $1
}

echo -e "Instalando dependências\n"

install github.com/gorilla/mux
install github.com/lib/pq
install github.com/auth0/go-jwt-middleware
install github.com/jedib0t/go-pretty/table

# GORM(https://gorm.io/docs/index.html)
install gorm.io/gorm
install gorm.io/driver/postgres
