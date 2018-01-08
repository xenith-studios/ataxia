//! Configuration module for Ataxia
//!
//!
use std::io::Read;
use std::fs::File;
use std::path::Path;

use toml;

use errors::{ConfigResult, ConfigResultExt};

#[derive(Deserialize, Debug)]
pub struct Config {
    main_port: u16,
    proxy_addr: String,
    engine_pid_file: String,
    engine_log_file: String,
}

impl Config {
    /// Read a Config file from the file specified at path.
    pub fn read_config(path: &Path) -> ConfigResult<Config> {
        let mut file: File = File::open(path)?;
        let mut input: String = String::new();
        file.read_to_string(&mut input)?;
        toml::from_str::<Config>(&input).chain_err(|| "Failed to parse TOML")
    }

    pub fn get_main_port(&self) -> u16 {
        self.main_port
    }

    pub fn get_proxy_addr(&self) -> &str {
        self.proxy_addr.as_ref()
    }

    pub fn get_pid_file(&self) -> &str {
        self.engine_pid_file.as_ref()
    }

    pub fn get_log_file(&self) -> &str {
        self.engine_log_file.as_ref()
    }
}
