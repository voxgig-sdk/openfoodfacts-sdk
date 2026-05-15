-- Openfoodfacts SDK error

local OpenfoodfactsError = {}
OpenfoodfactsError.__index = OpenfoodfactsError


function OpenfoodfactsError.new(code, msg, ctx)
  local self = setmetatable({}, OpenfoodfactsError)
  self.is_sdk_error = true
  self.sdk = "Openfoodfacts"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function OpenfoodfactsError:error()
  return self.msg
end


function OpenfoodfactsError:__tostring()
  return self.msg
end


return OpenfoodfactsError
