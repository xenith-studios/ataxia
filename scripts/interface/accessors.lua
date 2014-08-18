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

function DataAccessor:accessPrimitive(datatype, field_name, value)
	if value ~= nil then
		assert(type(value) == datatype, "attempt to set improper primitive datatype")
		self._context:set(self._id, field_name, value)
		SetObjectData(self._id, field_name, value) -- remove later for delayed commits
		return
	end

	ret = self._context:get(self._id, field_name)

	if ret == nil then
		ret = assert(GetObjectData(self._id, field_name), "call to access with nonsense field name")
		self._context:set(self._id, field_name, ret)
	end

	assert(type(ret) == datatype, "access failed to return proper primitive datatype")
	return ret
end

function DataAccessor:access(datatype, field_name, value)
	assert(datatype ~= nil, "call to access with no object type")
	assert(field_name ~= nil, "call to access with no field name")

	if datatype == "number" or datatype == "string" or datatype == "boolean" then
		return self:accessPrimitive(datatype, field_name, value)
	end

	if value ~= nil then
		assert(datatype._typestr == value._typestr, "attempt to set improper datatype")
		self._context:set(self._id, field_name, value)
		SetObjectData(self._id, field_name, value._id) -- remove later for delayed commits
		return
	end

	ret = self._context:get(self._id, field_name)

	if ret == nil then
		id = self:accessPrimitive("string", field_name)
		ret = datatype:create(id, self._context)
		self._context:set(self._id, field_name, ret)
	end

	assert(datatype._typestr == ret._typestr, "access failed to return proper datatype")
	return ret
end

function DataAccessor:accessDict(datatype, field_name, key, value)
	assert(datatype ~= nil, "call to accessDict with no object type")
	assert(field_name ~= nil, "call to accessDict with no field name")
	assert(key ~= nil, "call to accessDict with no key")

	if value ~= nil then
		SetDictData(self._id, field_name, key, value)
	else
		-- fetch the id from the dict
		id = GetDictData(self._id, field_name, key) -- annoyingly returns empty strings if not found
		if id ~= nil and id ~= "" then
			return datatype:create(id, self._context)
		end
	end

	return nil
end
