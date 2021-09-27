#!/usr/bin/env bash

sudo apt install php7.4-sqlite3 php7.4-mbstring php7.4-xml;

composer install;
cp -vi .env.bak .env;
