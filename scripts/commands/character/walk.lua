function trim(s)
	return s:match'^%s*(.*%S)' or ''
end

function starts(String, Start)
	return string.sub(String,1,string.len(Start))==Start
end

function do_walk (char_id, args)
	dirs = {"north", "east", "south", "west", "up", "down"}
	rdir = {"south", "west", "north", "east", "down", "up"}
	dir = nil

	for d = 0, 5 do
		if starts(dirs[d+1], trim(args)) then
			dir = d
			break
		end
	end

	if dir == nil then
		SendToChar(char_id, "What direction do you want to walk?\n")
		return
	end

	room_id = GetCharacterData(char_id, "room")
	exit_id = GetRoomExit(room_id, dir)

	if exit_id == "" then
		SendToChar(char_id, "There is no exit in that direction.\n")
		return
	end

	dest_id = GetRoomExitData(exit_id, "destination")

	if dest_id == "" then
		SendToChar(char_id, "That direction appears to go nowhere.\n")
		return
	end

	name = GetCharacterData(char_id, "name")
	SendToChar(char_id, "You walk " .. dirs[dir+1] .. ".\n")

	SetCharacterData(char_id, "room", "")
--	SendToRoom(room_id, name .. " leaves " .. dirs[dir+1] .. ".\n")
--	SendToRoom(dest_id, name .. " has arrived from the " .. rdir[dir+1] .. ".\n")
	SetCharacterData(char_id, "room", dest_id)

	do_look(char_id, "")
end

function do_north (char_id, args)
	do_walk(char_id, "north")
end

function do_south (char_id, args)
	do_walk(char_id, "south")
end

function do_east (char_id, args)
	do_walk(char_id, "east")
end

function do_west (char_id, args)
	do_walk(char_id, "west")
end

function do_up (char_id, args)
	do_walk(char_id, "up")
end

function do_down (char_id, args)
	do_walk(char_id, "down")
end
