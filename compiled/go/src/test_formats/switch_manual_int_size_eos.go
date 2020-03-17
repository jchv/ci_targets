// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import (
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"bytes"
)

type SwitchManualIntSizeEos struct {
	Chunks []*SwitchManualIntSizeEos_Chunk
	_io *kaitai.Stream
	_root *SwitchManualIntSizeEos
	_parent interface{}
}

func (this *SwitchManualIntSizeEos) Read(io *kaitai.Stream, parent interface{}, root *SwitchManualIntSizeEos) (err error) {
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
		tmp2 := new(SwitchManualIntSizeEos_Chunk)
		err = tmp2.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Chunks = append(this.Chunks, tmp2)
	}
	return err
}
type SwitchManualIntSizeEos_Chunk struct {
	Code uint8
	Size uint32
	Body *SwitchManualIntSizeEos_ChunkBody
	_io *kaitai.Stream
	_root *SwitchManualIntSizeEos
	_parent *SwitchManualIntSizeEos
	_raw_Body []byte
}

func (this *SwitchManualIntSizeEos_Chunk) Read(io *kaitai.Stream, parent *SwitchManualIntSizeEos, root *SwitchManualIntSizeEos) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp3, err := this._io.ReadU1()
	if err != nil {
		return err
	}
	this.Code = tmp3
	tmp4, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Size = uint32(tmp4)
	tmp5, err := this._io.ReadBytes(int(this.Size))
	if err != nil {
		return err
	}
	tmp5 = tmp5
	this._raw_Body = tmp5
	_io__raw_Body := kaitai.NewStream(bytes.NewReader(this._raw_Body))
	tmp6 := new(SwitchManualIntSizeEos_ChunkBody)
	err = tmp6.Read(_io__raw_Body, this, this._root)
	if err != nil {
		return err
	}
	this.Body = tmp6
	return err
}
type SwitchManualIntSizeEos_ChunkBody struct {
	Body interface{}
	_io *kaitai.Stream
	_root *SwitchManualIntSizeEos
	_parent *SwitchManualIntSizeEos_Chunk
	_raw_Body []byte
}

func (this *SwitchManualIntSizeEos_ChunkBody) Read(io *kaitai.Stream, parent *SwitchManualIntSizeEos_Chunk, root *SwitchManualIntSizeEos) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	switch (this._parent.Code) {
	case 17:
		tmp7, err := this._io.ReadBytesFull()
		if err != nil {
			return err
		}
		tmp7 = tmp7
		this._raw_Body = tmp7
		_io__raw_Body := kaitai.NewStream(bytes.NewReader(this._raw_Body))
		tmp8 := new(SwitchManualIntSizeEos_ChunkBody_ChunkMeta)
		err = tmp8.Read(_io__raw_Body, this, this._root)
		if err != nil {
			return err
		}
		this.Body = tmp8
	case 34:
		tmp9, err := this._io.ReadBytesFull()
		if err != nil {
			return err
		}
		tmp9 = tmp9
		this._raw_Body = tmp9
		_io__raw_Body := kaitai.NewStream(bytes.NewReader(this._raw_Body))
		tmp10 := new(SwitchManualIntSizeEos_ChunkBody_ChunkDir)
		err = tmp10.Read(_io__raw_Body, this, this._root)
		if err != nil {
			return err
		}
		this.Body = tmp10
	default:
		tmp11, err := this._io.ReadBytesFull()
		if err != nil {
			return err
		}
		tmp11 = tmp11
		this._raw_Body = tmp11
	}
	return err
}
type SwitchManualIntSizeEos_ChunkBody_ChunkMeta struct {
	Title string
	Author string
	_io *kaitai.Stream
	_root *SwitchManualIntSizeEos
	_parent *SwitchManualIntSizeEos_ChunkBody
}

func (this *SwitchManualIntSizeEos_ChunkBody_ChunkMeta) Read(io *kaitai.Stream, parent *SwitchManualIntSizeEos_ChunkBody, root *SwitchManualIntSizeEos) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp12, err := this._io.ReadBytesTerm(0, false, true, true)
	if err != nil {
		return err
	}
	this.Title = string(tmp12)
	tmp13, err := this._io.ReadBytesTerm(0, false, true, true)
	if err != nil {
		return err
	}
	this.Author = string(tmp13)
	return err
}
type SwitchManualIntSizeEos_ChunkBody_ChunkDir struct {
	Entries []string
	_io *kaitai.Stream
	_root *SwitchManualIntSizeEos
	_parent *SwitchManualIntSizeEos_ChunkBody
}

func (this *SwitchManualIntSizeEos_ChunkBody_ChunkDir) Read(io *kaitai.Stream, parent *SwitchManualIntSizeEos_ChunkBody, root *SwitchManualIntSizeEos) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	for {
		tmp14, err := this._io.EOF()
		if err != nil {
			return err
		}
		if tmp14 {
			break
		}
		tmp15, err := this._io.ReadBytes(int(4))
		if err != nil {
			return err
		}
		tmp15 = tmp15
		this.Entries = append(this.Entries, string(tmp15))
	}
	return err
}