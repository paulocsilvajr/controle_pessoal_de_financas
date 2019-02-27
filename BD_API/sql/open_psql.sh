#!/bin/bash

USER_DB="pi"

if [ "$USER" == "$USER_DB" ]; then
    psql controle_pessoal_de_financas
else
    echo "Altere para o usuário $USER_DB"
fi
