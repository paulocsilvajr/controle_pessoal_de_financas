FROM golang:1.18-alpine

WORKDIR /go/src/github.com/paulocsilvajr/controle_pessoal_de_financas

RUN apk add --no-cache tree file
