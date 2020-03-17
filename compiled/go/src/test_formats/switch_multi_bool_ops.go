// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type SwitchMultiBoolOps struct {
	Opcodes []*SwitchMultiBoolOps_Opcode
	_io *kaitai.Stream
	_root *SwitchMultiBoolOps
	_parent interface{}
}

func (this *SwitchMultiBoolOps) Read(io *kaitai.Stream, parent interface{}, root *SwitchMultiBoolOps) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	for {
		tmp1, err := this._io.EOF()
		if err != nil {
			return err
		}
		if tmp1 {
			break
		}
		tmp2 := new(SwitchMultiBoolOps_Opcode)
		err = tmp2.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Opcodes = append(this.Opcodes, tmp2)
	}
	return err
}
type SwitchMultiBoolOps_Opcode struct {
	Code uint8
	Body uint64
	_io *kaitai.Stream
	_root *SwitchMultiBoolOps
	_parent *SwitchMultiBoolOps
}

func (this *SwitchMultiBoolOps_Opcode) Read(io *kaitai.Stream, parent *SwitchMultiBoolOps, root *SwitchMultiBoolOps) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp3, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Code = tmp3
	var tmp4 uint8;
	var tmp5 bool;
	if (this.Code != 10) {
		tmp5 = true
	} else {
		tmp5 = false
	}
	if ( ((this.Code > 0) && (this.Code <= 8) && (tmp5)) ) {
		tmp4 = this.Code
	} else {
		tmp4 = 0
	}
	switch (tmp4) {
	case 1:
		tmp6, err := this._io.ReadU1()
		if err != nil {
			return err
		}
		this.Body = uint64(tmp6)
	case 2:
		tmp7, err := this._io.ReadU2le()
		if err != nil {
			return err
		}
		this.Body = uint64(tmp7)
	case 4:
		tmp8, err := this._io.ReadU4le()
		if err != nil {
			return err
		}
		this.Body = uint64(tmp8)
	case 8:
		tmp9, err := this._io.ReadU8le()
		if err != nil {
			return err
		}
		this.Body = uint64(tmp9)
	}
	return err
}