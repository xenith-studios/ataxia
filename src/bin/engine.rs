extern crate clap;
#[macro_use]
extern crate log;
extern crate simplelog;

extern crate ataxia;

include!(concat!(env!("OUT_DIR"), "/version.rs"));

use std::path::Path;
use std::fs::File;

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
                .help("The config file to use")
                .short("c")
                .long("config")
                .value_name("FILE")
                .takes_value(true)
                .default_value("data/ataxia.toml"),
        )
        .arg(
            Arg::with_name("proxy_addr")
                .help("Listen address and port of the proxy")
                .short("l")
                .long("listen")
                .value_name("address:port")
                .takes_value(true),
        )
        .arg(
            Arg::with_name("pid_file")
                .help("The filename to write the PID into")
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

    // Load settings from config file
    let config_path = Path::new(
        matches
            .value_of("config")
            .expect("Unable to specify config file path."),
    );
    let config = ataxia::config::Config::read_config(config_path)
        .expect("Unable to load the configuration.");

    // Initialize logging subsystem
    CombinedLogger::init(vec![
        TermLogger::new(
            if debug {
                LogLevelFilter::Debug
            } else if verbose {
                LogLevelFilter::Info
            } else {
                LogLevelFilter::Warn
            },
            Config::default(),
        ).expect("Failed to intitialize terminal logging"),
        WriteLogger::new(
            if debug {
                LogLevelFilter::Debug
            } else {
                LogLevelFilter::Info
            },
            Config::default(),
            File::create(config.get_log_file()).expect("Failed to create logfile"),
        ),
    ]).expect("Failed to initialize logging!");
    info!("Loading Ataxia Engine, compiled on {}", ATAXIA_COMPILED);

    // Clean up from previous unclean shutdown if necessary

    // Set up callbacks for signals
    // Write PID file

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
    //   Remove PID file
}
