// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

(function (root, factory) {
  if (typeof define === 'function' && define.amd) {
    define(['kaitai-struct/KaitaiStream'], factory);
  } else if (typeof module === 'object' && module.exports) {
    module.exports = factory(require('kaitai-struct/KaitaiStream'));
  } else {
    root.PositionToEnd = factory(root.KaitaiStream);
  }
}(this, function (KaitaiStream) {
var PositionToEnd = (function() {
  function PositionToEnd(_io, _parent, _root) {
    this._io = _io;
    this._parent = _parent;
    this._root = _root || this;

    this._read();
  }
  PositionToEnd.prototype._read = function() {
  }

  var IndexObj = PositionToEnd.IndexObj = (function() {
    function IndexObj(_io, _parent, _root) {
      this._io = _io;
      this._parent = _parent;
      this._root = _root || this;

      this._read();
    }
    IndexObj.prototype._read = function() {
      this.foo = this._io.readU4le();
      this.bar = this._io.readU4le();
    }

    return IndexObj;
  })();
  Object.defineProperty(PositionToEnd.prototype, 'index', {
    get: function() {
      if (this._m_index !== undefined)
        return this._m_index;
      var _pos = this._io.pos;
      this._io.seek((this._io.size - 8));
      this._m_index = new IndexObj(this._io, this, this._root);
      this._io.seek(_pos);
      return this._m_index;
    }
  });

  return PositionToEnd;
})();
return PositionToEnd;
}));
