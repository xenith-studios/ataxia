# Ataxia Game Engine

## About

Ataxia is a modern MUD engine written in Rust and Go. It utilizes Lua for commands and game logic. It uses separate processes for the game engine (written in Rust) and network proxy (written in Go).

### Features

The separate process model allows Ataxia to support the following features:

- Reload the engine process without disconnecting players. Avoids common "copyover/hotboot" hacks.
- The proxy can support various different communication technologies such as:
  - telnet
  - ssh
  - websockets (to enable an HTML client)
- Allows the network proxy to present a unified front-end to allow connecting to multiple backend game engines:
  - Live game
  - Building server
  - Test game for live feature testing

## Install

### Dependencies

First, install Rust. The game engine is written to work with Rust 1.22, but it should work with most 1.x versions.
See: https://www.rust-lang.org/

Next, install Go. The network proxy is written to work with Go 1.9, but it should work with most 1.x versions.
See: https://golang.org/

You will also want to install Dep, Go's upcoming dependency management tool.
See: https://github.com/golang/dep

### Building from Source

You will need to clone this repository into your `$GOPATH` to allow Go to build properly. I'll admit, the Go tooling is a little odd, so you'll want to do something similar to:

```sh
$ git clone https://github.com/xenith-studios/ataxia/ $GOPATH/src/github.com/xenith-studios/
```

After that, you can build the game:

```sh
$ make bootstrap
$ make
```

This will install all dependencies and build ataxia. (Make will automatically call cargo build and go build.)

Modify data/engine.toml and data/proxy.toml

Run Ataxia:

```sh
$ bin/ataxia-proxy &
$ bin/ataxia-engine
```

# Contributing

If you would also like to develop Ataxia, you will need to install the following additional tools:

- Go
    - goimports
    - golint
- Rust
    - rustfmt (This will eventually be installed as part of cargo by rustup, but it is currently in heavy flux)
    - clippy (It currently only works with nightly, so you will have to install the nightly toolchain alongside the stable toolchain with rustup)

To perform a full compile, including all lints:

```sh
$ make full
```

To run all tests:

```sh
$ make test
```

# Directory Layout

    src/
        All Rust source code for the engine
    cmd/
        Go binary code for the proxy
    internal/
        Go internal library code for the proxy
    bin/
        The location of compiled binary files and scripts for running the engine
    docs/
        User and developer documentation
    logs/
        Location of stored log files
    tools/
        Helper scripts and tools for developers
    data/
        On-disk data files, such as config files and world files
    scripts/
        On-disk storage location for all Lua scripts
    scripts/interface
        Helper scripts that set up the data interface between Rust and Lua
    scripts/command
        All in-game commands

# License

`ataxia` is primarily distributed under the terms of both the MIT License and
the Apache License (Version 2.0), with portions covered by various BSD-like
licenses. Previous versions of this code were licensed under the BSD three-clause license.

See LICENSE-APACHE and LICENSE-MIT for details.
