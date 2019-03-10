#ifndef ENUM_OF_VALUE_INST_H_
#define ENUM_OF_VALUE_INST_H_

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "kaitai/kaitaistruct.h"

#include <stdint.h>

#if KAITAI_STRUCT_VERSION < 7000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.7 or later is required"
#endif

class enum_of_value_inst_t : public kaitai::kstruct {

public:

    enum animal_t {
        ANIMAL_DOG = 4,
        ANIMAL_CAT = 7,
        ANIMAL_CHICKEN = 12
    };

    enum_of_value_inst_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent = 0, enum_of_value_inst_t* p__root = 0);

private:
    void _read();

public:
    ~enum_of_value_inst_t();

private:
    bool f_pet_3;
    animal_t m_pet_3;

public:
    animal_t pet_3();

private:
    bool f_pet_4;
    animal_t m_pet_4;

public:
    animal_t pet_4();

private:
    animal_t m_pet_1;
    animal_t m_pet_2;
    enum_of_value_inst_t* m__root;
    kaitai::kstruct* m__parent;

public:
    animal_t pet_1() const { return m_pet_1; }
    animal_t pet_2() const { return m_pet_2; }
    enum_of_value_inst_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
};

#endif  // ENUM_OF_VALUE_INST_H_