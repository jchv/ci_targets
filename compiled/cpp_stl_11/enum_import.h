#pragma once

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "kaitai/kaitaistruct.h"
#include <stdint.h>
#include <memory>

#if KAITAI_STRUCT_VERSION < 9000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.9 or later is required"
#endif

class enum_import_t : public kaitai::kstruct {

public:

    enum_import_t(kaitai::kstream* p__io, std::unique_ptr<kaitai::kstruct> p__parent = nullptr, enum_import_t* p__root = nullptr);

private:
    void _read();

public:
    ~enum_import_t();

private:
    enum_0_t::animal_t m_pet_1;
    enum_deep_t::container1_t::container2_t::animal_t m_pet_2;
    enum_import_t* m__root;
    std::unique_ptr<kaitai::kstruct> m__parent;

public:
    enum_0_t::animal_t pet_1() const { return m_pet_1; }
    enum_deep_t::container1_t::container2_t::animal_t pet_2() const { return m_pet_2; }
    enum_import_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent.get(); }
};
