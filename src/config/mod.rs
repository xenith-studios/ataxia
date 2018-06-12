//! Configuration module for Ataxia
//!
//!
use std::fs::File;
use std::io::Read;
use std::path::Path;

use toml;

use failure::Error;

#[derive(Deserialize, Debug)]
pub struct Config {
    proxy_addr: String,
    engine_pid_file: String,
    engine_log_file: String,
}

impl Config {
    /// Read a Config file from the file specified at path.
    pub fn read_config(path: &Path) -> Result<Config, Error> {
        let mut input = String::new();
        File::open(path)?.read_to_string(&mut input)?;
        let toml = toml::from_str::<Config>(&input)?;
        Ok(toml)
    }

    pub fn get_proxy_addr(&self) -> &str {
        self.proxy_addr.as_ref()
    }
    pub fn set_proxy_addr(&mut self, addr: &str) {
        self.proxy_addr = addr.to_string();
    }

    pub fn get_pid_file(&self) -> &str {
        self.engine_pid_file.as_ref()
    }
    pub fn set_pid_file(&mut self, file: &str) {
        self.engine_pid_file = file.to_string();
    }

    pub fn get_log_file(&self) -> &str {
        self.engine_log_file.as_ref()
    }
}
