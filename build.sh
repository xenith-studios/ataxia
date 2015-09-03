#!/bin/bash

sh tools/release-edit.sh

echo "Running go fmt, go vet, and golint..."
go fmt ./...
goimports -w .
go vet ./...
golint ./... | egrep -v "_string.go"

echo "Building ataxia..."
go build ./cmd/ataxia
mv ataxia bin/

echo "Done. Binary found at bin/ataxia."
