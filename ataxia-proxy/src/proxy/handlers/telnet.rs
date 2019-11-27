//! Telnet contains code specifically to handle network I/O for a telnet connection
//!
use crate::proxy::NetSock;
use failure;
use futures::prelude::*;
use log::info;
use std::collections::BTreeMap;
use std::sync::atomic::{AtomicUsize, Ordering};
use std::sync::{Arc, Mutex};
use tokio::net::TcpListener;
use tokio::prelude::*;
use uuid::Uuid;

/// A player connection
#[derive(Debug)]
pub struct Socket {
    uuid: Uuid,
    stream: tokio::net::TcpStream,
}

impl Socket {
    #[must_use]
    pub fn new(stream: tokio::net::TcpStream) -> Self {
        Self {
            uuid: Uuid::new_v4(),
            stream,
        }
    }

    pub async fn handle(mut self) {
        self.stream
            .write_all(b"You have connected. Goodbye!\r\n")
            .await
            .unwrap();
    }
}

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
    /// * Returns `tokio::io::Error` if the server can't bind to the listen port
    ///
    pub async fn new(
        address: String,
        clients: Arc<Mutex<BTreeMap<usize, NetSock>>>,
        id_counter: Arc<AtomicUsize>,
    ) -> Result<Self, failure::Error> {
        let listener = TcpListener::bind(&address).await?;
        info!("Listening for telnet clients on {}", address);
        Ok(Self {
            listener,
            clients,
            id_counter,
        })
    }
    /// Start the listener loop, which will spawn individual connections into the runtime
    pub async fn run(mut self) {
        let mut incoming = self.listener.incoming();
        while let Some(Ok(stream)) = incoming.next().await {
            let client_id = self.id_counter.fetch_add(1, Ordering::Relaxed);
            tokio::spawn(async move {
                info!(
                    "Telnet client connected: ID: {}, remote_addr: {}",
                    client_id,
                    stream.peer_addr().unwrap()
                );
                // Create account/socket struct
                let socket = Socket::new(stream);
                socket.handle().await;
            });
        }
    }
}
