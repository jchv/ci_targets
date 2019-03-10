// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include <memory>
#include "repeat_until_sized.h"

#include <memory>

repeat_until_sized_t::repeat_until_sized_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent, repeat_until_sized_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = this;
    m_records = nullptr;
    m__raw_records = nullptr;
    m__io__raw_records = nullptr;
    _read();
}

void repeat_until_sized_t::_read() {
    m__raw_records = new std::vector<std::string>();
    m__io__raw_records = new std::vector<kaitai::kstream*>();
    m_records = new std::vector<std::unique_ptr<record_t>>();
    {
        int i = 0;
        std::unique_ptr<record_t> _;
        do {
            std::string _buf = m__io->read_bytes(5);
            m__raw_records->push_back(std::move(_buf));
            kaitai::kstream* io__raw_records = new kaitai::kstream(_buf);
            m__io__raw_records->push_back(io__raw_records);
            _ = std::unique_ptr<record_t>(new record_t(io__raw_records, this, m__root));
            m_records->push_back(std::move(_));
            i++;
        } while (!(_->marker() == 170));
    }
}

repeat_until_sized_t::~repeat_until_sized_t() {
    delete m__raw_records;
    for (std::vector<kaitai::kstream*>::iterator it = m__io__raw_records->begin(); it != m__io__raw_records->end(); ++it) {
        delete *it;
    }
    delete m__io__raw_records;
}

repeat_until_sized_t::record_t::record_t(kaitai::kstream* p__io, repeat_until_sized_t* p__parent, repeat_until_sized_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = p__root;
    _read();
}

void repeat_until_sized_t::record_t::_read() {
    m_marker = m__io->read_u1();
    m_body = m__io->read_u4le();
}

repeat_until_sized_t::record_t::~record_t() {
}