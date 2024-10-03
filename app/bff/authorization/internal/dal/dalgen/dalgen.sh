#!/bin/bash

dalgen3 --xml=$1 --db=papercraft --go2=github.com/lingyicute/papercraft-server/app/bff/authorization/internal/dal/dataobject

gofmt -w ../dao/mysql_dao/*.go
gofmt -w ../dataobject/*.go
