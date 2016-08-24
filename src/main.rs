#[macro_use]
extern crate log;
extern crate env_logger;
extern crate clap;

include!(concat!(env!("OUT_DIR"), "/package.rs"));

use clap::{Arg, App};

fn main() {
    println!("Compiled on {}", ATAXIA_COMPILED);

    // Set up and parse the command-line arguments
    let matches = App::new("Ataxia Engine")
        .version(env!("CARGO_PKG_VERSION"))
        .author("Xenith Studios (see AUTHORS)")
        .about(env!("CARGO_PKG_DESCRIPTION"))
        .arg(Arg::with_name("config")
            .help("The config file to use")
            .short("c")
            .long("config")
            .value_name("FILE")
            .takes_value(true)
            .default_value("data/config.toml"))
        .arg(Arg::with_name("listen_addr")
            .help("Listen address and port")
            .short("l")
            .long("listen")
            .value_name("address:port")
            .takes_value(true)
            .default_value("*:9000"))
        .arg(Arg::with_name("pid_file")
            .help("The filename to write the PID into")
            .short("p")
            .long("pid")
            .value_name("FILE")
            .takes_value(true)
            .default_value("data/ataxia.pid"))
        .arg(Arg::with_name("hotboot")
            .help("Recover by performing a hotboot (you should never specify this manually!)")
            .short("H")
            .long("hotboot"))
        .arg(Arg::with_name("descriptor")
            .help("Descriptior to use when hotbooting")
            .short("D")
            .long("descriptor")
            .takes_value(true)
            .requires("hotboot")
            .default_value("0"))
        .arg(Arg::with_name("debug")
            .help("Enable debugging output")
            .short("d")
            .multiple(true))
        .arg(Arg::with_name("verbose")
            .help("Enable verbose output")
            .short("v")
            .multiple(true))
        .get_matches();

    // Load settings from config file
    if let Some(c) = matches.value_of("config") {
        println!("Value for -c: {}", c);

    }
    // Clean up from previous unclean shutdown if necessary (not hotbooting)
    //   Delete PID file if it exists

    // Set up callbacks for signals
    // Write PID file

    // Initialize
    //   Seed rand
    //   Logging
    env_logger::init().expect("Failed to initialize logging.");
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

    // Recover from hotboot if necessary
    if matches.is_present("hotboot") {
        recover();
    }

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

// When hotboot is called, this function will save game and world state,
// save each player state, and save the player list. Then it will do some
// cleanup (including closing the database) and call Exec to reload the
// running program.
// TODO: This is currently a stub function to lay out future functionality.
// fn hotboot() {
// Save game state
// Save socket and player list
// Cleanup and close database connection
// Exec to reload
// If we got here, something went wrong. Exit with error.
// }


// When recovering from a hotboot, recover will restore the game and world state,
// restore the player list, and restore each player state. Once that is done, it
// will then reconnect each active descriptor to the associated player.
// TODO: This is currently a stub function to lay out future functionality.
fn recover() {
    // Fill out functionality
}
