#!/bin/bash

source /etc/profile

cd /root/mini_proj

git remote -v | grep fetch

git pull

make

systemctl restart miniproj

sleep 3

netstat -ntlp | grep livingserver

exit $?

