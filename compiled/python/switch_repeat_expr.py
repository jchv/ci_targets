# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

from pkg_resources import parse_version
from kaitaistruct import __version__ as ks_version, KaitaiStruct, KaitaiStream, BytesIO


if parse_version(ks_version) < parse_version('0.7'):
    raise Exception("Incompatible Kaitai Struct Python API: 0.7 or later is required, but you have %s" % (ks_version))

class SwitchRepeatExpr(KaitaiStruct):
    def __init__(self, _io, _parent=None, _root=None):
        self._io = _io
        self._parent = _parent
        self._root = _root if _root else self
        self._read()

    def _read(self):
        self.code = self._io.read_u1()
        self.size = self._io.read_u4le()
        self._raw_body = [None] * (1)
        self.body = [None] * (1)
        for i in range(1):
            _on = self.code
            if _on == 17:
                self._raw_body[i] = self._io.read_bytes(self.size)
                io = KaitaiStream(BytesIO(self._raw_body[i]))
                self.body[i] = self._root.One(io, self, self._root)
            elif _on == 34:
                self._raw_body[i] = self._io.read_bytes(self.size)
                io = KaitaiStream(BytesIO(self._raw_body[i]))
                self.body[i] = self._root.Two(io, self, self._root)
            else:
                self.body[i] = self._io.read_bytes(self.size)


    class One(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.first = self._io.read_bytes_full()


    class Two(KaitaiStruct):
        def __init__(self, _io, _parent=None, _root=None):
            self._io = _io
            self._parent = _parent
            self._root = _root if _root else self
            self._read()

        def _read(self):
            self.second = self._io.read_bytes_full()


