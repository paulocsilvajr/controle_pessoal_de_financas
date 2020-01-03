# Controle Pessoal de Finanças(CPF)
### API desenvolvida no Ubuntu 18.04 em Golang 1.12

Este repositório contém a API de um software desenvolvido em go para disponibilizar rotas para as operações de inclusão, alteração, exclusão e consulta de lançamentos contábeis em um sistema de controle de finanças. Toda a comunicação com a API é feita por token gerado ao fazer login e mantido por tempo pré definido em arquivo de configuração. A API faz a comunicação e interpretação dos dados das aplicações WEB e Mobile com o banco de dados usando requisições GET/POST/PUT/DELETE usando o formato JSON para representar as informações compartilhadas.

### Pré-requisitos

Para testar e compilar, deve-se ter instalado e configurado o **Go**. O link do [github](https://github.com/golang/go/wiki/Ubuntu) demonstra como instalar via **apt**. Também pode-se instalar usando o utilitário de instalação [**Instalador de programas**](https://github.com/paulocsilvajr/instalador-programas). Para que a compilação ocorra corretamente, este branch(**API**) deve ser clonado dentro da pasta **$HOME/go/src**

Como go é uma linguagem compilada, deve-se gerar o executável usando o script "compile_static_linux_amd64.sh"(executável em bin/) ou executar o script "run.sh" que compila e executa o programa em pasta temporária.

Para executar os scripts(*.sh) deve-se conceder privilégio de execução a cada um.

O banco de dados utilizado é o **PostgreSQL 11**. Atualmente nos testes é usado o **Docker** para criar o container do Postgres e PgAdmin. Pode-se replicar esse ambiente utilizando o repositório [docker-code](https://bitbucket.org/paulocsilvajr/docker-code/src/master/) pasta **yml/postgreSQL_pgadmin4**. Senhas e usuários em arquivo .YML de repositório apresentado. Em branch **DESENVOLVIMENTO** na pasta DOC/sql/, pode-se encontrar backups.sql do banco de dados, use a versão com data mais atual para restaurar o banco **controle_pessoal_financas**.

Execute o script **install_dependencies.sh** antes da primeira execução do código para instalar as dependências da API em GO.

### Arquivos

```
bin/: Pasta com os binários compilados pelo script "compile_static_linux_amd64.sh". Ela é gerado ao executar o script citado anteriormente.
compile_static_linux_amd64.sh: Script que compila a API usando como compilador C o [MUSL](https://www.musl-libc.org/). Faz a instalação do Musl via apt caso não esteja instalado. Foi usado o Musl para gerar executável estaticamente lincado, evitando bibliotecas dinamicamente lincadas em executável de API. Também copia as chaves/vertificados para usar o protocolo HTTPS.
config/: Pasta com o arquivo de configuração inicial da API "config.json". Nele pode-se alterar configurações do banco de dados, duração do token, host e porta padrão e protocolo usado(HTTP ou HTTPS). Também contém informações de rotas da API. Ao executar a API compilada na primeira vez, é gerado na mesma pasta do executável uma cópia definida em arquivo "config.go" na função "criarConfigPadrao()".
controller/: Pasta com todos os handlers/controllers da API. Faz a comunicação entre as rotas definidas e as funções DAO que comunicam com o banco de dados.
cURL/: Pasta com exemplo de script para comunicar com API via comando "curl"
dao/: Pasta com os arquivos que fazem a comunicação com o banco de dados, convertendo os modelos em instruções SQL e vice-versa.
frase_key.txt: Arquivo com a frase secreta e desafio de chaves/certificados gerados para as requisições HTTPS.
gerar_keys.sh: Script para gerar as chaves/certificados utilizados no protocolo HTTPS.
go_clean.sh: Script para limpar o cache de teste. Deve ser executado sempre que o cache atrapalhar os testes unitários não retornando os valores desejados.
go_test_all_json.sh: Script para efetuar todos os testes criados e gerar como saída um JSON com os resultados.
go_test_all.sh: Script para efetuar todos os teste unitários criados em suas respectivas pastas/packages.
go_test_time.sh: Variação do script "go_test_all.sh" que retorna o tempo de execução e informações do sistema após efetuar os testes.
helper/: Pasta com funções auxiliares usadas em várias partes da API.
install_dependencies.sh: Script para instalar as dependências da API em Go.
keys/: Pasta com as chaves/certificados autoassinados da requisições HTTPS.
logger/: Pasta com as funções referentes ao gerador de log do sistema, tanto em arquivo, quanto em tela.
logs/: Pasta com os logs do sistema. Ao compilar e executar a API na primeira vez, será criado uma pasta log no diretório da API compilada.
main.go: Arquivo principal da API que inicia todos os pacotes necessários para a execução da mesma.
model/: Pasta com a definição dos modelos da API
postman/: Pasta com exemplo de requisições usando o programa [Postman](https://www.getpostman.com/).
push_api.sh/: Script para fazer o push para os repositórios remotos na branch api.
README.md: Este arquivo de ajuda e apresentação da API.
route/: Pasta com as rotas e funções associadas a rotas da API.
run_ngrok.sh: Script para executar a API a partir de endereço remoto disponibilizado pelo serviço gratuíto do [ngrok](https://ngrok.com/). Deve ser somente executado após executar o script run.sh. Este script somente disponibiliza um endereço externo público para a API rodando na sua máquina.
run.sh: Script que executa e roda a API, disponibiliza uma rota local(localhost ou IP máquina) usando a porta informada em arquivo config/config.json, por padrão a rota:porta https://localhost:8085. Para alterar o host e a porta, deve-se executar o arquivo compilado da API na primeira vez para gerar o arquivo de configuração config.json e depois alterar os campos "host" e "porta". Após alterar o arquivo de configuração, reexecute o arquivo compilado da API que usará as configurações definidas no arquivo de configuração. Para a API funcionar, o banco de dados deve estar rodando de acordo com as configurações de config.json.

```

### Licença

[Licença MIT](https://github.com/paulocsilvajr/controle_pessoal_de_financas/blob/api/license_mit.txt), arquivo em anexo no repositório.

### Contato

Paulo Carvalho da Silva Junior - pauluscave@gmail.com
