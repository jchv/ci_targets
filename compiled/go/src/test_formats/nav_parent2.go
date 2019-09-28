// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import (
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"io"
)

type NavParent2 struct {
	OfsTags uint32
	NumTags uint32
	Tags []*NavParent2_Tag
	_io *kaitai.Stream
	_root *NavParent2
	_parent interface{}
}

func (this *NavParent2) Read(io *kaitai.Stream, parent interface{}, root *NavParent2) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp1, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.OfsTags = uint32(tmp1)
	tmp2, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.NumTags = uint32(tmp2)
	this.Tags = make([]*NavParent2_Tag, this.NumTags)
	for i := range this.Tags {
		tmp3 := new(NavParent2_Tag)
		err = tmp3.Read(this._io, this, this._root)
		if err != nil {
			return err
		}
		this.Tags[i] = tmp3
	}
	return err
}
type NavParent2_Tag struct {
	Name string
	Ofs uint32
	NumItems uint32
	_io *kaitai.Stream
	_root *NavParent2
	_parent *NavParent2
	_f_tagContent bool
	tagContent *NavParent2_Tag_TagChar
}

func (this *NavParent2_Tag) Read(io *kaitai.Stream, parent *NavParent2, root *NavParent2) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp4, err := this._io.ReadBytes(int(4))
	if err != nil {
		return err
	}
	tmp4 = tmp4
	this.Name = string(tmp4)
	tmp5, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.Ofs = uint32(tmp5)
	tmp6, err := this._io.ReadU4le()
	if err != nil {
		return err
	}
	this.NumItems = uint32(tmp6)
	return err
}
func (this *NavParent2_Tag) TagContent() (v *NavParent2_Tag_TagChar, err error) {
	if (this._f_tagContent) {
		return this.tagContent, nil
	}
	thisIo := this._root._io
	_pos, err := thisIo.Pos()
	if err != nil {
		return nil, err
	}
	_, err = thisIo.Seek(int64(this.Ofs), io.SeekStart)
	if err != nil {
		return nil, err
	}
	switch (this.Name) {
	case "RAHC":
		tmp7 := new(NavParent2_Tag_TagChar)
		err = tmp7.Read(thisIo, this, this._root)
		if err != nil {
			return nil, err
		}
		this.tagContent = tmp7
	}
	_, err = thisIo.Seek(_pos, io.SeekStart)
	if err != nil {
		return nil, err
	}
	this._f_tagContent = true
	return this.tagContent, nil
}
type NavParent2_Tag_TagChar struct {
	Content string
	_io *kaitai.Stream
	_root *NavParent2
	_parent *NavParent2_Tag
}

func (this *NavParent2_Tag_TagChar) Read(io *kaitai.Stream, parent *NavParent2_Tag, root *NavParent2) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	tmp8, err := this._io.ReadBytes(int(this._parent.NumItems))
	if err != nil {
		return err
	}
	tmp8 = tmp8
	this.Content = string(tmp8)
	return err
}