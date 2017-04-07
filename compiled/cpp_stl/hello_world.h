#ifndef HELLO_WORLD_H_
#define HELLO_WORLD_H_

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include <kaitai/kaitaistruct.h>
#include <kaitai/kaitaistream.h>

#include <stdint.h>
#include <vector>
#include <sstream>
#include <algorithm>

#if KAITAI_STRUCT_VERSION < 7000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.7 or later is required"
#endif

class hello_world_t : public kaitai::kstruct {

public:

    hello_world_t(kaitai::kstream* p_io, kaitai::kstruct* p_parent = 0, hello_world_t* p_root = 0);
    ~hello_world_t();

private:
    uint8_t m_one;
    hello_world_t* m__root;
    kaitai::kstruct* m__parent;

public:
    uint8_t one() const { return m_one; }
    hello_world_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
};

#endif  // HELLO_WORLD_H_
