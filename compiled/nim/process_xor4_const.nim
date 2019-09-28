# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

import ../../../runtime/nim/kaitai

type
  ProcessXor4Const* = ref object
    key*: seq[byte]
    buf*: seq[byte]
    root*: ProcessXor4Const
    parent*: ref RootObj
    raw_buf*: seq[byte]

proc read*(_: typedesc[ProcessXor4Const], stream: KaitaiStream, root: ProcessXor4Const, parent: ref RootObj): owned ProcessXor4Const =
  result = new(ProcessXor4Const)
  let root = if root == nil: cast[ProcessXor4Const](result) else: root
  result.key = readBytes(stream, int(4))
  result.buf = readBytesFull(stream)
  result.root = root
  result.parent = parent

proc fromFile*(_: typedesc[ProcessXor4Const], filename: string): owned ProcessXor4Const =
  var stream = newKaitaiStream(filename)
  ProcessXor4Const.read(stream, nil, nil)