// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

var CastToTop = (function() {
  function CastToTop(_io, _parent, _root) {
    this._io = _io;
    this._parent = _parent;
    this._root = _root || this;

    this._read();
  }
  CastToTop.prototype._read = function() {
    this.code = this._io.readU1();
  }
  Object.defineProperty(CastToTop.prototype, 'header', {
    get: function() {
      if (this._m_header !== undefined)
        return this._m_header;
      var _pos = this._io.pos;
      this._io.seek(1);
      this._m_header = new CastToTop(this._io, null, null);
      this._io.seek(_pos);
      return this._m_header;
    }
  });
  Object.defineProperty(CastToTop.prototype, 'headerCasted', {
    get: function() {
      if (this._m_headerCasted !== undefined)
        return this._m_headerCasted;
      this._m_headerCasted = this.header;
      return this._m_headerCasted;
    }
  });

  return CastToTop;
})();

// Export for amd environments
if (typeof define === 'function' && define.amd) {
  define('CastToTop', [], function() {
    return CastToTop;
  });
}

// Export for CommonJS
if (typeof module === 'object' && module && module.exports) {
  module.exports = CastToTop;
}
