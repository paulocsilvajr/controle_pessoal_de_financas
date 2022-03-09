FROM golang:1.16-buster

WORKDIR /go/src/controle_pessoal_de_financas
COPY install_dependencies.sh .

RUN ./install_dependencies.sh
RUN apt update && apt install -y musl musl-dev musl-tools tree file
