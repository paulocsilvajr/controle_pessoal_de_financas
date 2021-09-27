#!/usr/bin/env bash

# fonte: https://www.itzgeek.com/how-tos/linux/ubuntu-how-tos/how-to-install-php-7-3-7-2-7-1-on-ubuntu-18-04-ubuntu-16-04.html
sudo apt update
sudo apt install -y curl wget gnupg2 ca-certificates lsb-release apt-transport-https
sudo apt-add-repository ppa:ondrej/php
sudo apt update
sudo apt install -y php7.4 php7.4-cli php7.4-common
