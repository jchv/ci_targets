// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"


type EnumIntRangeU_Constants int
const (
	EnumIntRangeU_Constants__Zero EnumIntRangeU_Constants = 0
	EnumIntRangeU_Constants__IntMax EnumIntRangeU_Constants = 4294967295
)
type EnumIntRangeU struct {
	F1 EnumIntRangeU_Constants
	F2 EnumIntRangeU_Constants
	_io *kaitai.Stream
	_root *EnumIntRangeU
	_parent interface{}
}

func (this *EnumIntRangeU) Read(io *kaitai.Stream, parent interface{}, root *EnumIntRangeU) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU4be()
	if err != nil {
		return err
	}
	this.F1 = EnumIntRangeU_Constants(tmp1)
	tmp2, err := this._io.ReadU4be()
	if err != nil {
		return err
	}
	this.F2 = EnumIntRangeU_Constants(tmp2)
	return err
}
