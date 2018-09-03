## Random Ideas
* Database for storage
  * SQLite (or PostgreSQL)
  * Maybe a K/V store of some kind?
* Area files
  * Areas defined in YAML (or similar markup language)
  * NPC templates defined in YAML
* Lua for scripting (or maybe [dyon](https://github.com/pistondevelopers/dyon) or [gluon](https://github.com/gluon-lang/gluon)?)
  * One Lua_State per thread
  * Lua bridges
    * [rlua](https://github.com/kyren/rlua)
* Multithreaded
  * Main thread for engine (for handling parsed game input, game state, npc updates, etc) and time (pulses/ticks)
  * Network thread for handing player communication and [gossip](https://gossip.haus) (Tokio event loop), might be implemented as a threadpool?
  * "Game" thread for non-time sensitive background events (weather, in-game time, etc)
  * Where to put slow, blocking actions (DNS calls, etc)?
* Variables spaces
  * Game data in Rust data structures
  * Player and world data (rooms, npcs) in Lua data structures?
    * Rust handles saving/loading data and converts into lua data structures
* Proxy process
  * All players connect to proxy, proxy manages player communications.
  * Proxy communicates to engine process via some data protocol (gRPC/protobuf, capnproto, thrift, etc)
  * This allows engine to only need to implement a single advanced protocol (https://github.com/mudcoders/ump can be an example)
  * Proxy can implement telnet, web sockets, etc. for different types of clients
  * This allows us to greatly simplify hotboot/copyover process. Engine doesn't need to worry about maintaining sockets anymore, the engine just reconnects to the proxy after it restarts. Don't even necessarily need the engine to exec itself, can just let process shutdown and reboot (let the startup/supervisor script handle that!)
  * Examples: https://github.com/evennia/evennia/wiki/Portal-And-Server and http://evennia.blogspot.com/2018/01/kicking-into-gear-from-distance.html

## Object hierarchy notes
* Entity-component system (ECS)
* Template objects/archetypes
* Objects all defined in scripts? or YAML?

## NPC notes
* Template-based system
  * Forms (animal, humanoid, plant, skeleton, etc)?
  * Example: npc object contains pointer to species object which defines species data (name, immunities, etc)
* NPC taxonomy

## Area/room notes
* Still many questions to be answered here:
* Regions are top-level
  * Regions can have subregions (mountain range region encased in continental region encased in planetary region)
* Zones belong to regions
  * "Traditional area"
  * City, dungeon, mountain path
  * Generally used for a higher-level of detail than region
    * Ex: Mountain region has a path zone through it (if path is important)
* Rooms are container objects
  * Generally placed as objects on the coordinate grid inside of a zone
  * Ex: building inside a city
* Mainly coordinate-based roomless system
  * Rooms can be placed in sequence inside zone for dungeon effect, if necessary
  * Room taxonomy? ("Terrain type")

## General functionality
* System of event hooks (before_event, on_event, after_event)
  * Sent to every object when events happen
  * Objects knows how to handle own events
  * Objects ignore events they can't handle
* Event-driven system?

## Commands
* Every command is defined as a separate lua script
  * Some core commands are built-in (like shutdown, reboot, load script)
* Command prefixes
  * Admin commands prefixed with /

## Settings/lore notes
* Modern fantasy
* Low magic (start off as no magic, must rediscover magic?)
  * Magic is not flashy, very subtle, but can be powerful (like LOTR or Dresden)
* Ages
  * Age of Creation
  * Age of Strife(?)
  * Age of Legends
  * Age of Legacy <-- game starts here?
* Big cities are bastions of civilization (safe zones), the farther you get from the cities the more dangerous the world gets.
