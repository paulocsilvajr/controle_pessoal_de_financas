#!/usr/bin/env bash

# ALTERE a variável a seguir com o diretório do projeto
# controle_pessoal_de_financas de acordo com a localização no seu computador.
# Geralmente localizado na pasta 'go' na home se seu usuário. Exemplo:
# /home/seu-usuario/go/src/github.com/usuario_github/controle_pessoal_de_financas
DIR_ABS=$HOME/go/src/github.com/paulocsilvajr/controle_pessoal_de_financas

docker run -it --rm -v $DIR_ABS:/go/src/controle_pessoal_de_financas golang_custom:controle_pessoal_de_financas ./API/v1/compile_static_linux_amd64.sh

sudo chown -vR $USER $DIR_ABS/API/v1/bin
