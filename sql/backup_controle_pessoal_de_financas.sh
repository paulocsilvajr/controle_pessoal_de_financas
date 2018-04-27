#!/bin/bash

BANCO='controle_pessoal_de_financas'
AGORA=$(date +%Y%m%d)
BACKUP=$AGORA"_"$BANCO".sql"

pg_dump $BANCO > $BACKUP && echo "Criado backup $BACKUP" || echo "Problema na criação do backup"
