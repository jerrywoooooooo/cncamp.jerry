#!/bin/bash
git clone https://gitee.com/jerrywoooooooo/cncamp.jerry.git

cd cncamp.jerry/http-server

go build -mod=mod .

pwd && ll
