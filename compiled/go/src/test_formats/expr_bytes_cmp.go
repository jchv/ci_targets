// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import (
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"bytes"
)

type ExprBytesCmp struct {
	One []byte
	Two []byte
	_io *kaitai.Stream
	_root *ExprBytesCmp
	_parent interface{}
	_f_isLe bool
	isLe bool
	_f_ack bool
	ack []byte
	_f_isGt2 bool
	isGt2 bool
	_f_isGt bool
	isGt bool
	_f_ack2 bool
	ack2 []byte
	_f_isEq bool
	isEq bool
	_f_isLt2 bool
	isLt2 bool
	_f_isGe bool
	isGe bool
	_f_hiVal bool
	hiVal []byte
	_f_isNe bool
	isNe bool
	_f_isLt bool
	isLt bool
}

func (this *ExprBytesCmp) Read(io *kaitai.Stream, parent interface{}, root *ExprBytesCmp) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadBytes(int(1))
	if err != nil {
		return err
	}
	tmp1 = tmp1
	this.One = tmp1
	tmp2, err := this._io.ReadBytes(int(3))
	if err != nil {
		return err
	}
	tmp2 = tmp2
	this.Two = tmp2
	return err
}
func (this *ExprBytesCmp) IsLe() (v bool, err error) {
	if (this._f_isLe) {
		return this.isLe, nil
	}
	tmp3, err := this.Ack2()
	if err != nil {
		return false, err
	}
	this.isLe = bool((bytes.Compare(this.Two, tmp3) <= 0))
	this._f_isLe = true
	return this.isLe, nil
}
func (this *ExprBytesCmp) Ack() (v []byte, err error) {
	if (this._f_ack) {
		return this.ack, nil
	}
	this.ack = []byte([]uint8{65, 67, 75})
	this._f_ack = true
	return this.ack, nil
}
func (this *ExprBytesCmp) IsGt2() (v bool, err error) {
	if (this._f_isGt2) {
		return this.isGt2, nil
	}
	tmp4, err := this.HiVal()
	if err != nil {
		return false, err
	}
	this.isGt2 = bool((bytes.Compare(tmp4, this.Two) > 0))
	this._f_isGt2 = true
	return this.isGt2, nil
}
func (this *ExprBytesCmp) IsGt() (v bool, err error) {
	if (this._f_isGt) {
		return this.isGt, nil
	}
	tmp5, err := this.Ack2()
	if err != nil {
		return false, err
	}
	this.isGt = bool((bytes.Compare(this.Two, tmp5) > 0))
	this._f_isGt = true
	return this.isGt, nil
}
func (this *ExprBytesCmp) Ack2() (v []byte, err error) {
	if (this._f_ack2) {
		return this.ack2, nil
	}
	this.ack2 = []byte([]uint8{65, 67, 75, 50})
	this._f_ack2 = true
	return this.ack2, nil
}
func (this *ExprBytesCmp) IsEq() (v bool, err error) {
	if (this._f_isEq) {
		return this.isEq, nil
	}
	tmp6, err := this.Ack()
	if err != nil {
		return false, err
	}
	this.isEq = bool(bytes.Equal(this.Two, tmp6))
	this._f_isEq = true
	return this.isEq, nil
}
func (this *ExprBytesCmp) IsLt2() (v bool, err error) {
	if (this._f_isLt2) {
		return this.isLt2, nil
	}
	this.isLt2 = bool((bytes.Compare(this.One, this.Two) < 0))
	this._f_isLt2 = true
	return this.isLt2, nil
}
func (this *ExprBytesCmp) IsGe() (v bool, err error) {
	if (this._f_isGe) {
		return this.isGe, nil
	}
	tmp7, err := this.Ack2()
	if err != nil {
		return false, err
	}
	this.isGe = bool((bytes.Compare(this.Two, tmp7) >= 0))
	this._f_isGe = true
	return this.isGe, nil
}
func (this *ExprBytesCmp) HiVal() (v []byte, err error) {
	if (this._f_hiVal) {
		return this.hiVal, nil
	}
	this.hiVal = []byte([]uint8{144, 67})
	this._f_hiVal = true
	return this.hiVal, nil
}
func (this *ExprBytesCmp) IsNe() (v bool, err error) {
	if (this._f_isNe) {
		return this.isNe, nil
	}
	tmp8, err := this.Ack()
	if err != nil {
		return false, err
	}
	this.isNe = bool((bytes.Compare(this.Two, tmp8) != 0))
	this._f_isNe = true
	return this.isNe, nil
}
func (this *ExprBytesCmp) IsLt() (v bool, err error) {
	if (this._f_isLt) {
		return this.isLt, nil
	}
	tmp9, err := this.Ack2()
	if err != nil {
		return false, err
	}
	this.isLt = bool((bytes.Compare(this.Two, tmp9) < 0))
	this._f_isLt = true
	return this.isLt, nil
}