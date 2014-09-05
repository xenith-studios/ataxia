#!/bin/bash
# Run the documentation server.

if [[ ! -d ".git" ]]; then
    echo "This command must be run at the root of the git repository."
    exit 1
fi

echo "Open this page in your web browser: http://localhost:6060/pkg/github.com/xenith-studios/ataxia/"
godoc -goroot="`pwd`" -http=:6060
