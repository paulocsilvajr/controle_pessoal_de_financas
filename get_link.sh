#!/bin/bash

echo $(cat links.txt | grep $1 | awk -F '::' '{print $NF}')
