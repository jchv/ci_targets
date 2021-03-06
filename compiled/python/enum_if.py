# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

from pkg_resources import parse_version
import kaitaistruct
from kaitaistruct import KaitaiStruct, KaitaiStream, BytesIO
from enum import Enum


if parse_version(kaitaistruct.__version__) < parse_version('0.9'):
    raise Exception("Incompatible Kaitai Struct Python API: 0.9 or later is required, but you have %s" % (kaitaistruct.__version__))

class EnumIf(KaitaiStruct):

    class Opcodes(Enum):
        a_string = 83
        a_tuple = 84
    def __init__(self, _io, _parent=None, _root=None):
        self._io = _io
        self._parent = _parent
        self._root = _root if _root else self
        self._read()

    def _read(self):
        self.op1 = EnumIf.Operation(self._io, self, self._root)
        self.op2 = EnumIf.Operation(self._io, self, self._root)
        self.op3 = EnumIf.Operation(self._io, self, self._root)

    class Operation(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.opcode = KaitaiStream.resolve_enum(EnumIf.Opcodes, self._io.read_u1())
            if self.opcode == EnumIf.Opcodes.a_tuple:
                self.arg_tuple = EnumIf.ArgTuple(self._io, self, self._root)

            if self.opcode == EnumIf.Opcodes.a_string:
                self.arg_str = EnumIf.ArgStr(self._io, self, self._root)



    class ArgTuple(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.num1 = self._io.read_u1()
            self.num2 = self._io.read_u1()


    class ArgStr(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.len = self._io.read_u1()
            self.str = (self._io.read_bytes(self.len)).decode(u"UTF-8")



