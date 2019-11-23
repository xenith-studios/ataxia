# Ataxia Game Engine

## About

Ataxia is a modern MUD/MUSH engine written in [Rust](https://www.rust-lang.org/). It utilizes Lua for commands and game logic. It uses separate processes for the game engine and network proxy.

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

The only dependency at this time is the Rust toolchain, which can be installed from the official channels via [rustup](https://www.rust-lang.org/en-US/install.html). The game is written to work with the Rust 2018 edition.

For best results, the minimum required version of Rust is 1.39.0. The code should compile with any stable or nightly version released after 2019-11-08.

For further information: https://www.rust-lang.org/

### Building from Source

Compiling from source is straightforward:

```sh
$ make
```

This will install all library dependencies and build ataxia. (Make will automatically call `cargo build` as needed.)

Modify the configuration files in `data/proxy.toml` and `data/engine.toml`

Run Ataxia:

```sh
$ ./bin/startup.py
```

### Development

If you would also like to develop Ataxia, you will need to install the following additional tools:

- Rust
    - rustfmt
    - clippy

You can install both tools via make:

```sh
$ make bootstrap
```

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
        Rust source code
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
