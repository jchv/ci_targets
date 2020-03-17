// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type EnumDeep struct {
	Pet1 EnumDeep_Container1_Animal
	Pet2 EnumDeep_Container1_Container2_Animal
	_io *kaitai.Stream
	_root *EnumDeep
	_parent interface{}
}

func (this *EnumDeep) Read(io *kaitai.Stream, parent interface{}, root *EnumDeep) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Pet1 = EnumDeep_Container1_Animal(tmp1)
	tmp2, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Pet2 = EnumDeep_Container1_Container2_Animal(tmp2)
	return err
}

type EnumDeep_Container1_Animal int
const (
	EnumDeep_Container1_Animal__Dog EnumDeep_Container1_Animal = 4
	EnumDeep_Container1_Animal__Cat EnumDeep_Container1_Animal = 7
	EnumDeep_Container1_Animal__Chicken EnumDeep_Container1_Animal = 12
)
type EnumDeep_Container1 struct {
	_io *kaitai.Stream
	_root *EnumDeep
	_parent interface{}
}

func (this *EnumDeep_Container1) Read(io *kaitai.Stream, parent interface{}, root *EnumDeep) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	return err
}

type EnumDeep_Container1_Container2_Animal int
const (
	EnumDeep_Container1_Container2_Animal__Canary EnumDeep_Container1_Container2_Animal = 4
	EnumDeep_Container1_Container2_Animal__Turtle EnumDeep_Container1_Container2_Animal = 7
	EnumDeep_Container1_Container2_Animal__Hare EnumDeep_Container1_Container2_Animal = 12
)
type EnumDeep_Container1_Container2 struct {
	_io *kaitai.Stream
	_root *EnumDeep
	_parent interface{}
}

func (this *EnumDeep_Container1_Container2) Read(io *kaitai.Stream, parent interface{}, root *EnumDeep) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	return err
}