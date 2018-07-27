//! Server module and related methods

use crate::Config;

/// Server data structure contains all related low-level data for running the game
/// TODO: This is a stub data structure for now
#[derive(Debug)]
pub struct Server {
    config: Config,
}

impl Server {
    /// Returns a new Server
    ///
    /// # Arguments
    ///
    /// * `config` - A Config structure, contains all necessary engine configuration
    ///
    /// # Errors
    ///
    ///
    pub fn new(config: Config) -> Result<Server, failure::Error> {
        // Initialize game
        //   Load initial game state
        //   Load database
        //   Load commands
        //   Load scripts
        //   Load world
        //   Load entities
        Ok(Server { config })
    }

    /// Run the big game loop
    pub fn run(self) -> Result<(), failure::Error> {
        // Clean up
        //   Save the world
        //   Shutdown Lua
        //   Flush pending database writes
        //   Close database connection
        Ok(())
    }
}
