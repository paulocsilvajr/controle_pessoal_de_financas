#!/usr/bin/env bash

DIR=$(dirname $0)
DIR_ABS=$(readlink -e $DIR)

echo $DIR_ABS

function compilar() {
    echo "Compilando no DOCKER ..."
    docker run -it --rm -v $DIR_ABS:/go/src/github.com/paulocsilvajr/controle_pessoal_de_financas golang_custom:controle_pessoal_de_financas sh -c "./API/v1/install_dependencies.sh && ./API/v1/compile_static_linux_amd64_docker.sh"
}

function corrigir_permissao() {
    echo "Alterar DONO de pasta 'bin' e seu conteúdo para '$USER'"
    sudo chown -vR $USER $DIR_ABS/API/v1/bin
}

compilar &&
    corrigir_permissao &&
    echo "END OF LINE."

