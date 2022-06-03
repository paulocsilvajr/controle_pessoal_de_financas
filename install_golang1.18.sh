#!/usr/bin/env bash

USUARIO=$( logname )
HOME_USUARIO=/home/$USUARIO

GOLANG=go1.18.1.linux-amd64
ARQUIVO=${GOLANG}.tar.gz
LINK=https://go.dev/dl/${ARQUIVO}

if [ -d "/usr/local/go" ]; then
    versao=$(/usr/local/go/bin/go version)
    echo "Removendo versão anterior do GO[$versao]"
    rm -rf /usr/local/go
fi

adicionar_go_em_profile() {
    if grep "/usr/local/go/bin" $1 > /dev/null; then
        echo "Arquivo $1 - OK"
    else
        echo 'export PATH=$PATH:/usr/local/go/bin' >> $1 &&
        echo 'export GOPATH=$HOME/go' >> $1 &&
        echo 'export PATH=$PATH:$GOPATH/bin' >> $1 &&
        echo "Atualizado PATH e criado variável de ambiente GOPATH"
    fi
}

wget -c $LINK &&
    tar -C /usr/local/ -xzf $GOLANG.tar.gz &&
    mkdir $HOME_USUARIO/go $HOME_USUARIO/go/bin $HOME_USUARIO/go/pkg $HOME_USUARIO/go/src ||
    chown -R $USUARIO $HOME_USUARIO/go &&
    adicionar_go_em_profile $HOME_USUARIO/.profile &&
    rm -v $ARQUIVO &&
    echo 'Disponibilizado variável $GOPATH, diretório ~/go e finalizado instalação.' &&
    echo -e "Para usuários do ZSH, execute o comando abaixo para reconhecer o executável go:\ncp -v ~/.profile ~/.zprofile"
