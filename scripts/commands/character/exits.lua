function do_exits (ch, args)
	dirs = {"north", "east", "south", "west", "up", "down"}
	exit_list = "[ "
	found = false

	for dir = 0, 5 do
		exit = ch:room():exits(tostring(dir))
		if exit ~= nil then
			exit_list = exit_list .. dirs[dir+1] .. " "
			found = true
		end
	end

	if found == false then
		exit_list = exit_list .. "none "
	end

	exit_list = exit_list .. "]"
	ch:Send(exit_list .. "\n")
end
