# Ataxia Game Engine #

## About ##

Ataxia is a modern MUD engine written in Rust. It utilizes asynchronous I/O via Futures to offload network
tasks to separate threads and uses Lua for commands and game logic.

## Install ##

First, install Rust. Ataxia is written to work with Rust 1.22, but it should work with most 1.x versions.
See: https://www.rust-lang.org/en-US/

The network proxy is written in Go. It is written to work with Go 1.9, but it should work with most 1.x versions.
See: https://golang.org/

You will also want to install Dep, Go's upcoming dependency management tool.
See: https://github.com/golang/dep

Once Rust and Go are installed:

    $ make

This will install all dependencies and build ataxia. (Make will automatically call cargo build and go build.)

Modify data/ataxia.toml and data/proxy.toml

Run Ataxia:

    $ bin/ataxia-proxy &
    $ cargo run --bin ataxia-engine

## Directory Layout ##

    src/
        All source code
    bin/
        The location of compiled binary files and helper scripts
    doc/
        User and developer documentation
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
        Helper scripts that set up the data interface between Rust and Lua
    scripts/command
        All in-game commands

## License ##

`ataxia` is primarily distributed under the terms of both the MIT License and
the Apache License (Version 2.0), with portions covered by various BSD-like
licenses. Previous versions of this code were licensed under the BSD three-clause license.

See LICENSE-APACHE and LICENSE-MIT for details.
