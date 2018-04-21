#!/bin/bash

is_root(){
    if ! [ $UID -eq 0 ]; then
        echo "Execute como root"
        exit 1
    fi
}

install_mysql-workbench(){
    apt install mysql-workbench -y && echo -e "Instalado MySQL Workbench\n\n"
}

is_root

install_mysql-workbench &&
    echo "END OF LINE."
