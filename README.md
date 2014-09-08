# Ataxia Game Engine #

## About ##

Ataxia is a modern MUD engine written Go. It heavily utilizes concurrency and uses Lua for commands
and game logic.

## Install ##

First, install Go. Ataxia is written to work with the current Go1 release. See: http://golang.org/doc/install

Second, you should install gpm and gvp:

    https://github.com/pote/gpm
    https://github.com/pote/gvp

Once everything is installed:

    $ gpm install
    $ ./build.sh

This will install all dependencies and build ataxia, putting the binary in bin/

Modify data/config.lua

Run Ataxia:

    $ ./bin/ataxia

## Directory Layout ##

    bin/
        The location of compiled binary files and helper scripts
    doc/
        User and developer documentation (not godoc documentation)
    log/
        Location of stored log files
    tools/
        Helper scripts and tools for developers
    data/
        On-disk data files, such as config files
    data/world
        World data files
    scripts/
        On-disk storage location for all Lua scripts
    scripts/interface
        Helper scripts that set up the data interface between Go and Lua
    scripts/command
        All in-game commands
