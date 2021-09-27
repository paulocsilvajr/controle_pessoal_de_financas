#!/usr/bin/env bash

SCRIPT=$1

docker container run -it --rm -v $(pwd)/app:/app php:7.4.3-cli-buster php $SCRIPT
