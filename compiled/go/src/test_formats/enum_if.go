// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"


type EnumIf_Opcodes int
const (
	EnumIf_Opcodes__AString EnumIf_Opcodes = 83
	EnumIf_Opcodes__ATuple EnumIf_Opcodes = 84
)
type EnumIf struct {
	Op1 *EnumIf_Operation
	Op2 *EnumIf_Operation
	Op3 *EnumIf_Operation
	_io *kaitai.Stream
	_root *EnumIf
	_parent interface{}
}

func (this *EnumIf) Read(io *kaitai.Stream, parent interface{}, root *EnumIf) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1 := new(EnumIf_Operation)
	err = tmp1.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Op1 = tmp1
	tmp2 := new(EnumIf_Operation)
	err = tmp2.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Op2 = tmp2
	tmp3 := new(EnumIf_Operation)
	err = tmp3.Read(this._io, this, this._root)
	if err != nil {
		return err
	}
	this.Op3 = tmp3
	return err
}
type EnumIf_Operation struct {
	Opcode EnumIf_Opcodes
	ArgTuple *EnumIf_ArgTuple
	ArgStr *EnumIf_ArgStr
	_io *kaitai.Stream
	_root *EnumIf
	_parent *EnumIf
}

func (this *EnumIf_Operation) Read(io *kaitai.Stream, parent *EnumIf, root *EnumIf) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp4, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Opcode = EnumIf_Opcodes(tmp4)
	if (this.Opcode == EnumIf_Opcodes__ATuple) {
		tmp5 := new(EnumIf_ArgTuple)
		err = tmp5.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.ArgTuple = tmp5
	}
	if (this.Opcode == EnumIf_Opcodes__AString) {
		tmp6 := new(EnumIf_ArgStr)
		err = tmp6.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.ArgStr = tmp6
	}
	return err
}
type EnumIf_ArgTuple struct {
	Num1 uint8
	Num2 uint8
	_io *kaitai.Stream
	_root *EnumIf
	_parent *EnumIf_Operation
}

func (this *EnumIf_ArgTuple) Read(io *kaitai.Stream, parent *EnumIf_Operation, root *EnumIf) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp7, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Num1 = tmp7
	tmp8, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Num2 = tmp8
	return err
}
type EnumIf_ArgStr struct {
	Len uint8
	Str string
	_io *kaitai.Stream
	_root *EnumIf
	_parent *EnumIf_Operation
}

func (this *EnumIf_ArgStr) Read(io *kaitai.Stream, parent *EnumIf_Operation, root *EnumIf) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp9, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Len = tmp9
	tmp10, err := this._io.ReadBytes(int(this.Len))
	if err != nil {
		return err
	}
	tmp10 = tmp10
	this.Str = string(tmp10)
	return err
}
