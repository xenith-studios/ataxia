//! Configuration module for Ataxia
use clap::Parser;
use serde::Deserialize;
use std::fs::File;
use std::io::Read;
use std::path::PathBuf;

/// Config structure for holding internal and external configuration data
#[derive(Debug, Deserialize, Parser)]
#[command(name = "Ataxia MUD",
    about = env!("CARGO_PKG_DESCRIPTION"),
    version = env!("CARGO_PKG_VERSION"),
)]
pub struct Config {
    #[arg(short = 'c', long = "config", default_value = "data/ataxia.toml")]
    #[serde(default)]
    config_file: PathBuf,
    #[arg(short = 'p', long = "pid")]
    pid_file: Option<String>,
    #[arg(short = 'T', long)]
    telnet_addr: Option<String>,
    #[arg(short = 'W', long)]
    ws_addr: Option<String>,
    #[arg(short = 'M', long)]
    mq_addr: Option<String>,
    #[arg(short = 'L', long)]
    log_file: Option<String>,
    #[arg(short, long)]
    #[serde(default)]
    debug: bool,
    #[arg(short, long)]
    #[serde(default)]
    verbose: bool,
}
impl Config {
    /// Returns a new Config
    /// Read configuration from the file path specified in the Clap arguments struct.
    ///
    /// # Errors
    ///
    /// * Returns `std::io::Error` if the config file can't be opened or read
    /// * Returns `toml::de::Error` if TOML parsing fails
    ///
    /// # Panics
    ///
    /// TODO: add possible panics here
    pub fn new() -> Result<Self, anyhow::Error> {
        let cli = Config::parse();

        // TODO: This is a very simplistic method that should be improved/strengthened
        let process_name = std::env::args()
            .next()
            .unwrap()
            .split('/')
            .last()
            .unwrap()
            .to_string();

        let mut input = String::new();
        File::open(&cli.config_file)?.read_to_string(&mut input)?;
        let mut config = toml::from_str::<Self>(&input)?;

        if let Some(pid_file) = cli.pid_file {
            config.pid_file = Some(pid_file);
        } else if config.pid_file.is_none() {
            // The PID file wasn't specified. Default to process name
            config.pid_file = Some(format!("data/{process_name}.pid"));
        }

        if let Some(log_file) = cli.log_file {
            config.log_file = Some(log_file);
        } else if config.log_file.is_none() {
            // The log file wasn't specified. Default to process name
            config.log_file = Some(format!("logs/{process_name}.log"));
        }

        if let Some(ws_addr) = cli.ws_addr {
            config.ws_addr = Some(ws_addr);
        }

        if let Some(telnet_addr) = cli.telnet_addr {
            config.telnet_addr = Some(telnet_addr);
        }

        if let Some(mq_addr) = cli.mq_addr {
            config.mq_addr = Some(mq_addr);
        }

        config.debug = cli.debug;

        config.verbose = cli.verbose;

        Ok(config)
    }

    /// Returns the listen address for player telnet connections
    #[must_use]
    pub fn telnet_addr(&self) -> &str {
        self.telnet_addr
            .as_ref()
            .map_or("", |telnet_addr| telnet_addr)
    }
    /// Set the listen address for player telnet connections
    pub fn set_telnet_addr(&mut self, addr: String) {
        self.telnet_addr = Some(addr);
    }

    /// Returns the listen address for player websocket connections
    #[must_use]
    pub fn ws_addr(&self) -> &str {
        self.ws_addr.as_ref().map_or("", |ws_addr| ws_addr)
    }
    /// Set the listen address for player websocket connections
    pub fn set_ws_addr(&mut self, addr: String) {
        self.ws_addr = Some(addr);
    }

    /// Returns the listen address for the message queue
    #[must_use]
    pub fn mq_addr(&self) -> &str {
        self.mq_addr.as_ref().map_or("", |mq_addr| mq_addr)
    }
    /// Set the listen address for the message queue
    pub fn set_mq_addr(&mut self, addr: String) {
        self.mq_addr = Some(addr);
    }

    /// Returns the file path to the pid file
    #[must_use]
    pub fn pid_file(&self) -> &str {
        self.pid_file.as_ref().map_or("", |pid_file| pid_file)
    }
    /// Set the file path to the pid file
    pub fn set_pid_file(&mut self, file: String) {
        self.pid_file = Some(file);
    }

    /// Returns the file path to the log file
    #[must_use]
    pub fn log_file(&self) -> &str {
        self.log_file.as_ref().map_or("", |log_file| log_file)
    }
    /// Set the file path to the log file
    pub fn set_log_file(&mut self, file: String) {
        self.log_file = Some(file);
    }

    /// Returns true if the debug CLI flag was specified
    #[must_use]
    pub fn debug(&self) -> bool {
        self.debug
    }

    /// Returns true if the verbose CLI flag was specified
    #[must_use]
    pub fn verbose(&self) -> bool {
        self.verbose
    }
}
