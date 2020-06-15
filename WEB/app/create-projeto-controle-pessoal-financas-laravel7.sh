#!/usr/bin/env bash

sudo apt install php-sqlite3
sudo apt install php-mbstring
sudo apt install php-xml
composer create-project --prefer-dist laravel/laravel controle-pessoal-financas-laravel 7.*
