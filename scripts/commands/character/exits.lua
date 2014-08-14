function do_exits (char_id, args)
	room_id = GetCharacterData(char_id, "room")
	dirs = {"north", "east", "south", "west", "up", "down"}
	exit_list = "[ "
	found = false

	for dir = 0, 5 do
		exit_id = GetRoomExit(room_id, dir)
		if exit_id ~= "" then
			exit_list = exit_list .. dirs[dir+1] .. " "
			found = true
		end
	end

	if found == false then
		exit_list = exit_list .. "none "
	end

	exit_list = exit_list .. "]"
	SendToChar(char_id, exit_list .. "\n")
end
