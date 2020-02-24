//! Engine module and related methods

use ataxia_core::Config;

/// Engine data structure contains all related low-level data for running the game
/// TODO: This is a stub data structure for now
#[derive(Debug)]
pub struct Engine {
    config: Config,
}

impl Engine {
    #![allow(clippy::new_ret_no_self)]
    /// Returns a new fully initialized game `Engine`
    ///
    /// # Arguments
    ///
    /// * `config` - A Config structure, contains all necessary configuration
    ///
    /// # Errors
    ///
    /// * Does not currently return any errors
    ///
    pub fn new(config: Config) -> Result<Self, failure::Error> {
        // Initialize game
        //   Set game start time
        //   Initialize Lua
        //   Load game data
        //   Load commands
        //   Load world data
        //   Load all entities (populate world)
        Ok(Self { config })
    }

    /// Run the big game loop
    ///
    /// # Errors
    ///
    /// * Does not currently return any errors
    pub fn run(self) -> Result<(), failure::Error> {
        let _ = self;
        // Main game loop
        /*loop {
            // Read network input channel and process all pending external events
            //   Did we get a new player login? If so, create the entity and add them to the game
            //   Did the player just quit? If so, remove them from the game and delete the entity
            //   Process all player commands
            // Process all server events (weather, time, zone updates, NPC updates, etc)
            // Process all output events and write them to the network output channel
            // Something something timing (ticks/pulses, sleep(pulse_time - how_long_this_loop_took))
        }*/

        // Game loop ends
        // Clean up
        //   Save the world
        //   Save game data
        //   Shutdown Lua
        Ok(())
    }
}
