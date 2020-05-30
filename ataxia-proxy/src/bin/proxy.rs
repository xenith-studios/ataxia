//! Binary source for the network proxy
//! There should be minimal functionality in this file. It exists mainly to set up the proxy and
//! call out to the library code.
#![warn(
    trivial_casts,
    trivial_numeric_casts,
    unsafe_code,
    unused_import_braces,
    unused_qualifications,
    clippy::all,
    clippy::pedantic,
    clippy::perf,
    clippy::style
)]

// Include this file to get access to the datetime of the last time we compiled
include!(concat!(env!("OUT_DIR"), "/version.rs"));

use std::fs::File;
use std::io::Write;
use std::path::PathBuf;
use std::process;

use log::{error, info};
use simplelog::{CombinedLogger, Config, LevelFilter, TermLogger, TerminalMode, WriteLogger};

#[allow(clippy::too_many_lines)]
fn main() -> Result<(), anyhow::Error> {
    // Load settings from config file while allowing command-line overrides
    let config = ataxia_core::Config::new().unwrap_or_else(|err| {
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
            TerminalMode::Mixed,
        ),
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
    let pid_file = PathBuf::from(config.pid_file());
    // FIXME: Remove once we have a startup/supervisor system in place to handle unclean shutdown
    if pid_file.exists() {
        std::fs::remove_file(&pid_file)?;
    }
    File::create(&pid_file)?.write_all(format!("{}", process::id()).as_ref())?;

    // Initialize support subsystems
    //   Environment
    //   Queues
    //   Database

    // Initialize Tokio async runtime and spin up the worker threadpool
    let rt = tokio::runtime::Runtime::new().expect("Unable to initialize the Tokio Runtime");

    // Initialize proxy and networking subsystems
    let server = ataxia_proxy::Proxy::new(config, rt).unwrap_or_else(|err| {
        error!("Unable to initialize the proxy: {}", err);
        std::process::exit(1);
    });

    // Start main loop
    if let Err(e) = server.run() {
        error!("Unresolved system error: {}", e);
        std::process::exit(1);
    }

    // If the loop exited without an error, we have a clean shutdown
    // Flush pending database writes and close database connection
    // Remove the PID file
    if pid_file.exists() {
        std::fs::remove_file(&pid_file)?;
    }

    info!("Clean shutdown");
    Ok(())
}
