//! Proxy module and related methods

pub mod handlers;

use std::collections::BTreeMap;
use std::sync::atomic::AtomicUsize;
use std::sync::{Arc, Mutex};

use self::handlers::{telnet, websockets};
use ataxia_core::Config;
use tokio::runtime::Runtime;

/// Socket enum that stores multiple types of sockets.
#[derive(Clone, Debug, Copy)]
pub enum NetSock {
    /// Telnet connection
    Telnet(usize),
    /// Websocket connection
    Websocket(usize),
}

/// Proxy data structure contains all related low-level data for running the network proxy
#[derive(Debug)]
pub struct Proxy {
    config: Config,
    clients: Arc<Mutex<BTreeMap<usize, NetSock>>>,
    telnet_server: telnet::Server,
    ws_server: websockets::Server,
    runtime: Runtime,
}

impl Proxy {
    #![allow(clippy::new_ret_no_self)]
    /// Returns a new fully initialized `Proxy` server
    ///
    /// # Arguments
    ///
    /// * `config` - A Config structure, contains all necessary configuration
    /// * 'rt' - The `tokio::runtime::Runtime` used to run the async I/O
    ///
    pub fn new(config: Config, mut rt: Runtime) -> Result<Self, failure::Error> {
        // Initialize the proxy
        let id_counter = Arc::new(AtomicUsize::new(1));
        let client_list = Arc::new(Mutex::new(BTreeMap::new()));
        let telnet_addr = config.telnet_addr().to_string();
        let ws_addr = config.ws_addr().to_string();
        //TODO: Set proxy start time

        Ok(Self {
            config,
            clients: client_list.clone(),
            telnet_server: rt.block_on(telnet::Server::new(
                telnet_addr,
                client_list.clone(),
                id_counter.clone(),
            ))?,
            ws_server: rt.block_on(websockets::Server::new(ws_addr, client_list, id_counter))?,
            runtime: rt,
        })
    }

    /// Run the main loop
    pub fn run(mut self) -> Result<(), failure::Error> {
        // Start the network I/O
        let telnet = self.runtime.spawn(self.telnet_server.run());
        let ws = self.runtime.spawn(self.ws_server.run());

        // Main loop
        /*loop {
            // Process all input events
            //   Send all processed events over MQ to engine process
            // Process all output events
            //   Send all processed output events to connections
            // Something something timing
        }*/
        // Main loop ends

        // Hold main thread open until the runtime has shutdown.
        // TODO: This can definitely be done better, cleanup later
        let _ = self.runtime.block_on(telnet);
        let _ = self.runtime.block_on(ws);
        // Clean up
        Ok(())
    }
}
