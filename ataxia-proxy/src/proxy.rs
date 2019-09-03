//! Proxy module and related methods

pub mod handlers;

use self::handlers::{telnet, websockets};
use ataxia_core::Config;
use std::collections::BTreeMap;
use std::sync::atomic::AtomicUsize;
use std::sync::{Arc, Mutex};

/// Socket enum that stores multiple types of sockets.
#[derive(Clone, Debug, Copy)]
pub enum NetSock {
    /// Telnet connection
    Telnet(usize),
    /// Websocket connection
    Websocket(usize),
}

impl NetSock {
    /// Send data to a connection
    pub fn send(&mut self, _data: &str) {
        //TODO: Add send handler for telnet sockets.
        match self {
            Self::Websocket(ref _socket) => {
                unimplemented!();
            }
            Self::Telnet(ref _socket) => {}
        }
    }
}

/// Proxy data structure contains all related low-level data for running the network proxy
#[derive(Debug)]
pub struct Proxy {
    config: Config,
    clients: Arc<Mutex<BTreeMap<usize, NetSock>>>,
    ws_server: websockets::Server,
    telnet_server: telnet::Server,
}

impl Proxy {
    #![allow(clippy::new_ret_no_self)]
    /// Returns a new fully initialized `Proxy` server
    ///
    /// # Arguments
    ///
    /// * `config` - A Config structure, contains all necessary configuration
    ///
    pub fn new(config: Config) -> Result<Self, failure::Error> {
        // Initialize the proxy
        //   Set process start time
        let client_list = Arc::new(Mutex::new(BTreeMap::new()));
        Ok(Self {
            config,
            clients: client_list.clone(),
            ws_server: websockets::Server {
                clients: client_list.clone(),
            },
            telnet_server: telnet::Server {
                clients: client_list.clone(),
            },
        })
    }

    /// Run the main loop
    pub async fn run(self) -> Result<(), failure::Error> {
        // Main loop
        let id_counter = Arc::new(AtomicUsize::new(1));

        let telnet = runtime::spawn(
            self.telnet_server
                .run(self.config.telnet_addr().to_string(), id_counter.clone()),
        );
        let websocket = runtime::spawn(
            self.ws_server
                .run(self.config.ws_addr().to_string(), id_counter.clone()),
        );
        /*loop {
            // Poll all connections
            //   Handle new connections
            //   Handle new disconnects/logouts
            // Process all input events
            //   Send all processed events over RPC to engine process
            // Process all output events
            //   Send all processed output events to connections
            // Something something timing
        }*/
        // Hold main thread open until server thread is done.
        telnet.await?;
        websocket.await?;
        // Main loop ends
        // Clean up
        Ok(())
    }
}
