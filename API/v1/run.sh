#!/usr/bin/env bash

function gorun() {
    # Usado função invocando o executável do Golang no diretório origem para garantir que rodará como root
    /usr/local/go/bin/go run $1
}

clear ; gorun main.go
