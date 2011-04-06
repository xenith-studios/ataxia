#!/bin/bash
# Run the documentation server.

if [[ ! -d ".git" ]]; then
    echo "This command must be run at the root of the git repository."
    exit 1
fi

godoc -path="." -http=:6060
