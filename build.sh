#!/bin/bash

sh tools/release-edit.sh

echo "Running go fmt..."
for dir in {engine,lua,game,handler,settings,utils}; do
    cd $dir;
    go fmt;
    cd ..;
done
go fmt

echo "Building ataxia..."
go build
mv ataxia bin/

echo "Done. Binary found at bin/ataxia."
