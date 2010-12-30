#!/bin/bash
#
set -e

DEPS=""
for dep in ${DEPS}; do
    cd $dep ; make clean || true; cd ..
done
make clean
