// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"


type EnumForUnknownId_Animal int
const (
	EnumForUnknownId_Animal__Dog EnumForUnknownId_Animal = 4
	EnumForUnknownId_Animal__Cat EnumForUnknownId_Animal = 7
	EnumForUnknownId_Animal__Chicken EnumForUnknownId_Animal = 12
)
type EnumForUnknownId struct {
	One EnumForUnknownId_Animal
	_io *kaitai.Stream
	_root *EnumForUnknownId
	_parent interface{}
}

func (this *EnumForUnknownId) Read(io *kaitai.Stream, parent interface{}, root *EnumForUnknownId) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.One = EnumForUnknownId_Animal(tmp1)
	return err
}
