#!/usr/bin/env bash

curl -L https://get.daocloud.io/docker/compose/releases/download/v2.11.2/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
docker-compose version
