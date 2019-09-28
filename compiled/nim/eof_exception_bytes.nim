# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

import ../../../runtime/nim/kaitai

type
  EofExceptionBytes* = ref object
    buf*: seq[byte]
    root*: EofExceptionBytes
    parent*: ref RootObj

proc read*(_: typedesc[EofExceptionBytes], stream: KaitaiStream, root: EofExceptionBytes, parent: ref RootObj): owned EofExceptionBytes =
  result = new(EofExceptionBytes)
  let root = if root == nil: cast[EofExceptionBytes](result) else: root
  result.buf = readBytes(stream, int(13))
  result.root = root
  result.parent = parent

proc fromFile*(_: typedesc[EofExceptionBytes], filename: string): owned EofExceptionBytes =
  var stream = newKaitaiStream(filename)
  EofExceptionBytes.read(stream, nil, nil)