FROM golang:1.18-alpine

WORKDIR /go/src/github.com/paulocsilvajr/

RUN apk add --no-cache tree file git openssl
RUN git clone --branch api https://github.com/paulocsilvajr/controle_pessoal_de_financas.git
RUN cd controle_pessoal_de_financas/ && ./API/v1/install_dependencies.sh

RUN mkdir -p controle_pessoal_de_financas/API/v1/keys
COPY API/v1/keys controle_pessoal_de_financas/API/v1/keys

RUN cd controle_pessoal_de_financas/ && ./API/v1/compile_static_linux_amd64_docker.sh

CMD ["sh"]
