Most of these questions are from Nick Gammon, posted on his [forums](http://www.gammon.com.au/forum/bbshowpost.php?bbsubject_id=5959)

# Terminology

* Entity: An entity is some object that exists within the engine. Some entities are visible in the game world, such as objects, agents, rooms, players, etc. Some entities are not visible in the game world and only have representation within the engine, such as timers and threads.
* EID: Entity ID. Every entity has a unique identifier attached to it. This is an internal identifier, so it is usually fairly cryptic.
* Object: An object is a physical entity that exists in the game world. This can be anything from a sword or bag up to a full-sized room.
* OID: Object ID. Every object has a unique identifier. This is an external identifier, designed to be fairly human-readable. 
* Agent: An agent is some "being" that interacts with the world. Specific examples include players and npcs.
* Region: The top-level hierarchical organizational construct for the game world. A region is a collection of coordinates on the game grid. Regions can contain other sub-regions.
* Zone: The mid-level organizational structure for the game world. A zone is a collection of rooms. Analogous to a classic mud "area". A zone can contain other sub-zones.
* Room: The lowest-level organizational structure for the game world. A container object.
* Reset: A timer that creates or destroys objects or agents at a specific interval or on a specific trigger.
* Trigger: A timer that triggers on a specific event that causes a specific action to occur.
* Timer: A timer emits an event when activated (for instance, when the countdown reaches zero). This event will generally activate a trigger.
* Event: Events are emitted by timers, triggers, player actions, etc. Many things can listen for emitted events. For example, a trap can listen for a "player enters the room" event to trigger it's effect.

# General Design

## The World

The game world will be separated into three main types of "areas": Regions, zones, and rooms. Regions will be built on a coordinate grid. Zones will be collections of rooms. Rooms will be container objects.

#### Region

A region is the top-level organizational structure, used to group areas of the world together into one logical unit. A region can contain other regions (sub-regions). A region can contain multiple zones. As an example, we can have a planet region which contains a continent region which contains a mountain range region. Regions are generally used to divide areas based on geography, culture, politics, etc. Most terrain locations will belong to multiple regions. For example, a specific location (coordinate) can belong to a terrain region (plains), a country region (Estonia), and a cultural region (German).

Regions represent stretches of terrain. Thus they represent geographical locations, as opposed to specific locations. In general, you will have large expanses of "region terrain" in between zones.

The vast majority of the time, regions will be invisible to players. They are used mainly as a building tool. If you want to do something specific, you should use a zone.

#### Zone

A zone is akin to a traditional mud area. A zone represents a specific location, such as a city, mountain path, dungeon, etc. A zone should represent an area of more detail than a region. Example zones: The Swamp Of Sorrows, The City of Barrowhaven, Shrek's Pass

Zones are collections of rooms laid out upon a grid.

#### Room

At the core, a room is basically a specialized container object. Each room can have a number of exits that can link either between rooms or between a room and a point on the main game grid (called the attachment point).

A room can be used for things such as buildings, caves, etc. A room, unlike a coordinate on the grid, can have a custom description written for it (within reason).

#### Resets / repops

The reset system will be a combination approach. There are three different types of resets:

* Static resets
** These specify agents that should always load in a specific location and will (usually) stay in that location.
** Static reset agents are loaded when the area loads and will stay around until killed or the area is unloaded.
** Static resets specify a respawn interval. After dying, the agent will respawn in the location after the interval.
** Ex: NPC shopkeeper, lair boss and minions. Can also be used to specify agents such as city guards that will walk a specific path or wander in a certain area
* Dynamic resets
** Think random encounter
** Dynamic resets generally take a list of possible agents and a percent chance to spawn for each
** Dynamic resets can be attached to a specific location, such as a room or grid location
** Dynamic resets can also be attached to an area, i.e. region, zone, or group of rooms (specified by room taxonomy system)
** Ex: Wandering monsters
* Object resets
** Spawn any game object, such as resources or items.
** Each object will have a respawn timer and respawn after the given interval.
** ==<small><em>(Should this be a separate type of reset? Should we just roll objects into static and dynamic resets?)</em></small>==

#### Geography

* Terrain taxonomy/terrain type

Ataxia will be able to support multiple planets, multiple continents, etc. Anywhere from a small 1000 room world up to 100k worlds with 10 different planets.

There will be a terrain taxonomy system.

* Weather

There will be an extensible weather system. It will be able to range anywhere from a basic rain/sunny to a large world-wide weather pattern generator that can generate multiple types of storms/weather patterns.

* Time

There will be a fairly simple time system. You will generally have simple day and night. The time speed can generally be specified per-game, even per-planet (for multi-planet games). There will, of course, be a light/dark system, where you would need lanterns/torches/magic/etc to see in the dark.

#### Scripting

We will use Lua for scripting. In general, most functionality will probably be in Lua space. The vast majority of commands will be scripts, with a few basic built-in commands. Keeping most functionality in Lua will enable us to attach specific functionality wherever we want. IE, triggered event attached to a room when someone walks in.

Lua scripts are all interpreted, but it is likely that most scripts will be loaded at boot-time, meaning they will already be in-memory. To avoid runtime errors crashing the mud, the main lua interpeter will be inside its own thread. This way we can catch run-away scripts (scripts that never return) and runtime errors.

#### Combat

Do you want fighting? If so, what are the rules? Do attacks take time (eg. a slow sword swing, or a fast dagger strike)? Can strikes be resisted, dodged, parried, evaded, and so on?

Do you have player-killing? (player vs. player). This can make your MUD more exciting, however you need to design in the rules. Can anyone attack anyone else? It is consensual? Can you only attack an opposing faction? What about NPCs (mobs)? What is to stop level 100 players wandering around killing your newbies 5 minutes after they connect to the game?

Do you have magic spells? If so, how do they work? Do they take time to cast? Can they be interrupted, resisted, repelled, reflected? What do they consume (eg. mana).

What "cooldown" do spells or abilities have? For example, a powerful spell might only be useable once every 10 minutes, thus you could say it has 10 minute cooldown.

Can you run from a battle? With what degree of success? Can you be chased? If so, how far?

Can combat occur in towns or "safe places"?

Do you have "area of effect" spells? (ie. a spell that hits all mobs, not just one)

Can characters freeze others in place (eg. to escape)?

Can characters hide (eg. to sneak past something)?

What happens if a player loses his/her connection in the middle of combat, intentionally or otherwise? Does their character keep fighting (or being attacked) until it is dead? It is removed from the game after 30 seconds?

If multiple characters are attacking one mob, which one does the mob fight back? The weakest? The first to attack it? The last to attack it? The one causing the most damage? The one who is nearest to death? The one who is healing others?

#### Races / classes / factions

Do you have different races (like orcs and goblins)? If so, what makes one different from another? What are you doing to make them balanced, so that one race is not all-powerful, or too weak?

Do you have different classes (like mages and warriors)? If so, what reason would players have for choosing one over the other? What are you doing to balance them? You probably need a paper-rock-scissors model, so that no one class (or race) can defeat, or be defeated by, all other ones.

Do you want different factions? (eg. good/evil, humans/martians, alliance/rebels). Factions can add interest to a game because they give a natural "side" that players can join. However if you have different factions then you have extra complexity (for example, deciding who fights who), and you would need cities/towns for each faction.

#### Character attributes

Most MUDs assign some sort of attributes to the characters (eg. strength, dexterity, wisdom, luck). What ones will you have? What does each one do? Try to be specific, not just "wisdom makes you smarter". For example, "each point of wisdom gives you 10 mana", or "each point of strength adds 1 extra damage to a sword-wielder".

How are the initial attributes assigned? Randomly? A table per class / race? 

How do players increase their attributes? Wearing equipment? Going up levels? Spells? Do they choose themselves what attribute is increased? (eg. each level they may be assigned 5 new attributes to distribute as they see fit).

Formulae

Work out in advance what the rules are for applying each attribute, not just vague talk about how "dexterity will be useful". For example:


Each 10 points in dexterity increases your chance to dodge a melee attack by 1%

Hit points lost in a melee attack = (attacker_attack_rating - defender_defense_rating) * weapon_damage_amount + dice_roll (2, 3)



Doing this will help focus on what attributes you want in your MUD, and what they will be used for.

#### Movement/Transport

Do you have transport systems (cars, boats, trains, aircraft, portals, teleports, horses)? If so, how do they work? What do they cost? Can anyone use them? What happens if a player logs out while on a boat?

#### Death

What happens when a character dies? Does it lose money, experience, equipment, or suffer damage? How long before it can be played again? Any penalties?

Can players resurrect other players? What are the requirements? How long does it take? Does it consume a reagent? Have a cooldown time?

#### Groups

Can players form groups? Is there a limit to a group's size? What happens if the group leader leaves? Is there a limit to level differences? (eg, no grouping level 1 players with level 100 players)?

Can players automatically follow other players?

What are the rules for looting mobs once killed (inside a group)?

What are the rules for gaining experience when a group kills a mob, or completes a quest? Do these vary if some group members are not in the same room? Or a long way away?

If a group kills a quest mob (eg. a boss) does the entire group get credit?

If a quest mob has a quest item (eg. a pendant) does the entire group get it, or only the player who looted the mob?

If a group kills a mob, but a character in the group dies half-way through the fight, does it get credit for the kill? If there is a quest item does s/he get that too?

#### Guilds

Can players form guilds (groups with some common purpose)? How? At what expense? Are there membership limits (eg. only orcs). 

What requirements are there to form a guild? Disband it? How do players get added to a guild? Removed from it? How do they remove themselves? Remove someone else? Do they have a guild meeting place?

#### Economy

Think about your game economy. MUD games tend to suffer from inflation because there is an unlimited amount of resources that can be obtained (you kill a mob, it comes back 5 minutes later, and mobs generally carry cash or valuable items). You need to design in "cash sinks" - ways that force players to spend their cash so there isn't too much of it around. For example, charge for transport, repairs, training, food, etc.

Can players trade with each other? (eg. sell a sword for 5 gold pieces). If so, how do they do that? Do they have to be in the same room or close by? If not, how would they do it? Do you have a mechanism to avoid cheating? (eg. a player pays his 5 gold but is not given his sword).

Do you want some sort of auction system? This could let players put goods up for auction, and let them sell to the highest bidder. You need to design the rules for such a thing.

How do players buy/sell goods? From shops? Wandering salesmen? Can they sell anything to anyone, or does it have to be a dealer in that product? What is the buy/sell markup? A fixed percentage or set for every object in the game? Are shops always open? Do they run out of goods? If a player sells something to a shopkeeper, is that then available for someone else to buy?

#### Quests

Do you have quests? (eg. "your task is to kill 5 kobolds", or "please take this letter to the mayor"). How are they defined (eg. in a script)? What are their rewards? Does one quest depend on another?

If you choose to have quests you will probably need to spend hours (if not weeks) designing a suitable set of quests, for each level that players will be in your game.

Can quests be shared with other players?

Who gives out quests?

Who accepts completed quests?

How do you abandon a quest? Re-gain it?

Is there a limit to the number of quests you can do?

#### Training

Do players need to train to learn things? From where? At what cost? What are the prerequisites? Can they learn everything, or do you limit them to choices (eg. you can learn 5 out of the available 10 skills)?

How are skills developed? Through practice? Trainers? Purchasing things?

#### Skills

Can players learn things (eg. metalwork)? Where do they learn? What materials are required? What cost? What can they make? What use are the made products (eg. swords could be sold). You may need to define, for each skill that players can learn, something like this:


Skill level required to make it.

Do they know how to make it? (eg. have the recipe)

What reagents are required? That is, the things that are consumed in the process (eg. leather to make shoes)

What tools are required? That is, things that are needed but are not consumed (eg. a leather-worker's needle).

Do they need to be at a certain place? Eg. a forge for metalwork, a fire for cooking.


#### Objects

Most MUDs have objects - that is, things that can be picked up, carried around and used. What are you planning to have? What will they do? Where do players get them? What use are they? Can they be bought or sold? Can they be made?

Does equipment wear out? To what rules? How is it repaired? At what cost? Where?

Can objects be discarded (left lying on the ground), or destroyed?

How many objects can a player carry at one time? Can they purchase or obtain extra capacity? Is capacity based on weight, number, volume?

What attributes do objects have? An object might have:


Armour class (wearing it increases armour)

Attack rating (wielding it gives X amount of attack, eg. a sword)

Attribute modification (eg. wearing it increases wisdom)

A use - eg. using a key unlocks a door, eating food increases health

Value - how much you can sell it for

Wear and tear - how much damage it has suffered

Minimum level number to use it

Restrictions on use - eg. only mages might use wands

Can a player carry more than one of it? If so, is there a limit?

Is it a quest item? If so, for which quest(s)?

Does it disappear if the player disconnects?

Can it be sold? Traded?


#### Mail

Do you have an in-game mail system? If so, what can be sent? Messages? Objects? Money?

Where are mailboxes? Does it cost to use them?

How long does mail take to be delivered?

#### Bank

Do you have a "bank" or similar storage place where players can deposit things that they cannot fit into their inventory?

How much can they store in the bank? Where is it/them? How much does it cost to use it?

#### Logging

Do you log everything that happens? Eg. what players say, who they kill, what trades are occurring?

In the event of a dispute, can the log be easily perused?

#### Levels

Do players have levels? (eg. you start at level 1, and hope to reach level 100). How many? How do you advance levels? What advantages are there to being a higher level?

What do players do when they reach the highest level?

#### Making money

Do you - the developers of the game - plan to make (real) money from it? Are you going to:


Charge by the hour to play?

Have a monthly subscription?

Have a free trial period?

Sell in-game goods or services for real money?

Sell the game itself (to other aspiring MUD admins)?


Who collects the money, and how?

#### Ownership

Who "owns" the game? If you are developing it from scratch, using a team of people, what happens if one leaves? Who (if anyone) owns:


The name of the game?

The computer program?

Any scripts used in a scripting language?

The room design (eg. room descriptions)

Other designs (eg. quests, object design)


Can someone leave your development team, take a copy of everything, and start a competing game?

#### Running the game

What server are you going to run it on? A commercial server or the private one belonging to one of the developers?

How are you going to administer it 24 hours a day / 7 days a week / 365 days a year? What happens if the chief (or only) programmer goes on holiday for a month and it crashes while s/he is gone?

Are you planning to have customer service people online all the time to help players? Will they be paid? How much?

What will you do to help players that are stuck (eg. in a room with no exits), or have lost all their goods through some bug (or their own fault)?

What method do players have to report bugs (eg. misspellings, or broken quests)?

#### Handling problem players

Eventually you will have a "problem" player. How do you deal with him/her? They may do things like:


Have an unsuitable name, like a swear-word, or a name that implies they are an admin.

Use abusive language.

Harrass or irritate other players.

Make life hard for other players by killing mobs that they are trying to kill or otherwise disrupting their quests.

Cheat, eg. by promising to pay for something and then not pay.

Exploit game mechanics or bugs to become super-powerful.

Camp out valuable resources or important mobs, getting to them before anyone else can.



You probably need some "self-help" mechanism for other players (eg. being able to ignore an annoying player), with a fall-back of being able to get help from a customer service person in extreme caes.

#### Accounts

Do players create "accounts"? That is, a central log-in point for all of their characters? Or is each character treated on its own? Is there a limit to how many characters one player can have?

Do you allow multi-playing - that is, one person controlling multiple characters simultaneously? It may be hard to stop simply based on IP address these days, as many IP addresses are shared by routers using Network Address Translation (NAT).

Do you "lock out" accounts for some reason (eg. non-payment, misbehaviour)?

#### Newbies

Every game will start of with "newbies" - even if they are experienced MUD players they will be new to your game. What do you do to help them? You might have help files, easy quests, a more experienced player assigned to help them, "newbie" chat channels, and so on.

What do new players get to choose when they start playing? Character name / gender / race / class / description?


#### Chatting

Players expect to talk to each other - do you have multiple chat channels? You probably don't want level 100 players having to put up with newbie questions in the middle of a difficult quest. However newbies will need help too. Do you "zone" your chats, so the chats are relevant to where people are (eg. zoned to the current town)?

#### Mob / NPC artificial intelligence

AI (artificial intelligence) is the programming that controls your mobs / NPCs (Non Player Characters) so they behave in a reasonably interesting way. How do you plan to have this work? Do all mobs behave the same, or do humanoids behave distinctly differently to (say) spiders?

If NPCs are spell-casters, what spells do they choose to cast?

Do NPCs run away if facing an overwhelming foe? If so, how far?

Do NPCs go and get help?

Inside what range do aggressive NPCs attack?

If the player runs away, does the NPC chase it, and if so, for how long?

#### Instanced dungeons

Do you plan to have "instanced dungeons"? These are areas of the game where players (or groups) get their own copy of a particular area. The point of this is to allow a group to play on its own, and not find that someone else has just killed the boss, or is harrassing the group. If you do plan on this, then it impacts on game design, because you may need to record an instance number alongside a player's location (eg. player Nick is in room 1002, instance 3).

#### Pets

Can players have pets? (eg. a pet dog) Is it for combat or local flavour? If for combat, how is it controlled? What happens when it dies? How is it resurrected / replaced? Does owning a pet cost anything (to purchase, or for upkeep)?

#### Game balance

What can you do to stop high-level players helping their friends with inappropriate equipment? (eg. level restrictions on good gear).

Is each class / race / faction balanced? You can be sure that if players work out that "orc warriors" are much more powerful than everything else, then soon the game will be filled with orc warriors.

Are attributes balanced? If one attribute (eg. strength) has a disproportional affect on combat, then every player will put all their points in strength (if they can).

#### Making characters different

You will probably want some sort of system that makes players choose between mutually exclusive features for their characters, so that you don't have, after 3 months, a MUD filled with dozens of characters who are all the same because they have all "maxed out" on every possible attribute.

Examples are:


Choosing different classes / races, where each one offers something the others don't

Choosing different professions, which offer different things in the game (eg. make armor OR make potions, but not both)

Wear different clothing, where a particular piece of armour can add protection or offensive power, but not both

Give players optional points to choose from, where they have to choose 20 out of a possible 50, so it is unlikely that every player will choose the same 20. Points might be spent in offense, defense, stealth, long life, short battles, and so on.



#### Saving things

Players expect their characters to be there next time they connect. When is game data saved? Every minute? Every time a player changes rooms? Do you save to lots of small files (eg. one per player, like SMAUG does), or one big file (like PennMUSH does).

Do you save to flat files or an SQL (or similar) database?

What is saved? If a player moves an object from one room to another does the game remember that?

If the server crashes and is restarted, what state is everything in? Are all mobs repopulated into their default positions? Or, does the game remember where they all were 5 minutes ago?

#### Building

How do builders extend the game? Dynamically while players are online? Offline by editing area files?

Do you have some sort of "builder interface" used only by builders to extend the game while it is running?

Do you provide some sort of web-page interface for builders?

Does the game need to be restarted to incorporate additional content?

# Specific Design

## The World

### Geography