import kaitai_struct_nim_runtime
import options

type
  Expr0* = ref Expr0Obj
  Expr0Obj* = object
    lenOf1*: uint16
    io*: KaitaiStream
    root*: Expr0
    parent*: ref RootObj
    mustBeF7Inst*: Option[int]
    mustBeAbc123Inst*: Option[string]

### Expr0 ###
proc mustBeF7*(this: Expr0): int
proc mustBeAbc123*(this: Expr0): string
proc mustBeF7(this: Expr0): int = 
  if isSome(this.mustBeF7Inst):
    return get(this.mustBeF7Inst)
  let mustBeF7Inst = (7 + 240)
  this.mustBeF7Inst = some(mustBeF7Inst)
  return get(this.mustBeF7Inst)

proc mustBeAbc123(this: Expr0): string = 
  if isSome(this.mustBeAbc123Inst):
    return get(this.mustBeAbc123Inst)
  let mustBeAbc123Inst = "abc" & "123"
  this.mustBeAbc123Inst = some(mustBeAbc123Inst)
  return get(this.mustBeAbc123Inst)

proc read*(_: typedesc[Expr0], io: KaitaiStream, root: Expr0, parent: ref RootObj): Expr0 =
  let this = new(Expr0)
  let root = if root == nil: cast[Expr0](result) else: root
  this.io = io
  this.root = root
  this.parent = parent

  let lenOf1 = this.io.readU2le()
  this.lenOf1 = lenOf1
  result = this

proc fromFile*(_: typedesc[Expr0], filename: string): Expr0 =
  Expr0.read(newKaitaiFileStream(filename), nil, nil)

proc `=destroy`(x: var Expr0Obj) =
  close(x.io)

