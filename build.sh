#!/bin/sh
go get -u github.com/go-yaml/yaml
mkdir -p artifacts
go build -ldflags "-s -w" -o artifacts/dogg