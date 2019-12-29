#pragma once

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "kaitai/kaitaistruct.h"
#include <stdint.h>
#include <memory>

#if KAITAI_STRUCT_VERSION < 9000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.9 or later is required"
#endif
class vlq_base128_le_t;

class imports_abs_t : public kaitai::kstruct {

public:

    imports_abs_t(kaitai::kstream* p__io, std::unique_ptr<kaitai::kstruct> p__parent = nullptr, imports_abs_t* p__root = nullptr);

private:
    void _read();

public:
    ~imports_abs_t();

private:
    std::unique_ptr<vlq_base128_le_t> m_len;
    std::string m_body;
    imports_abs_t* m__root;
    std::unique_ptr<kaitai::kstruct> m__parent;

public:
    vlq_base128_le_t* len() const { return m_len.get(); }
    std::string body() const { return m_body; }
    imports_abs_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent.get(); }
};
