// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package io.kaitai.struct.testformats;

import io.kaitai.struct.KaitaiStruct;
import io.kaitai.struct.KaitaiStream;
import java.io.IOException;
import java.util.ArrayList;
import java.nio.charset.Charset;

public class RepeatNStrz extends KaitaiStruct {
    public static RepeatNStrz fromFile(String fileName) throws IOException {
        return new RepeatNStrz(new KaitaiStream(fileName));
    }

    public RepeatNStrz(KaitaiStream _io) {
        this(_io, null, null);
    }

    public RepeatNStrz(KaitaiStream _io, KaitaiStruct _parent) {
        this(_io, _parent, null);
    }

    public RepeatNStrz(KaitaiStream _io, KaitaiStruct _parent, RepeatNStrz _root) {
        super(_io);
        this._parent = _parent;
        this._root = _root == null ? this : _root;
        _read();
    }
    private void _read() {
        this.qty = this._io.readU4le();
        lines = new ArrayList<String>((int) (qty()));
        for (int i = 0; i < qty(); i++) {
            this.lines.add(new String(this._io.readBytesTerm(0, false, true, true), Charset.forName("UTF-8")));
        }
    }
    private long qty;
    private ArrayList<String> lines;
    private RepeatNStrz _root;
    private KaitaiStruct _parent;
    public long qty() { return qty; }
    public ArrayList<String> lines() { return lines; }
    public RepeatNStrz _root() { return _root; }
    public KaitaiStruct _parent() { return _parent; }
}
