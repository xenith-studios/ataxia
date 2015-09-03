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
	Keywords         string `json:"keywords"`
	ShortDescription string `json:"short_descr"`
	LongDescription  string `json:"long_descr"`
	Description      string `json:"description"`
	Race             string `json:"race"`
	ActFlags         string `json:"act_flags"`
	AffFlags         string `json:"aff_flags"`
	Alignment        int    `json:"alignment"`
	Group            string `json:"group"`
	Level            int    `json:"level"`
	Hitroll          int    `json:"hitroll"`
	HPDice           string `json:"hp_dice"`
	ManaDice         string `json:"mana_dice"`
	DamageDice       string `json:"damage_dice"`
	DamageType       string `json:"damage_type"`
	ACPierce         int    `json:"ac_pierce"`
	ACBash           int    `json:"ac_bash"`
	ACSlash          int    `json:"ac_slash"`
	ACExotic         int    `json:"ac_exotic"`
	OffFlags         string `json:"off_flags"`
	ImmFlags         string `json:"imm_flags"`
	ResFlags         string `json:"res_flags"`
	VulnFlags        string `json:"vuln_flags"`
	StartPos         string `json:"start_pos"`
	DefaultPos       string `json:"default_pos"`
	Sex              string `json:"sex"`
	Wealth           int    `json:"wealth"`
	FormFlags        string `json:"form_flags"`
	PartFlags        string `json:"part_flags"`
	Size             string `json:"size"`
	Material         string `json:"material"`
	// remove_flags (hack)
	// mobprogs
}

// ObjectTemplate is the template data structure for objects
type ObjectTemplate struct {
	Keywords         string `json:"keywords"`
	ShortDescription string `json:"short_descr"`
	Description      string `json:"description"`
	Material         string `json:"material"`
	ItemType         string `json:"item_type"`
	ExtraFlags       string `json:"extra_flags"`
	WearFlags        string `json:"wear_flags"`
	Value0           string `json:"value0"`
	Value1           string `json:"value1"`
	Value2           string `json:"value2"`
	Value3           string `json:"value3"`
	Value4           string `json:"value4"`
	Level            int    `json:"level"`
	Weight           int    `json:"weight"`
	Cost             int    `json:"cost"`
	Condition        string `json:"condition"`
	//	added_affects	[]map[string]int
	//	added_flags		[]map[string]int (more complex, needs struct)
	ExtraDescription map[string]string `json:"extra_descr"`
}

// RoomTemplate is the template data structure for rooms
type RoomTemplate struct {
	Name             string                      `json:"name"`
	Description      string                      `json:"description"`
	TeleDest         int                         `json:"tele_dest"`
	RoomFlags        string                      `json:"room_flags"`
	SectorType       int                         `json:"sector_type"`
	HealRate         int                         `json:"heal_rate"`
	ManaRate         int                         `json:"mana_rate"`
	Clan             string                      `json:"Clan"`
	Guild            string                      `json:"Guild"`
	Owner            string                      `json:"Owner"`
	Exits            map[string]RoomExitTemplate `json:"Exits"`
	ExtraDescription map[string]string           `json:"extra_descr"`
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
	ID          string `json:"id"`
	DestVnum    string `json:"dest_vnum"`
	Destination *Room  `json:"destination"`
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
		for dirStr, exitTemplate := range roomTemplate.Exits {
			dir, _ := strconv.Atoi(dirStr)
			exit := NewRoomExit()
			exit.DestVnum = strconv.Itoa(exitTemplate.Vnum)
			room.exits[dir] = exit
		}

		area.rooms[vnum] = room
		area.World.AddRoom(room)
	}

	// resolve exits to room pointers (for now, this is only intra-area)
	for _, room := range area.rooms {
		for dir, exit := range room.exits {
			dest := area.World.LookupRoom(exit.DestVnum)
			if dest == nil {
				log.Println("Couldn't find room destination for vnum", exit.DestVnum)
				delete(room.exits, dir)
			} else {
				exit.Destination = dest
				area.World.AddRoomExit(exit)
			}
		}
	}
}
