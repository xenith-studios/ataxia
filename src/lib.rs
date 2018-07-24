//! Ataxia game engine library code
#![deny(
    missing_debug_implementations, missing_copy_implementations, trivial_casts,
    trivial_numeric_casts, unsafe_code, unstable_features, unused_import_braces,
    unused_qualifications
)]
#![warn(missing_docs)]
extern crate failure;
#[macro_use]
extern crate serde_derive;
extern crate clap;
extern crate toml;

pub mod config;
pub mod server;

use failure::Error;

pub use config::Config;
pub use server::Server;

pub fn init() -> Result<(), Error> {
    // Initialize game
    //   Load initial game state
    //   Load database
    //   Load commands
    //   Load scripts
    //   Load world
    //   Load entities
    Ok(())
}

pub fn run() -> Result<(), Error> {
    // Clean up
    //   Save the world
    //   Shutdown Lua
    //   Flush pending database writes
    //   Close database connection
    Ok(())
}
