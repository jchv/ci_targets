from construct import *
from construct.lib import *

imported_and_rel = Struct(
	'one' / Int8ub,
	'two' / LazyBound(lambda: imported_root),
)

_schema = imported_and_rel
