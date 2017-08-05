#ifndef INSTANCE_STD_H_
#define INSTANCE_STD_H_

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include <kaitai/kaitaistruct.h>
#include <kaitai/kaitaistream.h>

#include <stdint.h>

#if KAITAI_STRUCT_VERSION < 7000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.7 or later is required"
#endif

class instance_std_t : public kaitai::kstruct {

public:

    instance_std_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent = 0, instance_std_t* p__root = 0);
    void _read();
    ~instance_std_t();

private:
    bool f_header;
    std::string m_header;

public:
    std::string header();

private:
    instance_std_t* m__root;
    kaitai::kstruct* m__parent;

public:
    instance_std_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
};

#endif  // INSTANCE_STD_H_
