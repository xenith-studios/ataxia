//! Proxy module and related methods

use crate::Config;

/// Proxy data structure contains all related low-level data for running the network proxy
/// TODO: This is a stub data structure for now
#[derive(Debug)]
pub struct Proxy {
    config: Config,
}

impl Proxy {
    #![allow(clippy::new_ret_no_self)]
    /// Returns a new fully initialized `Proxy` server
    ///
    /// # Arguments
    ///
    /// * `config` - A Config structure, contains all necessary configuration
    ///
    pub fn new(config: Config) -> Result<Self, failure::Error> {
        // Initialize the proxy
        //   Set process start time
        Ok(Self { config })
    }

    /// Run the main loop
    pub fn run(self) -> Result<(), failure::Error> {
        // Main loop
        /*loop {
            // Poll all connections
            //   Handle new connections
            //   Handle new disconnects/logouts
            // Process all input events
            //   Send all processed events over RPC to engine process
            // Process all output events
            //   Send all processed output events to connections
            // Something something timing
        }*/

        // Main loop ends
        // Clean up
        Ok(())
    }
}
