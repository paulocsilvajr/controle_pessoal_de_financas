#!/bin/bash

docker exec -it postgresql_pgadmin4_postgres-compose_1 createdb -U postgres --encoding=UTF8 controle_pessoal_financas
