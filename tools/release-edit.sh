#!/usr/bin/env bash

RELFILE="src/cmd/ataxia/package.go"

if [[ ! -d ".git" ]]; then
  echo "Must be run at the root of the git repository (.git directory not found)"
fi

DATE=`date`
RELEASE=`git tag | egrep "^v([0-9]+\.?)+" | tail -n 1 | cut -d' ' -f 2`

if [[ -z "$RELEASE" ]]; then
  RELEASE="development release"
fi

echo "Updating constants:"
echo "  ATAXIA_VERSION  = '$RELEASE'"
echo "  ATAXIA_COMPILED = '$DATE'"

cat >"$RELFILE" <<EOF
package main

const (
	ATAXIA_VERSION  = "$RELEASE"
	ATAXIA_COMPILED = "$DATE"
)
EOF