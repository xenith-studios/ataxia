extern crate clap;
#[macro_use]
extern crate log;
extern crate simplelog;

extern crate ataxia;

include!(concat!(env!("OUT_DIR"), "/version.rs"));

use std::fs::File;
use std::io::Write;
use std::path::Path;
use std::process;

use clap::{App, Arg};
use simplelog::*;

fn main() {
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

    let debug = match matches.occurrences_of("debug") {
        0 => false,
        1 | _ => true,
    };

    let verbose = match matches.occurrences_of("verbose") {
        0 => false,
        1 | _ => true,
    };

    // Load settings from config file while allowing command-line overrides
    let config_path = Path::new(
        matches
            .value_of("config")
            .expect("Unable to specify config file path"),
    );

    let mut config =
        ataxia::config::Config::read_config(config_path).expect("Unable to load the configuration");

    if let Some(pid_file) = matches.value_of("pid_file") {
        config.set_pid_file(pid_file);
    }

    if let Some(proxy_addr) = matches.value_of("proxy_addr") {
        config.set_proxy_addr(proxy_addr);
    }

    let config = config;

    // Initialize logging subsystem
    CombinedLogger::init(vec![
        TermLogger::new(
            if debug {
                LevelFilter::Debug
            } else if verbose {
                LevelFilter::Info
            } else {
                LevelFilter::Warn
            },
            Config::default(),
        ).expect("Failed to intitialize terminal logging"),
        WriteLogger::new(
            if debug {
                LevelFilter::Debug
            } else {
                LevelFilter::Info
            },
            Config::default(),
            File::create(config.get_log_file()).expect("Failed to create logfile"),
        ),
    ]).expect("Failed to initialize logging");
    info!("Loading Ataxia Engine, compiled on {}", ATAXIA_COMPILED);

    // Clean up from previous unclean shutdown if necessary
    // TODO: Should this be handled by the startup/supervisor script?
    let pid_file = Path::new(config.get_pid_file());
    if pid_file.exists() {
        std::fs::remove_file(pid_file).expect("Couldn't remove stale PID file");
    }

    // Write PID to file
    File::create(config.get_pid_file())
        .expect("Couldn't create PID file")
        .write_all(format!("{}", process::id()).as_ref())
        .expect("Couldn't write PID to file");

    // TODO: Set up callbacks for catching signals

    // Initialize
    //   Seed rand
    //   Environment
    //   Queues
    //   Database
    //   Lua

    // Initialize engine
    // Load initial game state
    //   Load database
    //   Load commands
    //   Load scripts
    //   Load world
    //   Load entities

    // Initialize networking event loop in dedicated thread
    // Spawn other threads?
    // Start main game loop

    // Shutdown is caught here?
    // Clean up
    //   Save the world
    //   Shutdown Lua
    //   Flush pending database writes
    //   Close database connection

    // TODO: Should this be handled by the startup/supervisor script? Or should the engine do it to signal a clean shutdown?
    if pid_file.exists() {
        std::fs::remove_file(pid_file).expect("Couldn't remove PID file");
    }
}
