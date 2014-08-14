function do_look (char_id, args)
	room_id = GetCharacterData(char_id, "room")
	room_name = GetRoomData(room_id, "name")
	room_desc = GetRoomData(room_id, "description")
	SendToChar(char_id, room_name.."\n"..room_desc.."\n")
	do_exits(char_id, "")
end
