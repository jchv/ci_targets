# This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

require 'kaitai/struct/struct'

unless Gem::Version.new(Kaitai::Struct::VERSION) >= Gem::Version.new('0.9')
  raise "Incompatible Kaitai Struct Ruby API: 0.9 or later is required, but you have #{Kaitai::Struct::VERSION}"
end

class EnumForUnknownId < Kaitai::Struct::Struct

  ANIMAL = {
    4 => :animal_dog,
    7 => :animal_cat,
    12 => :animal_chicken,
  }
  I__ANIMAL = ANIMAL.invert
  def initialize(_io, _parent = nil, _root = self)
    super(_io, _parent, _root)
    _read
  end

  def _read
    @one = Kaitai::Struct::Stream::resolve_enum(ANIMAL, @_io.read_u1)
    self
  end
  attr_reader :one
end
