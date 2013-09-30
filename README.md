# Ataxia Game Engine #

## About ##

Ataxia is a modern MUD engine written Go. It heavily utilizes concurrency and uses Lua for configuration and game logic.

## Install ##

First, install Go. Ataxia is written to work with the current Go1 release. See: http://golang.org/doc/install

Once Go is installed:

    $ make

This will install all dependencies and build ataxia, putting the binary in bin/

Modify etc/config.lua

Run Ataxia:

    $ ./bin/ataxia
