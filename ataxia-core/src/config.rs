//! Configuration module for Ataxia
use serde::Deserialize;
use std::fs::File;
use std::io::Read;
use std::path::Path;

/// Config structure for holding internal and external configuration data
#[derive(Deserialize, Debug)]
pub struct Config {
    #[serde(default)]
    ws_addr: String,
    #[serde(default)]
    telnet_addr: String,
    #[serde(default)]
    mq_addr: String,
    pid_file: String,
    log_file: String,
    #[serde(default)]
    debug: bool,
    #[serde(default)]
    verbose: bool,
}

impl Config {
    #![allow(clippy::new_ret_no_self)]
    /// Returns a new Config
    /// Read configuration from the file path specified in the Clap arguments struct.
    ///
    /// # Arguments
    ///
    /// * `matches` - A `clap::ArgMatches` structure containing command-line arguments and default values
    ///
    /// # Errors
    ///
    /// * Returns `std::io::Error` if the config file can't be opened or read
    /// * Returns `toml::de::Error` if TOML parsing fails
    ///
    pub fn new(matches: &clap::ArgMatches<'_>) -> Result<Self, failure::Error> {
        let path = Path::new(matches.value_of("config").unwrap_or("data/ataxia.toml"));

        let mut input = String::new();
        File::open(path)?.read_to_string(&mut input)?;
        let mut config = toml::from_str::<Self>(&input)?;

        if let Some(pid_file) = matches.value_of("pid_file") {
            config.pid_file = pid_file.to_string();
        }

        if let Some(ws_addr) = matches.value_of("ws_addr") {
            config.ws_addr = ws_addr.to_string();
        }

        if let Some(telnet_addr) = matches.value_of("telnet_addr") {
            config.telnet_addr = telnet_addr.to_string();
        }

        if let Some(mq_addr) = matches.value_of("mq_addr") {
            config.mq_addr = mq_addr.to_string();
        }

        config.debug = match matches.occurrences_of("debug") {
            0 => false,
            1 | _ => true,
        };

        config.verbose = match matches.occurrences_of("verbose") {
            0 => false,
            1 | _ => true,
        };

        Ok(config)
    }

    /// Returns the listen address for player websocket connections
    pub fn ws_addr(&self) -> &str {
        self.ws_addr.as_ref()
    }
    /// Set the listen address for player websocket connections
    pub fn set_ws_addr(&mut self, addr: String) {
        self.ws_addr = addr;
    }

    /// Returns the listen address player telnet connections
    pub fn telnet_addr(&self) -> &str {
        self.telnet_addr.as_ref()
    }
    /// Set the listen address for player telnet connections
    pub fn set_telnet_addr(&mut self, addr: String) {
        self.telnet_addr = addr;
    }

    /// Returns the listen address of the message queue
    pub fn mq_addr(&self) -> &str {
        self.mq_addr.as_ref()
    }
    /// Set the listen address of the message queue
    pub fn set_mq_addr(&mut self, addr: String) {
        self.mq_addr = addr;
    }

    /// Returns the file path to the pid file
    pub fn pid_file(&self) -> &str {
        self.pid_file.as_ref()
    }
    /// Set the file path to the pid file
    pub fn set_pid_file(&mut self, file: String) {
        self.pid_file = file;
    }

    /// Returns the file path to the log file
    pub fn log_file(&self) -> &str {
        self.log_file.as_ref()
    }

    /// Returns true if the debug CLI flag was specified
    pub fn debug(&self) -> bool {
        self.debug
    }

    /// Returns true if the verbose CLI flag was specified
    pub fn verbose(&self) -> bool {
        self.verbose
    }
}
