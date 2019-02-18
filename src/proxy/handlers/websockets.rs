use serde::{Serialize, Deserialize};
use std::thread;
use std::sync::{Arc, Mutex};
use uuid::Uuid;
use ws::{listen, Handler, Sender, Result as WSResult, Message, CloseCode, Handshake};
use crate::proxy::NetSock;
use std::collections::BTreeMap;
use ataxia_events::EventLoop;

#[derive(Clone, Debug)]
pub struct Socket {
    out: Sender,
    clients: Arc<Mutex<BTreeMap<String, NetSock>>>,
    events: Arc<EventLoop>,
    uuid: String
}

impl Socket {
    pub fn send(&mut self, msg: &str) -> WSResult<()> {
        self.out.send(Message::from(msg))
    }
}

impl Handler for Socket {

    fn on_open(&mut self, shake: Handshake) -> WSResult<()> {
        let mut guard = self.clients.lock().unwrap();
        (*guard).insert(self.uuid.clone(), NetSock::Websocket(Arc::new(Mutex::new(self.clone()))));
        Ok(())
    }

    fn on_message(&mut self, msg: Message) -> WSResult<()> {
        // Handle message data
        Ok(())
    }

    fn on_close(&mut self, code: CloseCode, reason: &str) {
        match code {
            CloseCode::Normal => println!("The client is done with the connection."),
            CloseCode::Away   => println!("The client is leaving the site."),
            _ => println!("The client encountered an error: {}", reason),
        }
    }
}

/// Start a websocket listen server bound to host and port. Returns a handle for the thread this is running in.
pub fn create_server(host: Option<String>, port: i32, clients: &Arc<Mutex<BTreeMap<String, NetSock>>>, eventloop: &Arc<EventLoop>) -> thread::JoinHandle<()> {
    let mut host_bind = String::from("127.0.0.1");
    let cl = clients.clone();
    let el = eventloop.clone();
    if let Some(h) = host {
        host_bind = h;
    }
    thread::spawn(move || {
        listen(&format!("{}:{}", host_bind, port), |out| Socket { out, clients: cl.clone(), events: el.clone(), uuid: Uuid::new_v4().to_hyphenated().to_string() } ).unwrap()
    })
}
