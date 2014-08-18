function do_look (ch, args)
	room_name = ch:room():name()
	room_desc = ch:room():description()
	ch:Send(room_name.."\n"..room_desc.."\n")
	do_exits(ch, "")
end

