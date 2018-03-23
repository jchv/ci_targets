from construct import *
from construct.lib import *

floating_points = Struct(
	'single_value' / Float32l,
	'double_value' / Float64l,
	'single_value_be' / Float32b,
	'double_value_be' / Float64b,
	'approximate_value' / Float32l,
	'single_value_plus_int' / Computed((self.single_value + 1)),
	'single_value_plus_float' / Computed((self.single_value + 0.5)),
	'double_value_plus_float' / Computed((self.double_value + 0.05)),
)

_schema = floating_points
