// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import (
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"io"
)

type NonStandard struct {
	Foo uint8
	Bar uint32
	_io *kaitai.Stream
	_root *NonStandard
	_parent interface{}
	_f_vi bool
	vi uint8
	_f_pi bool
	pi uint8
}

func (this *NonStandard) Read(io *kaitai.Stream, parent interface{}, root *NonStandard) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Foo = tmp1
	switch (this.Foo) {
	case 42:
		tmp2, err := this._io.ReadU2le()
		if err != nil {
			return err
		}
		this.Bar = uint32(tmp2)
	case 43:
		tmp3, err := this._io.ReadU4le()
		if err != nil {
			return err
		}
		this.Bar = uint32(tmp3)
	}
	return err
}
func (this *NonStandard) Vi() (v uint8, err error) {
	if (this._f_vi) {
		return this.vi, nil
	}
	this.vi = uint8(this.Foo)
	this._f_vi = true
	return this.vi, nil
}
func (this *NonStandard) Pi() (v uint8, err error) {
	if (this._f_pi) {
		return this.pi, nil
	}
	_pos, err := this._io.Pos()
	if err != nil {
		return 0, err
	}
	_, err = this._io.Seek(int64(0), io.SeekStart)
	if err != nil {
		return 0, err
	}
	tmp4, err := this._io.ReadU1()
	if err != nil {
		return 0, err
	}
	this.pi = tmp4
	_, err = this._io.Seek(_pos, io.SeekStart)
	if err != nil {
		return 0, err
	}
	this._f_pi = true
	this._f_pi = true
	return this.pi, nil
}
