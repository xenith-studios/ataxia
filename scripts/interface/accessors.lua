DataAccessor = {
	_id = nil,
	_context = nil
}

function DataAccessor:new(obj)
	assert(obj ~= nil, "attempt to instantiate abstract class DataAccessor")
	setmetatable(obj, self)
	self.__index = self
	return obj
end

function DataAccessor:accessString(field_name, value)
	assert(field_name ~= nil, "call to accessString with no field name")

	if value ~= nil then
		self._context:set(self._id, field_name, value)
		SetObjectData(self._id, field_name, value) -- remove later for delayed commits
	else
		value = self._context:get(self._id, field_name)

		if value == nil then
			value = assert(GetObjectData(self._id, field_name), "call to access with nonsense field name")
			self._context:set(self._id, field_name, value)
		end
	end

	return value
end

function DataAccessor:accessObject(obj_type, field_name, value)
	assert(field_name ~= nil, "call to accessObject with no field name")

	if value ~= nil then
		self._context:set(self._id, field_name, value)
		SetObjectData(self._id, field_name, value._id) -- remove later for delayed commits
	else
		value = self._context:get(self._id, field_name)

		if value == nil then
			id = self:accessString(field_name)
			value = obj_type:create(id, self._context)
			self._context:set(self._id, field_name, value)
		end
	end

	return value
end

function DataAccessor:accessDict(obj_type, field_name, key, value)
	assert(field_name ~= nil, "call to accessDict with no field name")
	assert(key ~= nil, "call to accessDict with no key")
	assert(obj_type ~= nil, "call to accessDict with no object type")

	if value ~= nil then
		SetDictData(self._id, field_name, key, value)
	else
		-- fetch the id from the dict
		id = GetDictData(self._id, field_name, key) -- annoyingly returns empty strings if not found
		if id ~= nil and id ~= "" then
			return obj_type:create(id, self._context)
		end
	end

	return nil
end
