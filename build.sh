#!/bin/bash

sh tools/release-edit.sh

echo "Running go fmt, go vet, and golint..."
for dir in {engine,lua,game,handler,utils}; do
    cd $dir;
    go fmt;
    goimports -w .;
    go vet *.go;
    golint;
    cd ..;
done
go fmt
goimports -w .
go vet *.go
golint

echo "Building ataxia..."
go build
mv ataxia bin/

echo "Done. Binary found at bin/ataxia."
