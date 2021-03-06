// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

(function (root, factory) {
  if (typeof define === 'function' && define.amd) {
    define(['kaitai-struct/KaitaiStream'], factory);
  } else if (typeof module === 'object' && module.exports) {
    module.exports = factory(require('kaitai-struct/KaitaiStream'));
  } else {
    root.ProcessCoerceUsertype1 = factory(root.KaitaiStream);
  }
}(this, function (KaitaiStream) {
var ProcessCoerceUsertype1 = (function() {
  function ProcessCoerceUsertype1(_io, _parent, _root) {
    this._io = _io;
    this._parent = _parent;
    this._root = _root || this;

    this._read();
  }
  ProcessCoerceUsertype1.prototype._read = function() {
    this.records = new Array(2);
    for (var i = 0; i < 2; i++) {
      this.records[i] = new Record(this._io, this, this._root);
    }
  }

  var Record = ProcessCoerceUsertype1.Record = (function() {
    function Record(_io, _parent, _root) {
      this._io = _io;
      this._parent = _parent;
      this._root = _root || this;

      this._read();
    }
    Record.prototype._read = function() {
      this.flag = this._io.readU1();
      if (this.flag == 0) {
        this._raw_bufUnproc = this._io.readBytes(4);
        var _io__raw_bufUnproc = new KaitaiStream(this._raw_bufUnproc);
        this.bufUnproc = new Foo(_io__raw_bufUnproc, this, this._root);
      }
      if (this.flag != 0) {
        this._raw__raw_bufProc = this._io.readBytes(4);
        this._raw_bufProc = KaitaiStream.processXorOne(this._raw__raw_bufProc, 170);
        var _io__raw_bufProc = new KaitaiStream(this._raw_bufProc);
        this.bufProc = new Foo(_io__raw_bufProc, this, this._root);
      }
    }
    Object.defineProperty(Record.prototype, 'buf', {
      get: function() {
        if (this._m_buf !== undefined)
          return this._m_buf;
        this._m_buf = (this.flag == 0 ? this.bufUnproc : this.bufProc);
        return this._m_buf;
      }
    });

    return Record;
  })();

  var Foo = ProcessCoerceUsertype1.Foo = (function() {
    function Foo(_io, _parent, _root) {
      this._io = _io;
      this._parent = _parent;
      this._root = _root || this;

      this._read();
    }
    Foo.prototype._read = function() {
      this.value = this._io.readU4le();
    }

    return Foo;
  })();

  return ProcessCoerceUsertype1;
})();
return ProcessCoerceUsertype1;
}));
