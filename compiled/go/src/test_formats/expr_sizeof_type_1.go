// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type ExprSizeofType1 struct {
	_io *kaitai.Stream
	_root *ExprSizeofType1
	_parent interface{}
	_f_sizeofBlock bool
	sizeofBlock int
	_f_sizeofSubblock bool
	sizeofSubblock int
}

func (this *ExprSizeofType1) Read(io *kaitai.Stream, parent interface{}, root *ExprSizeofType1) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	return err
}
func (this *ExprSizeofType1) SizeofBlock() (v int, err error) {
	if (this._f_sizeofBlock) {
		return this.sizeofBlock, nil
	}
	this.sizeofBlock = int(11)
	this._f_sizeofBlock = true
	return this.sizeofBlock, nil
}
func (this *ExprSizeofType1) SizeofSubblock() (v int, err error) {
	if (this._f_sizeofSubblock) {
		return this.sizeofSubblock, nil
	}
	this.sizeofSubblock = int(4)
	this._f_sizeofSubblock = true
	return this.sizeofSubblock, nil
}
type ExprSizeofType1_Block struct {
	A uint8
	B uint32
	C []byte
	D *ExprSizeofType1_Block_Subblock
	_io *kaitai.Stream
	_root *ExprSizeofType1
	_parent interface{}
}

func (this *ExprSizeofType1_Block) Read(io *kaitai.Stream, parent interface{}, root *ExprSizeofType1) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.A = tmp1
	tmp2, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.B = uint32(tmp2)
	tmp3, err := this._io.ReadBytes(int(2))
	if err != nil {
		return err
	}
	tmp3 = tmp3
	this.C = tmp3
	tmp4 := new(ExprSizeofType1_Block_Subblock)
	err = tmp4.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.D = tmp4
	return err
}
type ExprSizeofType1_Block_Subblock struct {
	A []byte
	_io *kaitai.Stream
	_root *ExprSizeofType1
	_parent interface{}
}

func (this *ExprSizeofType1_Block_Subblock) Read(io *kaitai.Stream, parent interface{}, root *ExprSizeofType1) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp5, err := this._io.ReadBytes(int(4))
	if err != nil {
		return err
	}
	tmp5 = tmp5
	this.A = tmp5
	return err
}
