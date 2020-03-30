import kaitai_struct_nim_runtime
import options

type
  TypeTernary_Dummy* = ref TypeTernary_DummyObj
  TypeTernary_DummyObj* = object
    value*: uint8
    io*: KaitaiStream
    root*: TypeTernary
    parent*: TypeTernary
  TypeTernary* = ref TypeTernaryObj
  TypeTernaryObj* = object
    difWoHack*: TypeTernary_Dummy
    difWithHack*: TypeTernary_Dummy
    io*: KaitaiStream
    root*: TypeTernary
    parent*: ref RootObj
    rawDifWoHack*: string
    rawDifWithHack*: string
    rawRawDifWithHack*: string
    isHackInst*: Option[bool]
    difInst*: Option[TypeTernary_Dummy]
    difValueInst*: Option[uint8]

### TypeTernary_Dummy ###
proc read*(_: typedesc[TypeTernary_Dummy], io: KaitaiStream, root: TypeTernary, parent: TypeTernary): TypeTernary_Dummy =
  let this = new(TypeTernary_Dummy)
  let root = if root == nil: cast[TypeTernary](result) else: root
  this.io = io
  this.root = root
  this.parent = parent

  let value = this.io.readU1()
  this.value = value
  result = this

proc fromFile*(_: typedesc[TypeTernary_Dummy], filename: string): TypeTernary_Dummy =
  TypeTernary_Dummy.read(newKaitaiFileStream(filename), nil, nil)

proc `=destroy`(x: var TypeTernary_DummyObj) =
  close(x.io)

### TypeTernary ###
proc isHack*(this: TypeTernary): bool
proc dif*(this: TypeTernary): TypeTernary_Dummy
proc difValue*(this: TypeTernary): uint8
proc isHack(this: TypeTernary): bool = 
  if isSome(this.isHackInst):
    return get(this.isHackInst)
  let isHackInst = true
  this.isHackInst = some(isHackInst)
  return get(this.isHackInst)

proc dif(this: TypeTernary): TypeTernary_Dummy = 
  if isSome(this.difInst):
    return get(this.difInst)
  let difInst = (if not(this.isHack): this.difWoHack else: this.difWithHack)
  this.difInst = some(difInst)
  return get(this.difInst)

proc difValue(this: TypeTernary): uint8 = 
  if isSome(this.difValueInst):
    return get(this.difValueInst)
  let difValueInst = this.dif.this.value
  this.difValueInst = some(difValueInst)
  return get(this.difValueInst)

proc read*(_: typedesc[TypeTernary], io: KaitaiStream, root: TypeTernary, parent: ref RootObj): TypeTernary =
  let this = new(TypeTernary)
  let root = if root == nil: cast[TypeTernary](result) else: root
  this.io = io
  this.root = root
  this.parent = parent

  if not(this.isHack):
    let rawDifWoHack = this.io.readBytes(int(1))
    this.rawDifWoHack = rawDifWoHack
    let rawDifWoHackIo = newKaitaiStringStream(rawDifWoHack)
    let difWoHack = TypeTernary_Dummy.read(rawDifWoHackIo, this.root, this)
    this.difWoHack = difWoHack
  let rawRawDifWithHack = this.io.readBytes(int(1))
  this.rawRawDifWithHack = rawRawDifWithHack
  let rawDifWithHack = rawRawDifWithHack.processXor(3)
  this.rawDifWithHack = rawDifWithHack
  let rawDifWithHackIo = newKaitaiStringStream(rawDifWithHack)
  let difWithHack = TypeTernary_Dummy.read(rawDifWithHackIo, this.root, this)
  this.difWithHack = difWithHack
  result = this

proc fromFile*(_: typedesc[TypeTernary], filename: string): TypeTernary =
  TypeTernary.read(newKaitaiFileStream(filename), nil, nil)

proc `=destroy`(x: var TypeTernaryObj) =
  close(x.io)

