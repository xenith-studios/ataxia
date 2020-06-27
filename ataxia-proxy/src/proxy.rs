//! Proxy module and related methods

pub mod handlers;

use std::collections::BTreeMap;
use std::sync::atomic::AtomicUsize;
use std::sync::Arc;
use std::thread;

use self::handlers::{telnet, websockets};
use ataxia_core::Config;
use crossbeam::crossbeam_channel::unbounded;
use log::info;
use tokio::runtime::Runtime;

/// Shorthand for the transmit half of the message channel.
type Tx = crossbeam::channel::Sender<Message>;

/// Shorthand for the receive half of the message channel.
type Rx = crossbeam::channel::Receiver<Message>;

/// Message is a wrapper for data passed between the main thread and the tasks
#[derive(Debug)]
pub enum Message {
    /// A message that should be broadcasted to others.
    Data(String),

    /// A new connection
    NewConnection(usize, Tx, String),
}

/// Proxy data structure contains all related low-level data for running the network proxy
#[derive(Debug)]
pub struct Proxy {
    config: Config,
    client_list: BTreeMap<usize, Tx>,
    telnet_server: telnet::Server,
    ws_server: websockets::Server,
    rx: Rx,
}

impl Proxy {
    #![allow(clippy::new_ret_no_self, clippy::eval_order_dependence)]
    /// Returns a new fully initialized `Proxy` server
    ///
    /// # Arguments
    ///
    /// * `config` - A Config structure, contains all necessary configuration
    /// * 'rt' - The `tokio::runtime::Runtime` used to run the async I/O
    ///
    /// # Errors
    ///
    /// * Returns an error if the `Server` fails to initialize
    pub async fn new(config: Config) -> Result<Self, anyhow::Error> {
        // Initialize the proxy
        let id_counter = Arc::new(AtomicUsize::new(1));
        let client_list = BTreeMap::new();
        let telnet_addr = config.telnet_addr().to_string();
        let ws_addr = config.ws_addr().to_string();
        let (tx, rx) = unbounded();
        //TODO: Set proxy start time

        Ok(Self {
            config,
            client_list,
            telnet_server: telnet::Server::new(telnet_addr, id_counter.clone(), tx.clone()).await?,
            ws_server: websockets::Server::new(ws_addr, id_counter, tx).await?,
            rx,
        })
    }

    /// Run the main loop
    ///
    /// # Errors
    ///
    /// * Does not currently return any errors
    pub fn run(mut self, mut runtime: Runtime) -> Result<(), anyhow::Error> {
        // Start the network I/O and put the executor into a background thread
        let telnet = runtime.spawn(self.telnet_server.run());
        let ws = runtime.spawn(self.ws_server.run());
        let executor = thread::spawn(move || {
            // Hold executor thread open until the network tasks have shutdown.
            // TODO: This can definitely be done better, cleanup later
            let _ = runtime.block_on(ws);
            let _ = runtime.block_on(telnet);
        });

        // Main loop
        while let Some(message) = self.rx.iter().next() {
            // Process all input events
            //   Send all processed events over MQ to engine process
            // Process all output events
            //   Send all processed output events to connections
            // Something something timing
            match message {
                Message::NewConnection(id, rx, name) => {
                    self.client_list.insert(id, rx);
                    info!("Player {} has connected on socket {}", name, id);
                }
                Message::Data(message) => {
                    info!("Received message: {}", message);
                }
            }
        }
        // Main loop ends

        // Wait for the networking to shut down
        executor.join().unwrap();
        // Clean up
        Ok(())
    }
}
