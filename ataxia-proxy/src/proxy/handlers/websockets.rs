//! Websocket contains code specifically to handle network I/O for a websocket connection
//!
use crate::proxy::NetSock;
use failure;
use futures::prelude::*;
use log::info;
use tokio::net::TcpListener;
//use tokio::prelude::*;
use std::collections::BTreeMap;
use std::sync::atomic::{AtomicUsize, Ordering};
use std::sync::{Arc, Mutex};
//use uuid::Uuid;

/// A player connection
#[derive(Clone, Debug)]
pub struct Socket {
    uuid: String,
}

impl Socket {}

/// Server data structure holding all the server state
#[derive(Debug)]
pub struct Server {
    listener: TcpListener,
    clients: Arc<Mutex<BTreeMap<usize, NetSock>>>,
    id_counter: Arc<AtomicUsize>,
}

impl Server {
    /// Returns a new Server
    ///
    /// # Arguments
    ///
    /// * `address` - A String containing the listen addr:port
    /// * `clients` - A shared binding to the client list
    /// * `id_counter` - A shared binding to a global connection counter
    ///
    /// # Errors
    ///
    /// * Returns tokio::io::Error if the server can't bind to the listen port
    ///
    pub async fn new(
        address: String,
        clients: Arc<Mutex<BTreeMap<usize, NetSock>>>,
        id_counter: Arc<AtomicUsize>,
    ) -> Result<Self, failure::Error> {
        let listener = TcpListener::bind(&address).await?;
        info!("Listening for websocket clients on {}", address);
        Ok(Self {
            listener,
            clients,
            id_counter,
        })
    }
    /// Start the listener loop, which will spawn individual connections into the runtime
    pub async fn run(self) {
        let mut incoming = self.listener.incoming();
        while let Some(stream) = incoming.next().await {
            let id_ref = self.id_counter.clone();
            let _clients_ref = self.clients.clone();
            tokio::spawn(async move {
                let _client_id = id_ref.fetch_add(1, Ordering::SeqCst);
                let stream = stream.unwrap();
                info!("Client connected: {}", stream.peer_addr().unwrap());
            });
        }
    }
}
