-- This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild
--
-- This file is compatible with Lua 5.3

local class = require("class")
require("kaitaistruct")
require("custom_fx_no_args")

ProcessCustomNoArgs = class.class(KaitaiStruct)

function ProcessCustomNoArgs:_init(io, parent, root)
  KaitaiStruct._init(self, io)
  self._parent = parent
  self._root = root or self
  self:_read()
end

function ProcessCustomNoArgs:_read()
  self._raw_buf = self._io:read_bytes(5)
  local _process__raw_buf = CustomFxNoArgs()
  self.buf = _process__raw_buf:decode(self._raw_buf)
end


