Room = DataAccessor:new({
	_typestr = "Room",

	-- strings
	name = function(self, value) return self:access("string", "name", value) end,
	description = function(self, value) return self:access("string", "description", value) end,

	-- objects
	exits = function(self, key, value) return self:accessDict(RoomExit, "exits", key, value) end,

	-- methods
	Send = function(self, value) SendToRoom(self._id, value) end
})

function Room:create(id, context)
	return Room:new{_id = id, _context = context}
end


RoomExit = DataAccessor:new({
	_typestr = "RoomExit",

	destination = function(self, value) return self:access(Room, "destination", value) end
})

function RoomExit:create(id, context)
	return RoomExit:new{_id = id, _context = context}
end
