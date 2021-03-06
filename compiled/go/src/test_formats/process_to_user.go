// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import (
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"bytes"
)

type ProcessToUser struct {
	Buf1 *ProcessToUser_JustStr
	_io *kaitai.Stream
	_root *ProcessToUser
	_parent interface{}
	_raw_Buf1 []byte
	_raw__raw_Buf1 []byte
}

func (this *ProcessToUser) Read(io *kaitai.Stream, parent interface{}, root *ProcessToUser) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadBytes(int(5))
	if err != nil {
		return err
	}
	tmp1 = tmp1
	this._raw__raw_Buf1 = tmp1
	this._raw_Buf1 = kaitai.ProcessRotateLeft(this._raw__raw_Buf1, int(3))
	_io__raw_Buf1 := kaitai.NewStream(bytes.NewReader(this._raw_Buf1))
	tmp2 := new(ProcessToUser_JustStr)
	err = tmp2.Read(_io__raw_Buf1, this, this._root)
	if err != nil {
		return err
	}
	this.Buf1 = tmp2
	return err
}
type ProcessToUser_JustStr struct {
	Str string
	_io *kaitai.Stream
	_root *ProcessToUser
	_parent *ProcessToUser
}

func (this *ProcessToUser_JustStr) Read(io *kaitai.Stream, parent *ProcessToUser, root *ProcessToUser) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp3, err := this._io.ReadBytesFull()
	if err != nil {
		return err
	}
	tmp3 = tmp3
	this.Str = string(tmp3)
	return err
}
