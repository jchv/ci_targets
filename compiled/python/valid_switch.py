# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

from pkg_resources import parse_version
import kaitaistruct
from kaitaistruct import KaitaiStruct, KaitaiStream, BytesIO


if parse_version(kaitaistruct.__version__) < parse_version('0.9'):
    raise Exception("Incompatible Kaitai Struct Python API: 0.9 or later is required, but you have %s" % (kaitaistruct.__version__))

class ValidSwitch(KaitaiStruct):
    def __init__(self, _io, _parent=None, _root=None):
        self._io = _io
        self._parent = _parent
        self._root = _root if _root else self
        self._read()

    def _read(self):
        self.a = self._io.read_u1()
        if not self.a == 80:
            raise kaitaistruct.ValidationNotEqualError(80, self.a, self._io, u"/seq/0")
        _on = self.a
        if _on == 80:
            self.b = self._io.read_u2le()
        else:
            self.b = self._io.read_u2be()
        if not self.b == 17217:
            raise kaitaistruct.ValidationNotEqualError(17217, self.b, self._io, u"/seq/1")


