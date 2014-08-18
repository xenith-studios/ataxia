Room = DataAccessor:new({
	-- strings
	name = function(self, value) return self:accessString("name", value) end,
	description = function(self, value) return self:accessString("description", value) end,

	-- objects
	exits = function(self, key, value) return self:accessDict(RoomExit, "exits", key, value) end,

	-- methods
	Send = function(self, value) SendToRoom(self._id, value) end
})

function Room:create(id, context)
	return Room:new{_id = id, _context = context}
end


RoomExit = DataAccessor:new({
	destination = function(self, value) return self:accessObject(Room, "destination", value) end
})

function RoomExit:create(id, context)
	return RoomExit:new{_id = id, _context = context}
end
