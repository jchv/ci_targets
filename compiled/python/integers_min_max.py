# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

from pkg_resources import parse_version
import kaitaistruct
from kaitaistruct import KaitaiStruct, KaitaiStream, BytesIO


if parse_version(kaitaistruct.__version__) < parse_version('0.9'):
    raise Exception("Incompatible Kaitai Struct Python API: 0.9 or later is required, but you have %s" % (kaitaistruct.__version__))

class IntegersMinMax(KaitaiStruct):
    def __init__(self, _io, _parent=None, _root=None):
        self._io = _io
        self._parent = _parent
        self._root = _root if _root else self
        self._read()

    def _read(self):
        self.unsigned_min = self._root.Unsigned(self._io, self, self._root)
        self.unsigned_max = self._root.Unsigned(self._io, self, self._root)
        self.signed_min = self._root.Signed(self._io, self, self._root)
        self.signed_max = self._root.Signed(self._io, self, self._root)

    class Unsigned(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.u1 = self._io.read_u1()
            self.u2le = self._io.read_u2le()
            self.u4le = self._io.read_u4le()
            self.u8le = self._io.read_u8le()
            self.u2be = self._io.read_u2be()
            self.u4be = self._io.read_u4be()
            self.u8be = self._io.read_u8be()


    class Signed(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.s1 = self._io.read_s1()
            self.s2le = self._io.read_s2le()
            self.s4le = self._io.read_s4le()
            self.s8le = self._io.read_s8le()
            self.s2be = self._io.read_s2be()
            self.s4be = self._io.read_s4be()
            self.s8be = self._io.read_s8be()



