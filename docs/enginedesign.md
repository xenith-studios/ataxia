### Networking notes

How to listen for incoming connections and spawn them onto the executor as futures:

```
let server = TcpStream::new()
    .incoming()
    .map_err(|e| println!("error = {:?}", e))
    .for_each(|s| { 
        tokio::spawn(TelnetHandler::new(s))
    })
```
TelnetHandler (and WebSocketHandler) will be a future that handles a single connection:
* Reading input
* Parsing input
* Send parsed/formatted input on channel to be sent to engine via mq
* Receive output from engine via channel
* Formatting output
* Writing output

There will be another future that handles shuffling messages back and forth via mq to the engine. 
It receives input from all client futures via mpsc channel, sends them out the mq to the engine.
It receives output from the engine via mq, then writes it to the proper mpsc channel to the correct client future.

* Task (main future) is the listener
* tokio::run the task, blocks thread until listener shuts down
* It spawns a task for the mq future
* It spawns new tasks (futures) for each client connection
* TelnetHandler/WebSocketHandler implements Future