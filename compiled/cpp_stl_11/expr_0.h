#pragma once

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "kaitai/kaitaistruct.h"
#include <stdint.h>
#include <memory>

#if KAITAI_STRUCT_VERSION < 9000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.9 or later is required"
#endif

class expr_0_t : public kaitai::kstruct {

public:

    expr_0_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent = nullptr, expr_0_t* p__root = nullptr);

private:
    void _read();

public:
    ~expr_0_t();

private:
    bool f_must_be_f7;
    int32_t m_must_be_f7;

public:
    int32_t must_be_f7();

private:
    bool f_must_be_abc123;
    std::string m_must_be_abc123;

public:
    std::string must_be_abc123();

private:
    uint16_t m_len_of_1;
    expr_0_t* m__root;
    kaitai::kstruct* m__parent;

public:
    uint16_t len_of_1() const { return m_len_of_1; }
    expr_0_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
};