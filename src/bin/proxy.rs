//! Binary source for the network proxy
//! There should be minimal functionality in this file. It exists mainly to set up the proxy and
//! call out to the library code.
#![deny(
    trivial_casts,
    trivial_numeric_casts,
    unsafe_code,
    unused_import_braces,
    unused_qualifications,
    clippy::all,
    clippy::pedantic
)]

include!(concat!(env!("OUT_DIR"), "/version.rs"));

use std::fs::File;
use std::io::Write;
use std::path::Path;
use std::process;

use clap::{App, Arg};
use log::{error, info};
use simplelog::*;

fn main() -> Result<(), failure::Error> {
    // Set up and parse the command-line arguments
    let matches = App::new("Ataxia Network Proxy")
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
                .default_value("data/proxy.toml"),
        )
        .arg(
            Arg::with_name("http_addr")
                .help("Address and port for http connections")
                .short("H")
                .long("http_addr")
                .value_name("address:port")
                .takes_value(true),
        )
        .arg(
            Arg::with_name("telnet_addr")
                .help("Address and port for telnet connections")
                .short("T")
                .long("telnet_addr")
                .value_name("address:port")
                .takes_value(true),
        )
        .arg(
            Arg::with_name("internal_addr")
                .help("Address and port for internal connections")
                .short("I")
                .long("internal_addr")
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
        )
        .expect("Failed to initialize terminal logging"), // FIXME: Remove expect once ? is supported for Option in failure
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
    info!(
        "Loading Ataxia Network Proxy, compiled on {}",
        ATAXIA_COMPILED
    );

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

    // Initialize proxy subsystem
    let server = ataxia::Proxy::new(config).unwrap_or_else(|err| {
        error!("Unable to initialize the proxy: {}", err);
        std::process::exit(1);
    });

    // Initialize async networking subsystem in a dedicated thread

    // Start main loop
    if let Err(e) = server.run() {
        error!("Unresolved system error: {}", e);
        std::process::exit(1);
    }

    // If the loop exited without an error, we have a clean shutdown
    // Flush pending database writes and close database connection
    // Remove the PID file
    if pid_file.exists() {
        std::fs::remove_file(pid_file)?;
    }

    Ok(())
}
