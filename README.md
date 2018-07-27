# Ataxia Game Engine

## About

Ataxia is a modern MUD/MUSH engine written in Rust and Go. It utilizes Lua for commands and game logic. It uses separate processes for the game engine (written in Rust) and network proxy (written in Go).

PLEASE NOTE THAT CURRENTLY THERE IS VERY LITTLE CODE/FEATURES WRITTEN.

### Planned Features

The separate process model allows Ataxia to support the following features:

- Reload the engine process without disconnecting players. Avoids common "copyover/hotboot" hacks.
  - Adds the ability to rollback players to a previous version of code if needed
- The proxy can support various different communication protocols such as:
  - telnet (with or without tls)
  - ssh
  - websockets (to enable an HTML client)
- Allows the network proxy to present a unified front-end to allow connecting to multiple backend game engines through a single connection/port:
  - Live game
  - Test game (for feature/bug testing)
  - Building interface
  - Admin interface
- The network proxy will manage the account and login/permissions system:
  - Allows granting permissions (building interface, test access) on a per-account basis
  - Allows tying multiple characters to a single account/login
- If you really want to, you can run the network proxy on a different server than the game engine

## Install

### Dependencies

First, install Rust. The game engine is written to work with the Rust 2018 edition, and it currently requires the nightly compiler until the edition stabilizes.
See: https://www.rust-lang.org/

Next, install Go. The network proxy is written to work with Go 1.10, but it should work with most 1.x versions.
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

Modify data/ataxia.toml

Run Ataxia:

```sh
$ bin/ataxia-proxy &
$ bin/ataxia-engine
```

### Development

If you would also like to develop Ataxia, you will need to install the following additional tools:

- Go
    - goimports
    - golint
- Rust
    - rustfmt (This is installed via rustup by default)
    - clippy (This currently only works with nightly, so you will have to install the nightly toolchain alongside the stable toolchain with rustup. Install clippy via `cargo +nightly install clippy`)

To perform a full compile, including all lints:

```sh
$ make full
```

To run all tests:

```sh
$ make test
```

## Directory Layout

    src/
        Rust source code for the engine
    cmd/ (binary),internal/ (library)
        Go source code for the proxy
    bin/
        The location of compiled binary files and scripts for running the game
    docs/
        User and developer documentation
    logs/
        Log files
    tools/
        Helper scripts and tools for developers
    data/
        On-disk data files (ex. config files, world files)
    scripts/
        Lua scripts
    scripts/interface
        Helper scripts that set up the data interface between Rust and Lua
    scripts/commands
        All in-game commands

## License

Licensed under either of:

- Apache License, Version 2.0, (LICENSE-APACHE or http://www.apache.org/licenses/LICENSE-2.0)
- MIT license (LICENSE-MIT or http://opensource.org/licenses/MIT)

at your option.

Previous versions of this code were licensed under the BSD three-clause license.

### Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted for inclusion in the work by you, as defined in the Apache-2.0 license, shall be dual licensed as above, without any additional terms or conditions.