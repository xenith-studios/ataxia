#!/bin/bash
#
set -e

DEPS="settings"
for dep in ${DEPS}; do
    cd $dep ; make nuke || true; cd ..
done
make nuke
