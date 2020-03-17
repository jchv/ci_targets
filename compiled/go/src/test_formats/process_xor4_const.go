// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type ProcessXor4Const struct {
	Key []byte
	Buf []byte
	_io *kaitai.Stream
	_root *ProcessXor4Const
	_parent interface{}
	_raw_Buf []byte
}

func (this *ProcessXor4Const) Read(io *kaitai.Stream, parent interface{}, root *ProcessXor4Const) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadBytes(int(4))
	if err != nil {
		return err
	}
	tmp1 = tmp1
	this.Key = tmp1
	tmp2, err := this._io.ReadBytesFull()
	if err != nil {
		return err
	}
	tmp2 = tmp2
	this._raw_Buf = tmp2
	this.Buf = kaitai.ProcessXOR(this._raw_Buf, []uint8{236, 187, 163, 20})
	return err
}