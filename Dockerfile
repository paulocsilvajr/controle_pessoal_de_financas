FROM golang:1.18-bullseye

WORKDIR /go/src/controle_pessoal_de_financas

RUN apt update && apt install -y musl musl-dev musl-tools tree file
