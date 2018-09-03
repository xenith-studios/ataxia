//! Binary source for the game engine
//! There should be minimal functionality in this code. It exists mainly to call out to the library code.
#![deny(
    trivial_casts,
    trivial_numeric_casts,
    unsafe_code,
    unused_import_braces,
    unused_qualifications
)]

include!(concat!(env!("OUT_DIR"), "/version.rs"));

use std::fs::File;
use std::io::Write;
use std::path::Path;
use std::process;

use clap::{App, Arg};
use log::{error, info, log};
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
        ).arg(
            Arg::with_name("proxy_addr")
                .help("Address and port of the network proxy process")
                .short("a")
                .long("addr")
                .value_name("address:port")
                .takes_value(true),
        ).arg(
            Arg::with_name("pid_file")
                .help("The filesystem path to the PID file")
                .short("p")
                .long("pid")
                .value_name("FILE")
                .takes_value(true),
        ).arg(
            Arg::with_name("debug")
                .help("Enable debugging output")
                .short("d"),
        ).arg(
            Arg::with_name("verbose")
                .help("Enable verbose output")
                .short("v"),
        ).get_matches();

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
            File::create(config.log_file())?,
        ),
    ])?;
    info!("Loading Ataxia Engine, compiled on {}", ATAXIA_COMPILED);

    // TODO: Figure out a system for catching/handling signals (SIGINT, SIGQUIT, SIGHUP)

    // Clean up from previous unclean shutdown if necessary

    // Write PID to file
    // TODO: Acquire lock on PID file as additional method of insuring only a single instance is running?
    let pid_path = config.pid_file().to_string();
    let pid_file = Path::new(&pid_path);
    // FIXME: Remove once we have a startup/supervisor system in place to handle unclean shutdown
    if pid_file.exists() {
        std::fs::remove_file(pid_file)?;
    }
    File::create(pid_file)?.write_all(format!("{}", process::id()).as_ref())?;

    // Initialize support subsystems
    //   Environment
    //   Queues
    //   Database

    // Initialize engine subsystem
    let server = ataxia::Server::new(config).unwrap_or_else(|err| {
        error!("Unable to initialize the engine: {}", err);
        std::process::exit(1);
    });

    // Initialize async networking subsystem in a dedicated thread

    // Start main game loop
    if let Err(e) = server.run() {
        error!("Unresolved engine error: {}", e);
        std::process::exit(1);
    }

    // If the game loop exited without an error, we have a clean shutdown
    // Flush pending database writes and close database connection
    // Remove the PID file
    if pid_file.exists() {
        std::fs::remove_file(pid_file)?;
    }

    Ok(())
}
