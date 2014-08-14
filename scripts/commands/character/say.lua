function do_say (char_id, args)
	name = GetCharacterData(char_id, "name")
	SendToOthers(char_id, string.format("%s says '%s'\n", name, args))
	SendToChar(char_id, string.format("You say '%s'\n", args))
end
