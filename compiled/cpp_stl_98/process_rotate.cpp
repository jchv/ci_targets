// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include <memory>
#include "process_rotate.h"

process_rotate_t::process_rotate_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent, process_rotate_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = this;
    m__io_buf1 = 0;
    m__io_buf2 = 0;
    m__io_buf3 = 0;
    _read();
}

void process_rotate_t::_read() {
    m__raw_buf1 = m__io->read_bytes(5);
    m_buf1 = kaitai::kstream::process_rotate_left(m__raw_buf1, 3);
    m__raw_buf2 = m__io->read_bytes(5);
    m_buf2 = kaitai::kstream::process_rotate_left(m__raw_buf2, 8 - (3));
    m_key = m__io->read_u1();
    m__raw_buf3 = m__io->read_bytes(5);
    m_buf3 = kaitai::kstream::process_rotate_left(m__raw_buf3, key());
}

process_rotate_t::~process_rotate_t() {
}