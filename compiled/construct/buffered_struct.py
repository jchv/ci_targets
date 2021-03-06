from construct import *
from construct.lib import *

buffered_struct__block = Struct(
	'number1' / Int32ul,
	'number2' / Int32ul,
)

buffered_struct = Struct(
	'len1' / Int32ul,
	'block1' / FixedSized(this.len1, LazyBound(lambda: buffered_struct__block)),
	'len2' / Int32ul,
	'block2' / FixedSized(this.len2, LazyBound(lambda: buffered_struct__block)),
	'finisher' / Int32ul,
)

_schema = buffered_struct
