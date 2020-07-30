//! Telnet contains code specifically to handle network I/O for a telnet connection
//!
use crate::proxy::{Message, Rx, Tx};
use futures::prelude::*;
use log::{error, info};
use std::sync::atomic::{AtomicUsize, Ordering};
use std::sync::Arc;
use tokio::net::TcpListener;
use tokio::sync::mpsc;
use tokio_util::codec::{Framed, LinesCodec};
use uuid::Uuid;

/// A player connection
#[derive(Debug)]
pub struct Socket {
    uuid: Uuid,
    id: usize,
    stream: Framed<tokio::net::TcpStream, LinesCodec>,
    rx: Rx,
    main_tx: Tx,
    addr: String,
}

impl Socket {
    #[must_use]
    /// Returns a new Socket
    ///
    /// # Arguments
    ///
    /// * `stream` - A `TcpStream` from Tokio
    pub fn new(
        stream: Framed<tokio::net::TcpStream, LinesCodec>,
        id: usize,
        rx: Rx,
        main_tx: Tx,
        addr: String,
    ) -> Self {
        Self {
            uuid: Uuid::new_v4(),
            id,
            stream,
            rx,
            main_tx,
            addr,
        }
    }

    /// Handle a connection
    ///
    /// # Arguments
    ///
    /// # Errors
    #[allow(clippy::mut_mut)]
    pub async fn handle(mut self) -> Result<(), anyhow::Error> {
        loop {
            futures::select! {
                data = self.rx.recv().fuse() => {
                    if let Some(message) = data {
                        match message {
                            Message::Data(_, message) => self.stream.send(message).await?,
                            _ => error!("Oops"),
                        }
                    } else {
                        // The main loop closed our channel to signal system shutdown
                        self.stream.send("The game is shutting down, goodbye!").await?;
                        break;
                    }
                },
                data = self.stream.next().fuse() => {
                    if let Some(message) = data {
                        if let Err(e) = self.main_tx.send(Message::Data(self.id, message?)) {
                            // The main loop channel was closed, most likely this signals a crash
                            self.stream.send("The game has crashed, sorry! Please try again.").await?;
                            break;
                        }
                    } else {
                        // The other end closed the connection. Nothing left to do.
                        break;
                    }
                },
            };
        }
        self.main_tx.send(Message::CloseConnection(self.id))?;
        Ok(())
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
        address: &str,
        id_counter: Arc<AtomicUsize>,
        tx: Tx,
    ) -> Result<Self, anyhow::Error> {
        let listener = TcpListener::bind(address).await?;
        info!("Listening for telnet clients on {:?}", address);
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
                    let addr: String = match stream.peer_addr() {
                        Ok(addr) => addr.to_string(),
                        Err(_) => "Unknown".to_string(),
                    };
                    let mut stream = Framed::new(stream, LinesCodec::new());
                    tokio::spawn(async move {
                        info!(
                            "Telnet client connected: ID: {}, remote_addr: {}",
                            client_id, addr
                        );
                        // Create a channel for this player
                        let (tx, rx) = mpsc::unbounded_channel();
                        stream.send("Welcome to the Ataxia Portal.").await.unwrap();
                        stream.send("Please enter your username:").await.unwrap();
                        let username = match stream.next().await {
                            Some(Ok(line)) => line,
                            // We didn't get a line so we return early here.
                            Some(Err(_)) | None => {
                                info!("Failed to get username from {}. Client disconnected.", addr);
                                return;
                            }
                        };
                        stream
                            .send(format!("Welcome, {}!", username))
                            .await
                            .unwrap();
                        main_tx
                            .send(Message::NewConnection(client_id, tx, username))
                            .unwrap();
                        // Create account/socket struct
                        let socket = Socket::new(stream, client_id, rx, main_tx, addr);
                        if let Err(e) = socket.handle().await {
                            error!("Client error: {}", e);
                        };
                    });
                }
            }
        }
    }
}
