//! Portal module and related methods

pub mod handlers;

use std::collections::BTreeMap;
use std::sync::atomic::AtomicUsize;
use std::sync::Arc;

//use self::handlers::{telnet, websockets};
use self::handlers::telnet;
use ataxia_core::Config;
use log::info;
use tokio::sync::mpsc;

/// Shorthand for the transmit half of the message channel.
type Tx = mpsc::UnboundedSender<Message>;

/// Shorthand for the receive half of the message channel.
type Rx = mpsc::UnboundedReceiver<Message>;

/// Message is a wrapper for data passed between the main thread and the tasks
#[derive(Debug)]
pub enum Message {
    /// A message that should be broadcasted to others.
    Data(usize, String),

    /// A new connection
    NewConnection(usize, Tx, String),

    /// Remove an existing connection
    CloseConnection(usize),
}

/// portal data structure contains all related low-level data for running the network portal
#[derive(Debug)]
pub struct Portal {
    config: Config,
    client_list: BTreeMap<usize, (String, Tx)>,
    telnet_server: telnet::Server,
    //ws_server: websockets::Server,
    rx: Rx,
}

impl Portal {
    #![allow(clippy::new_ret_no_self, clippy::eval_order_dependence)]
    /// Returns a new fully initialized `portal` server
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
        // Initialize the portal
        let id_counter = Arc::new(AtomicUsize::new(1));
        let client_list = BTreeMap::new();
        let telnet_addr = config.telnet_addr().to_string();
        //let ws_addr = config.ws_addr().to_string();
        let (tx, rx) = mpsc::unbounded_channel();
        //TODO: Set portal start time

        Ok(Self {
            config,
            client_list,
            telnet_server: telnet::Server::new(&telnet_addr, id_counter.clone(), tx.clone())
                .await?,
            //ws_server: websockets::Server::new(&ws_addr, id_counter, tx).await?,
            rx,
        })
    }

    /// Run the main loop
    ///
    /// # Errors
    ///
    /// * Does not currently return any errors
    pub async fn run(mut self) -> Result<(), anyhow::Error> {
        // Start the network I/O servers
        tokio::spawn(self.telnet_server.run());
        //tokio::spawn(self.ws_server.run());

        // Main loop
        while let Some(message) = self.rx.recv().await {
            // Process all input events
            //   Send all processed events over MQ to engine process
            // Process all output events
            //   Send all processed output events to connections
            // Something something timing
            match message {
                Message::NewConnection(id, rx, name) => {
                    info!("Player {} has connected on socket {}", name, id);
                    self.client_list.values().for_each(|(_, tx)| {
                        tx.send(Message::Data(id, format!("{} has joined the chat.", name)))
                            .unwrap();
                    });
                    self.client_list.insert(id, (name, rx));
                }
                Message::CloseConnection(id) => {
                    if let Some((name, _)) = self.client_list.remove(&id) {
                        info!("Player {} has disconnected on socket {}", name, id);
                        self.client_list.values().for_each(|(_, tx)| {
                            tx.send(Message::Data(id, format!("{} has left the chat.", name)))
                                .unwrap();
                        });
                    }
                }
                Message::Data(id, message) => {
                    if let Some((name, _)) = self.client_list.get(&id) {
                        info!("Received message from {}: {}", name, message);
                        self.client_list.values().for_each(|(tx_name, tx)| {
                            if tx_name != name {
                                tx.send(Message::Data(id, format!("{}: {}", name, message)))
                                    .unwrap();
                            }
                        });
                    }
                }
            }
        }
        // Main loop ends

        // Clean up
        Ok(())
    }
}
