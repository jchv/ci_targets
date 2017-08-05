#ifndef INDEX_SIZES_H_
#define INDEX_SIZES_H_

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include <kaitai/kaitaistruct.h>
#include <kaitai/kaitaistream.h>

#include <stdint.h>
#include <vector>

#if KAITAI_STRUCT_VERSION < 7000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.7 or later is required"
#endif

class index_sizes_t : public kaitai::kstruct {

public:

    index_sizes_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent = 0, index_sizes_t* p__root = 0);
    void _read();
    ~index_sizes_t();

private:
    uint32_t m_qty;
    std::vector<uint32_t>* m_sizes;
    std::vector<std::string>* m_bufs;
    index_sizes_t* m__root;
    kaitai::kstruct* m__parent;

public:
    uint32_t qty() const { return m_qty; }
    std::vector<uint32_t>* sizes() const { return m_sizes; }
    std::vector<std::string>* bufs() const { return m_bufs; }
    index_sizes_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
};

#endif  // INDEX_SIZES_H_
