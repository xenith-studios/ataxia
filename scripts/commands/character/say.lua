function do_say (ch, args)
	SendToOthers(ch._id, string.format("%s says '%s'\n", ch:name(), args))
	ch:Send(string.format("You say '%s'\n", args))
end
