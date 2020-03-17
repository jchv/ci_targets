#ifndef NESTED_SAME_NAME_H_
#define NESTED_SAME_NAME_H_

// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

#include "kaitai/kaitaistruct.h"
#include <stdint.h>

#if KAITAI_STRUCT_VERSION < 9000L
#error "Incompatible Kaitai Struct C++/STL API: version 0.9 or later is required"
#endif

class nested_same_name_t : public kaitai::kstruct {

public:
    class main_t;
    class dummy_obj_t;

    nested_same_name_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent = 0, nested_same_name_t* p__root = 0);

private:
    void _read();

public:
    ~nested_same_name_t();

    class main_t : public kaitai::kstruct {

    public:
        class foo_obj_t;

        main_t(kaitai::kstream* p__io, nested_same_name_t* p__parent = 0, nested_same_name_t* p__root = 0);

    private:
        void _read();

    public:
        ~main_t();

        class foo_obj_t : public kaitai::kstruct {

        public:

            foo_obj_t(kaitai::kstream* p__io, nested_same_name_t::main_t* p__parent = 0, nested_same_name_t* p__root = 0);

        private:
            void _read();

        public:
            ~foo_obj_t();

        private:
            std::string m_data;
            nested_same_name_t* m__root;
            nested_same_name_t::main_t* m__parent;

        public:
            std::string data() const { return m_data; }
            nested_same_name_t* _root() const { return m__root; }
            nested_same_name_t::main_t* _parent() const { return m__parent; }
        };

    private:
        int32_t m_main_size;
        foo_obj_t* m_foo;
        nested_same_name_t* m__root;
        nested_same_name_t* m__parent;

    public:
        int32_t main_size() const { return m_main_size; }
        foo_obj_t* foo() const { return m_foo; }
        nested_same_name_t* _root() const { return m__root; }
        nested_same_name_t* _parent() const { return m__parent; }
    };

    class dummy_obj_t : public kaitai::kstruct {

    public:
        class foo_t;

        dummy_obj_t(kaitai::kstream* p__io, nested_same_name_t* p__parent = 0, nested_same_name_t* p__root = 0);

    private:
        void _read();

    public:
        ~dummy_obj_t();

        class foo_t : public kaitai::kstruct {

        public:

            foo_t(kaitai::kstream* p__io, kaitai::kstruct* p__parent = 0, nested_same_name_t* p__root = 0);

        private:
            void _read();

        public:
            ~foo_t();

        private:
            nested_same_name_t* m__root;
            kaitai::kstruct* m__parent;

        public:
            nested_same_name_t* _root() const { return m__root; }
            kaitai::kstruct* _parent() const { return m__parent; }
        };

    private:
        nested_same_name_t* m__root;
        nested_same_name_t* m__parent;

    public:
        nested_same_name_t* _root() const { return m__root; }
        nested_same_name_t* _parent() const { return m__parent; }
    };

private:
    main_t* m_main_data;
    dummy_obj_t* m_dummy;
    nested_same_name_t* m__root;
    kaitai::kstruct* m__parent;

public:
    main_t* main_data() const { return m_main_data; }
    dummy_obj_t* dummy() const { return m_dummy; }
    nested_same_name_t* _root() const { return m__root; }
    kaitai::kstruct* _parent() const { return m__parent; }
};

#endif  // NESTED_SAME_NAME_H_