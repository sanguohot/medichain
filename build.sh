#!/bin/bash

GO11MODULE=on
go build --ldflags "-linkmode external -extldflags -static" -a -o main -v
