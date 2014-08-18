function execute_character_action (char_id, func_name, args)
	ch = Character:create(char_id, Context:new())
	_G[func_name](ch, args)
end
