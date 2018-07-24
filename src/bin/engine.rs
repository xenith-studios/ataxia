//! Binary source for the game engine
//! There should be minimal functionality in this code. It exists mainly to call out to the library code.
#![deny(
    trivial_casts, trivial_numeric_casts, unsafe_code, unstable_features, unused_import_braces,
    unused_qualifications
)]
extern crate clap;
#[macro_use]
extern crate log;
extern crate failure;
extern crate simplelog;

extern crate ataxia;

include!(concat!(env!("OUT_DIR"), "/version.rs"));

use std::fs::File;
use std::io::Write;
use std::path::Path;
use std::process;

use clap::{App, Arg};
use simplelog::*;

fn main() -> Result<(), failure::Error> {
    // Set up and parse the command-line arguments
    let matches = App::new("Ataxia Engine")
        .version(env!("CARGO_PKG_VERSION"))
        .author("Xenith Studios (see AUTHORS)")
        .about(env!("CARGO_PKG_DESCRIPTION"))
        .arg(
            Arg::with_name("config")
                .help("The filesystem path to the config file")
                .short("c")
                .long("config")
                .value_name("FILE")
                .takes_value(true)
                .default_value("data/ataxia.toml"),
        )
        .arg(
            Arg::with_name("proxy_addr")
                .help("Address and port of the network proxy process")
                .short("a")
                .long("addr")
                .value_name("address:port")
                .takes_value(true),
        )
        .arg(
            Arg::with_name("pid_file")
                .help("The filesystem path to the PID file")
                .short("p")
                .long("pid")
                .value_name("FILE")
                .takes_value(true),
        )
        .arg(
            Arg::with_name("debug")
                .help("Enable debugging output")
                .short("d"),
        )
        .arg(
            Arg::with_name("verbose")
                .help("Enable verbose output")
                .short("v"),
        )
        .get_matches();

    // Load settings from config file while allowing command-line overrides
    let config = ataxia::Config::new(&matches).unwrap_or_else(|err| {
        eprintln!("Unable to load the configuration file: {}", err);
        std::process::exit(1);
    });

    // Initialize logging subsystem
    CombinedLogger::init(vec![
        TermLogger::new(
            if config.debug() {
                LevelFilter::Debug
            } else if config.verbose() {
                LevelFilter::Info
            } else {
                LevelFilter::Warn
            },
            Config::default(),
        ).expect("Failed to initialize terminal logging"),
        WriteLogger::new(
            if config.debug() {
                LevelFilter::Debug
            } else {
                LevelFilter::Info
            },
            Config::default(),
            File::create(config.get_log_file())?,
        ),
    ])?;
    info!("Loading Ataxia Engine, compiled on {}", ATAXIA_COMPILED);

    // Clean up from previous unclean shutdown if necessary
    // TODO: Should this be handled by the startup/supervisor script?
    let pid_file = Path::new(config.get_pid_file());
    if pid_file.exists() {
        std::fs::remove_file(pid_file)?;
    }

    // Write PID to file
    File::create(config.get_pid_file())?.write_all(format!("{}", process::id()).as_ref())?;

    // TODO: Figure out a system for catching/handling signals (SIGINT, SIGQUIT, SIGHUP)

    // Initialize support subsystems
    //   Seed rand
    //   Environment
    //   Queues
    //   Database
    //   Lua

    // Initialize engine subsystem
    if let Err(e) = ataxia::init() {
        eprintln!("Unresolved engine error during setup: {}", e);
        std::process::exit(1);
    }

    // Initialize async networking subsystem in a dedicated thread?

    // Start main game loop
    if let Err(e) = ataxia::run() {
        eprintln!("Unresolved engine error: {}", e);
        std::process::exit(1);
    }

    // If the game loop exited without an error, we have a clean shutdown
    // TODO: Should this be handled by the startup/supervisor script? Or should the engine do it to signal a clean shutdown?
    if pid_file.exists() {
        std::fs::remove_file(pid_file)?;
    }

    Ok(())
}
