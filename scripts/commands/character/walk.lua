function trim(s)
	return s:match'^%s*(.*%S)' or ''
end

function starts(String, Start)
	return string.sub(String,1,string.len(Start))==Start
end

function do_walk (ch, args)
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
		ch:Send("What direction do you want to walk?\n")
		return
	end

	from_room = ch:room()
	exit = from_room:exits(tostring(dir))

	if exit == nil then
		ch:Send("There is no exit in that direction.\n")
		return
	end

	to_room = exit:destination()
	if to_room == nil then
		ch:Send("That direction appears to go nowhere.\n")
		return
	end

	ch:Send("You walk " .. dirs[dir+1] .. ".\n")

	ch:room(nil)
--	from_room:Send(ch:name() .. " leaves " .. dirs[dir+1] .. ".\n")
--	to_room:Send(ch:name() .. " has arrived from the " .. rdir[dir+1] .. ".\n")
	ch:room(to_room)
	do_look(ch, "")
end

function do_north (ch, args)
	do_walk(ch, "north")
end

function do_south (ch, args)
	do_walk(ch, "south")
end

function do_east (ch, args)
	do_walk(ch, "east")
end

function do_west (ch, args)
	do_walk(ch, "west")
end

function do_up (ch, args)
	do_walk(ch, "up")
end

function do_down (ch, args)
	do_walk(ch, "down")
end
