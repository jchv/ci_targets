import kaitai_struct_nim_runtime
import options

type
  IfInstances* = ref object of KaitaiStruct
    parent*: KaitaiStruct
    neverHappensInst*: Option[uint8]

proc neverHappens*(this: IfInstances): uint8
proc read*(_: typedesc[IfInstances], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): IfInstances =
  template this: untyped = result
  this = new(IfInstances)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent


proc neverHappens(this: IfInstances): uint8 = 
  if isSome(this.neverHappensInst):
    return get(this.neverHappensInst)
  if false:
    let pos = this.io.pos()
    this.io.seek(100500)
    this.neverHappensInst = some(this.io.readU1())
    this.io.seek(pos)
  return get(this.neverHappensInst)

proc fromFile*(_: typedesc[IfInstances], filename: string): IfInstances =
  IfInstances.read(newKaitaiFileStream(filename), nil, nil)

