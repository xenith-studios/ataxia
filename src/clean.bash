#!/bin/bash
#
set -e

DEPS="lua settings server game"
for dep in ${DEPS}; do
    cd $dep ; make nuke || true; cd ..
done
make nuke
