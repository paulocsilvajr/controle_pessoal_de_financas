#!/usr/bin/env bash

ORIGEM=`dirname $0`

sudo apt install lnav -qqq

lnav $ORIGEM/storage/logs/laravel.log
