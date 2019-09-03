//! Telnet contains code specifically to handle network I/O for a telnet connection
//!
use crate::proxy::NetSock;
use failure;
use futures::prelude::*;
use log::info;
use runtime::net::TcpListener;
use std::collections::BTreeMap;
use std::sync::atomic::{AtomicUsize, Ordering};
use std::sync::{Arc, Mutex};
//use uuid::Uuid;

/// An player connection
#[derive(Clone, Debug)]
pub struct Socket {
    uuid: String,
}

impl Socket {}

/// Server data structure holding all the server state
#[derive(Clone, Debug)]
pub struct Server {
    pub clients: Arc<Mutex<BTreeMap<usize, NetSock>>>,
}

impl Server {
    /// Async entry point for the telnet server
    pub async fn run(
        self,
        address: String,
        id_counter: Arc<AtomicUsize>,
    ) -> Result<(), failure::Error> {
        let mut socket = TcpListener::bind(&address)?;
        info!("Listening for telnet clients on {}", address);
        let mut incoming = socket.incoming();
        while let Some(stream) = incoming.next().await {
            let id_ref = id_counter.clone();
            let _clients_ref = self.clients.clone();
            runtime::spawn(async move {
                let _client_id = id_ref.fetch_add(1, Ordering::SeqCst);
                let stream = stream?;
                info!("Client connected: {}", stream.peer_addr()?);
                Ok::<(), failure::Error>(())
            });
        }
        Ok(())
    }
}
