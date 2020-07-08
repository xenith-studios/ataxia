//! Telnet contains code specifically to handle network I/O for a telnet connection
//!
use crate::proxy::{Message, Rx, Tx};
use futures::prelude::*;
use log::{error, info};
use std::sync::atomic::{AtomicUsize, Ordering};
use std::sync::Arc;
use tokio::net::TcpListener;
use tokio::prelude::*;
use tokio::sync::mpsc;
use uuid::Uuid;

/// A player connection
#[derive(Debug)]
pub struct Socket {
    uuid: Uuid,
    id: usize,
    stream: tokio::net::TcpStream,
    rx: Rx,
    main_tx: Tx,
}

impl Socket {
    #[must_use]
    /// Returns a new Socket
    ///
    /// # Arguments
    ///
    /// * `stream` - A `TcpStream` from Tokio
    pub fn new(stream: tokio::net::TcpStream, id: usize, rx: Rx, main_tx: Tx) -> Self {
        Self {
            uuid: Uuid::new_v4(),
            id,
            stream,
            rx,
            main_tx,
        }
    }

    /// Handle a connection
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
    id_counter: Arc<AtomicUsize>,
    main_tx: Tx,
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
        id_counter: Arc<AtomicUsize>,
        tx: Tx,
    ) -> Result<Self, anyhow::Error> {
        let listener = TcpListener::bind(&address).await?;
        info!("Listening for telnet clients on {}", address);
        Ok(Self {
            listener,
            id_counter,
            main_tx: tx,
        })
    }
    /// Start the listener loop, which will spawn individual connections into the runtime
    pub async fn run(mut self) {
        let mut incoming = self.listener.incoming();
        while let Some(connection) = incoming.next().await {
            match connection {
                Err(e) => error!("Accept failed: {:?}", e),
                Ok(stream) => {
                    let client_id = self.id_counter.fetch_add(1, Ordering::Relaxed);
                    let main_tx = self.main_tx.clone();
                    tokio::spawn(async move {
                        info!(
                            "Telnet client connected: ID: {}, remote_addr: {}",
                            client_id,
                            stream.peer_addr().unwrap()
                        );
                        // Create a channel for this player
                        let (tx, rx) = mpsc::unbounded_channel();
                        main_tx
                            .send(Message::NewConnection(client_id, tx, "Unknown".to_string()))
                            .unwrap();
                        // Create account/socket struct
                        let socket = Socket::new(stream, client_id, rx, main_tx);
                        socket.handle().await;
                    });
                }
            }
        }
    }
}
