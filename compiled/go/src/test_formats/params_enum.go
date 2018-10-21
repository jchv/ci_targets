// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"


type ParamsEnum_Animal int
const (
	ParamsEnum_Animal__Dog ParamsEnum_Animal = 4
	ParamsEnum_Animal__Cat ParamsEnum_Animal = 7
	ParamsEnum_Animal__Chicken ParamsEnum_Animal = 12
)
type ParamsEnum struct {
	One ParamsEnum_Animal
	InvokeWithParam *ParamsEnum_WithParam
	_io *kaitai.Stream
	_root *ParamsEnum
	_parent interface{}
}

func (this *ParamsEnum) Read(io *kaitai.Stream, parent interface{}, root *ParamsEnum) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.One = ParamsEnum_Animal(tmp1)
	tmp2 := new(ParamsEnum_WithParam)
	err = tmp2.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.InvokeWithParam = tmp2
	return err
}
type ParamsEnum_WithParam struct {
	_io *kaitai.Stream
	_root *ParamsEnum
	_parent *ParamsEnum
	_f_isCat bool
	isCat bool
}

func (this *ParamsEnum_WithParam) Read(io *kaitai.Stream, parent *ParamsEnum, root *ParamsEnum) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	return err
}
func (this *ParamsEnum_WithParam) IsCat() (v bool, err error) {
	if (this._f_isCat) {
		return this.isCat, nil
	}
	this.isCat = bool(this.EnumeratedOne == ParamsEnum_Animal__Cat)
	this._f_isCat = true
	return this.isCat, nil
}
