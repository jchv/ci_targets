import kaitai_struct_nim_runtime
import options
import encodings

template defineEnum(typ) =
  type typ* = distinct int64
  proc `==`*(x, y: typ): bool {.borrow.}

type
  ExprArray* = ref object of KaitaiStruct
    aint*: seq[uint32]
    afloat*: seq[float64]
    astr*: seq[string]
    parent*: KaitaiStruct
    aintFirstInst*: Option[uint32]
    afloatSizeInst*: Option[int]
    astrSizeInst*: Option[int]
    aintMinInst*: Option[uint32]
    afloatMinInst*: Option[float64]
    aintSizeInst*: Option[int]
    aintLastInst*: Option[uint32]
    afloatLastInst*: Option[float64]
    astrFirstInst*: Option[string]
    astrLastInst*: Option[string]
    aintMaxInst*: Option[uint32]
    afloatFirstInst*: Option[float64]
    astrMinInst*: Option[string]
    astrMaxInst*: Option[string]
    afloatMaxInst*: Option[float64]

proc read*(_: typedesc[ExprArray], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): ExprArray

proc aintFirst*(this: ExprArray): uint32
proc afloatSize*(this: ExprArray): int
proc astrSize*(this: ExprArray): int
proc aintMin*(this: ExprArray): uint32
proc afloatMin*(this: ExprArray): float64
proc aintSize*(this: ExprArray): int
proc aintLast*(this: ExprArray): uint32
proc afloatLast*(this: ExprArray): float64
proc astrFirst*(this: ExprArray): string
proc astrLast*(this: ExprArray): string
proc aintMax*(this: ExprArray): uint32
proc afloatFirst*(this: ExprArray): float64
proc astrMin*(this: ExprArray): string
proc astrMax*(this: ExprArray): string
proc afloatMax*(this: ExprArray): float64

proc read*(_: typedesc[ExprArray], io: KaitaiStream, root: KaitaiStruct, parent: KaitaiStruct): ExprArray =
  template this: untyped = result
  this = new(ExprArray)
  let root = if root == nil: cast[KaitaiStruct](this) else: root
  this.io = io
  this.root = root
  this.parent = parent

  for i in 0 ..< 4:
    this.aint.add(this.io.readU4le())
  for i in 0 ..< 3:
    this.afloat.add(this.io.readF8le())
  for i in 0 ..< 3:
    this.astr.add(convert(this.io.readBytesTerm(0, false, true, true), srcEncoding = "UTF-8"))

proc aintFirst(this: ExprArray): uint32 = 
  if isSome(this.aintFirstInst):
    return get(this.aintFirstInst)
  this.aintFirstInst = some(this.aint[0])
  if isSome(this.aintFirstInst):
    return get(this.aintFirstInst)

proc afloatSize(this: ExprArray): int = 
  if isSome(this.afloatSizeInst):
    return get(this.afloatSizeInst)
  this.afloatSizeInst = some(len(this.afloat))
  if isSome(this.afloatSizeInst):
    return get(this.afloatSizeInst)

proc astrSize(this: ExprArray): int = 
  if isSome(this.astrSizeInst):
    return get(this.astrSizeInst)
  this.astrSizeInst = some(len(this.astr))
  if isSome(this.astrSizeInst):
    return get(this.astrSizeInst)

proc aintMin(this: ExprArray): uint32 = 
  if isSome(this.aintMinInst):
    return get(this.aintMinInst)
  this.aintMinInst = some(min(this.aint))
  if isSome(this.aintMinInst):
    return get(this.aintMinInst)

proc afloatMin(this: ExprArray): float64 = 
  if isSome(this.afloatMinInst):
    return get(this.afloatMinInst)
  this.afloatMinInst = some(min(this.afloat))
  if isSome(this.afloatMinInst):
    return get(this.afloatMinInst)

proc aintSize(this: ExprArray): int = 
  if isSome(this.aintSizeInst):
    return get(this.aintSizeInst)
  this.aintSizeInst = some(len(this.aint))
  if isSome(this.aintSizeInst):
    return get(this.aintSizeInst)

proc aintLast(this: ExprArray): uint32 = 
  if isSome(this.aintLastInst):
    return get(this.aintLastInst)
  this.aintLastInst = some(this.aint[^1])
  if isSome(this.aintLastInst):
    return get(this.aintLastInst)

proc afloatLast(this: ExprArray): float64 = 
  if isSome(this.afloatLastInst):
    return get(this.afloatLastInst)
  this.afloatLastInst = some(this.afloat[^1])
  if isSome(this.afloatLastInst):
    return get(this.afloatLastInst)

proc astrFirst(this: ExprArray): string = 
  if isSome(this.astrFirstInst):
    return get(this.astrFirstInst)
  this.astrFirstInst = some(this.astr[0])
  if isSome(this.astrFirstInst):
    return get(this.astrFirstInst)

proc astrLast(this: ExprArray): string = 
  if isSome(this.astrLastInst):
    return get(this.astrLastInst)
  this.astrLastInst = some(this.astr[^1])
  if isSome(this.astrLastInst):
    return get(this.astrLastInst)

proc aintMax(this: ExprArray): uint32 = 
  if isSome(this.aintMaxInst):
    return get(this.aintMaxInst)
  this.aintMaxInst = some(max(this.aint))
  if isSome(this.aintMaxInst):
    return get(this.aintMaxInst)

proc afloatFirst(this: ExprArray): float64 = 
  if isSome(this.afloatFirstInst):
    return get(this.afloatFirstInst)
  this.afloatFirstInst = some(this.afloat[0])
  if isSome(this.afloatFirstInst):
    return get(this.afloatFirstInst)

proc astrMin(this: ExprArray): string = 
  if isSome(this.astrMinInst):
    return get(this.astrMinInst)
  this.astrMinInst = some(min(this.astr))
  if isSome(this.astrMinInst):
    return get(this.astrMinInst)

proc astrMax(this: ExprArray): string = 
  if isSome(this.astrMaxInst):
    return get(this.astrMaxInst)
  this.astrMaxInst = some(max(this.astr))
  if isSome(this.astrMaxInst):
    return get(this.astrMaxInst)

proc afloatMax(this: ExprArray): float64 = 
  if isSome(this.afloatMaxInst):
    return get(this.afloatMaxInst)
  this.afloatMaxInst = some(max(this.afloat))
  if isSome(this.afloatMaxInst):
    return get(this.afloatMaxInst)

proc fromFile*(_: typedesc[ExprArray], filename: string): ExprArray =
  ExprArray.read(newKaitaiFileStream(filename), nil, nil)

