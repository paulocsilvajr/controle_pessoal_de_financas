#!/usr/bin/env bash

# fontes:
# https://hub.docker.com/_/php
# https://stackoverflow.com/questions/17157721/how-to-get-a-docker-containers-ip-address-from-the-host

MYIMAGE=my-php-app
MYCONTAINER=my-running-php-app
COR="\033[0;34m%s\033[0m"

if [ "$1" == "-r" ]; then
    docker stop $MYCONTAINER &&
        docker rm $MYCONTAINER && printf "$COR\n" "Removido container '$MYCONTAINER'"
else
    docker stop $MYCONTAINER &&
        docker rm $MYCONTAINER &&
        printf "$COR\n" "PARADO e REMOVIDO '$MYCONTAINER' que estava rodando em docker"

    docker rmi $MYIMAGE --force
    docker build -t $MYIMAGE . &&
        printf "$COR\n" "Criado imagem: '$MYIMAGE' com arquivos de pasta src/"

    printf "\n$COR" "Criado conteiner: "
    docker run -d --name $MYCONTAINER $MYIMAGE

    printf "$COR" "IP: "
    echo "http://$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $MYCONTAINER)"
fi
