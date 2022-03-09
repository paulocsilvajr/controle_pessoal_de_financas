#!/bin/bash

if [ -e $1 ]; then
    sudo cp -vi $1 /home/docker/postgresql-data/ && \
            docker exec -it postgresqlpgadmin4_postgres-compose_1 psql -U postgres -d controle_pessoal_financas -f /var/lib/postgresql/data/$1
else
    echo 'Informe o arquivo de backup'
fi
