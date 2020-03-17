// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include <memory>
#include "str_encodings.h"

str_encodings_t::str_encodings_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent, str_encodings_t* p__root) : kaitai::kstruct(p__io) {
    m__parent = p__parent;
    m__root = this;
    _read();
}

void str_encodings_t::_read() {
    m_len_of_1 = m__io->read_u2le();
    m_str1 = kaitai::kstream::bytes_to_str(m__io->read_bytes(len_of_1()), std::string("ASCII"));
    m_len_of_2 = m__io->read_u2le();
    m_str2 = kaitai::kstream::bytes_to_str(m__io->read_bytes(len_of_2()), std::string("UTF-8"));
    m_len_of_3 = m__io->read_u2le();
    m_str3 = kaitai::kstream::bytes_to_str(m__io->read_bytes(len_of_3()), std::string("SJIS"));
    m_len_of_4 = m__io->read_u2le();
    m_str4 = kaitai::kstream::bytes_to_str(m__io->read_bytes(len_of_4()), std::string("CP437"));
}

str_encodings_t::~str_encodings_t() {
}