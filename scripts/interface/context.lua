Context = {
	accessed = {} -- id, field_name, object
}

function Context:new(obj)
	obj = obj or {}
	setmetatable(obj, self)
	self.__index = self
	return obj
end

function Context:get(id, field_name)
	if self.accessed[id] == nil then
--		print(id, field_name, "-->", "nil id")
		return nil
	end

--	print(id, field_name, "-->", self.accessed[id][field_name])
	return self.accessed[id][field_name]
end

function Context:set(id, field_name, value)
--	print(id, field_name, "<--", value)

	if self.accessed[id] == nil then
		self.accessed[id] = {}
	end

	self.accessed[id][field_name] = value
end
