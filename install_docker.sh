#!/usr/bin/env bash

sudo apt-get update &&
    sudo apt-get -y install docker docker-compose &&
    sudo usermod -aG docker $USER && echo -e "Adicionado grupo 'docker' no usuário $USER" &&
    su - $USER
