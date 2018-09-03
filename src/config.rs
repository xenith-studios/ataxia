//! Configuration module for Ataxia
use serde_derive::Deserialize;
use std::fs::File;
use std::io::Read;
use std::path::Path;

/// Config structure for holding internal and external configuration data
#[derive(Deserialize, Debug)]
pub struct Config {
    proxy_addr: String,
    engine_pid_file: String,
    engine_log_file: String,
    #[serde(default)]
    debug: bool,
    #[serde(default)]
    verbose: bool,
}

impl Config {
    /// Returns a new Config
    /// Read configuration from the file path specified in the argument structure.
    ///
    /// # Arguments
    ///
    /// * `matches` - A clap::ArgMatches structure containing command-line arguments and default values
    ///
    /// # Errors
    ///
    ///
    pub fn new(matches: &clap::ArgMatches<'_>) -> Result<Config, failure::Error> {
        let path = Path::new(matches.value_of("config").unwrap_or("data/ataxia.toml"));

        let mut input = String::new();
        File::open(path)?.read_to_string(&mut input)?;
        let mut config = toml::from_str::<Config>(&input)?;

        if let Some(pid_file) = matches.value_of("pid_file") {
            config.set_pid_file(pid_file);
        }

        if let Some(proxy_addr) = matches.value_of("proxy_addr") {
            config.set_proxy_addr(proxy_addr);
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

    /// Returns the listen address of the network proxy process
    pub fn proxy_addr(&self) -> &str {
        self.proxy_addr.as_ref()
    }
    /// Set the listen address of the network proxy process
    pub fn set_proxy_addr(&mut self, addr: &str) {
        self.proxy_addr = addr.to_string();
    }

    /// Returns the file path to the pid file
    pub fn pid_file(&self) -> &str {
        self.engine_pid_file.as_ref()
    }
    /// Set the file path to the pid file
    pub fn set_pid_file(&mut self, file: &str) {
        self.engine_pid_file = file.to_string();
    }

    /// Returns the file path to the log file
    pub fn log_file(&self) -> &str {
        self.engine_log_file.as_ref()
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
