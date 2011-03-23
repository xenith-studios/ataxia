#!/bin/bash

echo "Ataxia Authors:" > AUTHORS
git log --pretty=format:"%an <%ae>" | sort -u >> AUTHORS
