//! Websockets contains code specifically to handle network I/O for a websocket connection
//!
use crate::portal::{Message, Rx, Tx};
use anyhow::anyhow;
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
    /// Returns a new Socket
    ///
    /// # Arguments
    ///
    /// * `stream` - A `TcpStream` from Tokio
    /// * `id` - A unique incrementing connection identifier
    /// * `main_tx` - Transmit channel to communicate with the main task
    ///
    /// # Errors
    ///
    /// * This function will return an error if any send/recv operations fail on any channel or
    /// stream, which would signal a connection error and disconnect the client.
    pub async fn new(
        stream: tokio::net::TcpStream,
        id: usize,
        main_tx: Tx,
    ) -> Result<Self, anyhow::Error> {
        let addr = match stream.peer_addr() {
            Ok(addr) => addr.to_string(),
            Err(_) => "Unknown".to_string(),
        };
        info!(
            "Websocket client connected: ID: {}, remote_addr: {}",
            id, addr
        );
        let mut stream = Framed::new(stream, LinesCodec::new());
        let (tx, rx) = mpsc::unbounded_channel();
        stream.send("Welcome to the Ataxia Portal.").await?;
        stream.send("Please enter your username:").await?;
        let username = match stream.next().await {
            Some(Ok(line)) => line,
            Some(Err(_)) | None => {
                return Err(anyhow!("Failed to get username from {}.", addr));
            }
        };
        stream.send(format!("Welcome, {}!", username)).await?;
        main_tx.send(Message::NewConnection(id, tx, username))?;
        Ok(Self {
            uuid: Uuid::new_v4(),
            id,
            stream,
            rx,
            main_tx,
            addr,
        })
    }

    /// Handle a connection
    ///
    /// # Errors
    ///
    /// * This function will return an error and disconnect the client if any send/recv fails.
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
        info!("Listening for websocket clients on {:?}", address);
        Ok(Self {
            listener,
            id_counter,
            main_tx: tx,
        })
    }
    /// Start the listener loop, which will spawn individual connections into the runtime
    pub async fn run(mut self) {
        while let Some(connection) = self.listener.next().await {
            match connection {
                Err(e) => error!("Accept failed: {:?}", e),
                Ok(stream) => {
                    let client_id = self.id_counter.fetch_add(1, Ordering::Relaxed);
                    let main_tx = self.main_tx.clone();
                    tokio::spawn(async move {
                        let socket = match Socket::new(stream, client_id, main_tx).await {
                            Ok(socket) => socket,
                            Err(e) => {
                                error!("Client disconnected: {}", e);
                                return;
                            }
                        };
                        if let Err(e) = socket.handle().await {
                            error!("Client error: {}", e);
                            return;
                        };
                    });
                }
            }
        }
    }
}
