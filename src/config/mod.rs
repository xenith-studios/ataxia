//! Configuration module for Ataxia
//!
//!
use std::io::Read;
use std::fs::File;
use std::path::Path;
use std::str::FromStr;

use toml::{Decoder, Value};
use rustc_serialize::Decodable;

use ::errors::{ConfigErrorKind, ConfigResult, ConfigError};

#[derive(RustcDecodable, Debug)]
pub struct Config {
    main_port: u16,
    admin_port: u16,
    build_port: u16,
    pid_file: String,
}

impl Config {
    /// Read a Config file from the file specified at path.
    pub fn read_config(path: &Path) -> ConfigResult<Config> {
        let mut file: File = try!(File::open(path));
        let mut toml: String = String::new();
        try!(file.read_to_string(&mut toml));
        toml.parse()
    }

    pub fn get_main_port(&self) -> u16 {
        self.main_port
    }
}

impl FromStr for Config {
    type Err = ConfigError;

    fn from_str(toml: &str) -> ConfigResult<Config> {
        let value: Value = try!(toml.parse().map_err(|vec| ConfigErrorKind::VecParserError(vec)));
        let mut decoder: Decoder = Decoder::new(value);
        Ok(try!(Self::decode(&mut decoder)))
    }
}
