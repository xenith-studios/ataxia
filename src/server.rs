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
        //   Set game start time
        //   Initialize Lua
        //   Load game data
        //   Load commands
        //   Load world data
        //   Load all entities (populate world)
        Ok(Server { config })
    }

    /// Run the big game loop
    pub fn run(self) -> Result<(), failure::Error> {
        // Main game loop
        /*loop {
            // Read network input channel and process all pending external events
            //   Did we get a new player login? If so, create the entity and add them to the game
            //   Did the player just quit? If so, remove them from the game and delete the entity
            // Process all server events (weather, time, zone updates, etc)
            // Process all output events and write them to the network output channel
            // Something something timing (ticks/pulses)
            break;
        }*/

        // Game loop ends
        // Clean up
        //   Save the world
        //   Save game data
        //   Shutdown Lua
        Ok(())
    }
}
