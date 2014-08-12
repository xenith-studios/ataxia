function do_say (player_id, args)
	name = GetPlayerData(player_id, "name")
	SendToAll(string.format("%s says '%s'", name, args))
end
