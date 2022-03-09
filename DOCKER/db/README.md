# Controle Pessoal de Finanças(CPF)

## Banco de dados(PostgreSQL 11) e PgAdmin 4.15 via DOCKER

### Detalhes de uso

Use o comando *'up.sh'* ou *'run.sh'* para baixar as imagens do docker e criar os containers necessários pra o PostgreSQL 11 e PGAdmin. Se executado do comando *'run.sh'*, será aberto o navegador Google Chrome na página do PGAdmin para a configuração inicial.

Será criado a pasta */home/docker/postgresql-data* e */home/docker/pgadmin* contendo os volumes dos containers do PostgreSQL e PGAdmin para manter as configurações e os bancos criados.

Na primeira conexão ao PGAdmin, deve ser alterado a configuração da conexão do servidor para(Guia 'Connection' de Add New Server):

* Hostname/address: postgresqlpgadmin4_postgres-compose_1
* Port: 5432
* Maintenance database: postgres
* Username: postgres
* Password: Postgres2019!
* Save password? SIM/marcar

Use o comando 'down.sh' para derrubar e remover os containers criados.

Detalhes da criação dos containers e imagens do docker no arquivo *'docker-compose.yml'*.

### Arquivos

```
docker-compose.yml: Arquivo com as configurações para criar os containers do PostgreSQL e PgAdmin na mesma rede do docker via docker-compose
down.sh: Script para derrubar os containers PostgreSQL e PgAdmin
fonte.txt: Arquivo com link de artigo base usado na criação do docker-composer.yml
get_network.sh: Script que mostra a configuração da rede criada via docker-composer.yml
log_pgadmin.sh: Script que exibe o log do container ref ao PgAdmin
log_postgres-f.sh: Script que exibe o log do container ref ao PostgreSQL, mantem exibindo log até receber sinal de finalização(CTRL+C)
log_postgres.sh: Script que exibe o log do container ref ao PostgreSQL
open_pgadmin4_firefox.sh: Abre o PgAdmin em navegador Mozilla Firefox
open_pgadmin4_google-chrome.sh: Abre o PgAdmin em navegador Google Chrome
parse.sh: Script que pega informações de docker-composer.yml e gera arquivo var_yml com informações dos containers
passwords.sh: Script que exibe as senhas contidas em docker-compose.yml
ps.sh: Script para exibir a condição/estado de cada container criado
README.md: Este arquivo com detalhes de branch docker
run.sh: Script para subir os container do PostgreSQL e PgAdmin e abrir o navegador para acessar o PgAdmin
up.sh: Script para subir os container do PostgreSQL e PgAdmin
var_yml: Arquivo gerado a partir de script parse.sh com informações obtidas de docker-compose.yml
```