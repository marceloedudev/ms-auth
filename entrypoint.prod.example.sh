#!/bin/bash

curl -sSL https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar zxf - -C .

./migrate.linux-amd64 -database postgres://user:pass@host:port/godb?sslmode=disable -path migrations up
