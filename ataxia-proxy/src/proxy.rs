//! Proxy module and related methods

pub mod handlers;

use ataxia_core::Config;
use ataxia_events::EventLoop;
use handlers::websockets::Socket as WSocket;
use std::collections::BTreeMap;
use std::sync::{Arc, Mutex};

/// Socket enum that stores multiple types of sockets.
#[derive(Clone, Debug)]
pub enum NetSock {
    /// Telnet connection
    Telnet, // Unused for now
    /// Websocket connection
    Websocket(Arc<Mutex<WSocket>>),
}

impl NetSock {
    /// Send data to a connection
    pub fn send(&mut self, data: &str) {
        //TODO: Add send handler for telnet sockets.
        match self {
            NetSock::Websocket(ref socket) => {
                let mut guard = socket.lock().unwrap();
                (*guard).send(data).unwrap();
            }
            NetSock::Telnet => {}
        }
    }
}

/// Proxy data structure contains all related low-level data for running the network proxy
/// TODO: This is a stub data structure for now
#[derive(Debug)]
pub struct Proxy {
    config: Config,
    clients: Arc<Mutex<BTreeMap<String, NetSock>>>,
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
        Ok(Self {
            config,
            clients: Arc::new(Mutex::new(BTreeMap::new())),
        })
    }

    /// Run the main loop
    pub fn run(self) -> Result<(), failure::Error> {
        // Main loop
        let eventloop: Arc<EventLoop> = Arc::new(EventLoop::new());
        let websocket_thread =
            handlers::websockets::create_server(None, 45678, &self.clients.clone(), &eventloop);
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
        websocket_thread.join().unwrap();
        // Main loop ends
        // Clean up
        Ok(())
    }
}
