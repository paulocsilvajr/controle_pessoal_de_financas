# Controle Pessoal de Finanças(CPF)
### API desenvolvida em Golang 1.12

Este repositório contém a API de um software desenvolvido em go para disponibilizar rotas para as operações de inclusão, alteração, exclusão e consulta de lançamentos contábeis em um sistema de controle de finanças. Toda a comunicação com a API é feita por token gerado ao fazer login e mantido por tempo pré definido em arquivo de configuração. A API faz a comunicação e interpretação dos dados das aplicações WEB e Mobile com o banco de dados usando requisições GET/POST/PUT/DELETE usando o formato JSON para representar as informações compartilhadas.

### Pré-requisitos

Para testar e compilar, deve-se ter instalado e configurado o **Go**. O link do [github](https://github.com/golang/go/wiki/Ubuntu) demonstra como instalar via **apt**. Também pode-se instalar usando o utilitário de instalação [**Instalador de programas**](https://github.com/paulocsilvajr/instalador-programas). Para que a compilação ocorra corretamente, este branch(**API**) deve ser clonado dentro da pasta **$HOME/go/src**

Como go é uma linguagem compilada, deve-se gerar o executável usando o script "compile_static_linux_amd64.sh"(executável em bin/) ou executar o script "run.sh" que compila e executa o programa em pasta temporária.

Para executar os scripts(*.sh) deve-se conceder privilégio de execução a cada um.

O banco de dados utilizado é o **PostgreSQL 11**. Atualmente nos testes é usado o **Docker** para criar o container do Postgres e PgAdmin. Pode-se replicar esse ambiente utilizando o repositório [docker-code](https://paulocsilvajr@bitbucket.org/paulocsilvajr/docker-code.git). Senhas e usuários em arquivo .YML de repositório apresentado. Em branch **DESENVOLVIMENTO** na pasta DOC/sql/, pode-se encontrar backups.sql do banco de dados, use a versão com data mais atual para restaurar o banco **controle_pessoal_financas**.

### Arquivos

```
arquivo.sh: Descrição do arquivo.
```

### Licença

[Licença MIT](https://github.com/paulocsilvajr/instalador-programas/blob/master/license_gpl.txt), arquivo em anexo no repositório.

### Contato

Paulo Carvalho da Silva Junior - pauluscave@gmail.com
