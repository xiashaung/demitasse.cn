#!/usr/bin/env sh

cd ./src

go mod download

go build .

mv demitasse.cn /home/www/api/

