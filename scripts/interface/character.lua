
Character = DataAccessor:new({
	-- strings
	name = function(self, value) return self:accessString("name", value) end,
	-- objects - keep our own copy of everything, for future change commits.  only accept
	-- an ID as a value, look it up ourselves and make a new object
	room = function(self, value) return self:accessObject(Room, "room", value) end,
	-- methods
	Send = function(self, value) SendToChar(self._id, value) end
})

function Character:create(id, context)
	return Character:new{_id = id, _context = context}
end
