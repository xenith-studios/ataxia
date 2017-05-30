//! Configuration module for Ataxia
//!
//!
use std::io::Read;
use std::fs::File;
use std::path::Path;

use toml;

use ::errors::{ConfigResult, ConfigResultExt};

#[derive(Deserialize, Debug)]
pub struct Config {
    main_port: u16,
    admin_port: u16,
    build_port: u16,
    proxy_pid_file: String,
    engine_pid_file: String,
    listen_addr: String,
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

    pub fn get_listen_addr(&self) -> &str {
        self.listen_addr.as_ref()
    }
}
