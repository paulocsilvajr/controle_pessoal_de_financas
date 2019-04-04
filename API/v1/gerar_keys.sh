#!/usr/bin/env bash

DIAS=365
PASTA=keys
CSR=$PASTA/new.ssl.csr
PRIVATE_KEY=$PASTA/privkey.pem
CERTIFICATE_KEY=$PASTA/new.cert.key
CERTIFICATE=$PASTA/new.cert.cert

openssl genrsa -out $PRIVATE_KEY -des3 2048 && \
    openssl req -new -key $PRIVATE_KEY -out $CSR && \
    openssl rsa -in $PRIVATE_KEY -out $CERTIFICATE_KEY && \
    openssl x509 -in $CSR -out $CERTIFICATE -req -signkey $CERTIFICATE_KEY -days $DIAS
