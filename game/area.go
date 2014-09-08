package game

import (
	"encoding/json"
	//	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/xenith-studios/ataxia/utils"
)

// MobileTemplate is the template data structure for mobs
type MobileTemplate struct {
	keywords    string
	short_descr string
	long_descr  string
	description string
	race        string
	act_flags   string
	aff_flags   string
	alignment   int
	group       string
	level       int
	hitroll     int
	hp_dice     string
	mana_dice   string
	damage_dice string
	damage_type string
	ac_pierce   int
	ac_bash     int
	ac_slash    int
	ac_exotic   int
	off_flags   string
	imm_flags   string
	res_flags   string
	vuln_flags  string
	start_pos   string
	default_pos string
	sex         string
	wealth      int
	form_flags  string
	part_flags  string
	size        string
	material    string
	// remove_flags (hack)
	// mobprogs
}

// ObjectTemplate is the template data structure for objects
type ObjectTemplate struct {
	keywords    string
	short_descr string
	description string
	material    string
	item_type   string
	extra_flags string
	wear_flags  string
	value0      string
	value1      string
	value2      string
	value3      string
	value4      string
	level       int
	weight      int
	cost        int
	condition   string
	//	added_affects	[]map[string]int
	//	added_flags		[]map[string]int (more complex, needs struct)
	extra_descr map[string]string
}

// RoomTemplate is the template data structure for rooms
type RoomTemplate struct {
	Name        string
	Description string
	Tele_dest   int
	Room_flags  string
	Sector_type int
	Heal_rate   int
	Mana_rate   int
	Clan        string
	Guild       string
	Owner       string
	Exits       map[string]RoomExitTemplate
	Extra_descr map[string]string
}

// RoomExitTemplate is the template data structure for room exits
type RoomExitTemplate struct {
	Description string
	Keywords    string
	Locks       int
	Key         int
	Vnum        int
}

// Room is a single room
type Room struct {
	ID          string
	Vnum        string
	Name        string
	Description string
	exits       map[int]*RoomExit
}

// NewRoom returns a new room
func NewRoom() *Room {
	return &Room{
		ID:    utils.UUID(),
		exits: make(map[int]*RoomExit),
	}
}

// RoomExit is a single room exit
type RoomExit struct {
	ID          string
	dest_vnum   string
	destination *Room
}

// NewRoomExit returns a new room exit
func NewRoomExit() *RoomExit {
	return &RoomExit{
		ID: utils.UUID(),
	}
}

// AreaHeader ##TODO
type AreaHeader struct {
	Credits  string
	Name     string
	Filename string
}

// AreaPrototype ##TODO
type AreaPrototype struct {
	Area          AreaHeader              `json:"AREA"`
	RoomTemplates map[string]RoomTemplate `json:"ROOMS"`
	//	mobileTemplates	map[string]MobileTemplate 	`json:"MOBILES"`
	//	objectTemplates	map[string]ObjectTemplate 	`json:"OBJECTS"`
	// Resets	[]ResetTemplate
	//	roomTemplates	map[string]RoomTemplate 	`json:"ROOMS"`
	// shops
	// specials
}

// Area is a single area
type Area struct {
	ID        string
	World     *World
	Prototype AreaPrototype
	rooms     map[string]*Room
}

// NewArea returns a new area
func NewArea(world *World) *Area {
	return &Area{
		ID:    utils.UUID(),
		World: world,
		rooms: make(map[string]*Room),
	}
}

// Load an area from a file
func (area *Area) Load(filename string) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal("Unable to read area file", filename)
	}

	log.Println("Loaded file", filename)
	err = json.Unmarshal(bytes, &area.Prototype)
	//	fmt.Printf("%+v\n", area.Prototype)

	if err != nil {
		log.Fatal("Unable to parse area file", filename)
	}

	log.Println("Loaded area", area.Prototype.Area.Name)
}

// Initialize a new area
func (area *Area) Initialize() {
	log.Println("Initializing area", area.Prototype.Area.Name)
	// make instance of each room, add the exits
	for vnum, roomTemplate := range area.Prototype.RoomTemplates {
		room := NewRoom()
		room.Vnum = vnum
		room.Name = roomTemplate.Name
		room.Description = roomTemplate.Description
		for dir_str, exitTemplate := range roomTemplate.Exits {
			dir, _ := strconv.Atoi(dir_str)
			exit := NewRoomExit()
			exit.dest_vnum = strconv.Itoa(exitTemplate.Vnum)
			room.exits[dir] = exit
		}

		area.rooms[vnum] = room
		area.World.AddRoom(room)
	}

	// resolve exits to room pointers (for now, this is only intra-area)
	for _, room := range area.rooms {
		for dir, exit := range room.exits {
			dest := area.World.LookupRoom(exit.dest_vnum)
			if dest == nil {
				log.Println("Couldn't find room destination for vnum", exit.dest_vnum)
				delete(room.exits, dir)
			} else {
				exit.destination = dest
				area.World.AddRoomExit(exit)
			}
		}
	}
}
