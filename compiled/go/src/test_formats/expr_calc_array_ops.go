// This is a generated file! Please edit source .ksy file and use kaitai-struct-compiler to rebuild

package test_formats

import "github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"

type ExprCalcArrayOps struct {
	_io *kaitai.Stream
	_root *ExprCalcArrayOps
	_parent interface{}
	_f_doubleArray bool
	doubleArray []float64
	_f_intArraySize bool
	intArraySize int
	_f_intArrayMax bool
	intArrayMax int
	_f_doubleArrayMax bool
	doubleArrayMax float64
	_f_strArrayMax bool
	strArrayMax string
	_f_strArrayMin bool
	strArrayMin string
	_f_doubleArrayMid bool
	doubleArrayMid float64
	_f_strArray bool
	strArray []string
	_f_doubleArraySize bool
	doubleArraySize int
	_f_strArrayFirst bool
	strArrayFirst string
	_f_strArrayLast bool
	strArrayLast string
	_f_strArrayMid bool
	strArrayMid string
	_f_doubleArrayLast bool
	doubleArrayLast float64
	_f_intArrayMin bool
	intArrayMin int
	_f_strArraySize bool
	strArraySize int
	_f_intArrayFirst bool
	intArrayFirst int
	_f_doubleArrayFirst bool
	doubleArrayFirst float64
	_f_intArrayMid bool
	intArrayMid int
	_f_doubleArrayMin bool
	doubleArrayMin float64
	_f_intArray bool
	intArray []int
	_f_intArrayLast bool
	intArrayLast int
}

func (this *ExprCalcArrayOps) Read(io *kaitai.Stream, parent interface{}, root *ExprCalcArrayOps) (err error) {
	this._io = io
	this._parent = parent
	this._root = root

	return err
}
func (this *ExprCalcArrayOps) DoubleArray() (v []float64, err error) {
	if (this._f_doubleArray) {
		return this.doubleArray, nil
	}
	this.doubleArray = []float64([]float64{10.0, 25.0, 50.0, 100.0, 3.14159})
	this._f_doubleArray = true
	return this.doubleArray, nil
}
func (this *ExprCalcArrayOps) IntArraySize() (v int, err error) {
	if (this._f_intArraySize) {
		return this.intArraySize, nil
	}
	tmp1, err := this.IntArray()
	if err != nil {
		return 0, err
	}
	this.intArraySize = int(len(tmp1))
	this._f_intArraySize = true
	return this.intArraySize, nil
}
func (this *ExprCalcArrayOps) IntArrayMax() (v int, err error) {
	if (this._f_intArrayMax) {
		return this.intArrayMax, nil
	}
	tmp4, err := this.IntArray()
	if err != nil {
		return 0, err
	}
	tmp2 := tmp4[0]
	tmp5, err := this.IntArray()
	if err != nil {
		return 0, err
	}
	for _, tmp3 := range tmp5 {
		if tmp2 < tmp3 {
			tmp2 = tmp3
		}
	}
	this.intArrayMax = int(tmp2)
	this._f_intArrayMax = true
	return this.intArrayMax, nil
}
func (this *ExprCalcArrayOps) DoubleArrayMax() (v float64, err error) {
	if (this._f_doubleArrayMax) {
		return this.doubleArrayMax, nil
	}
	tmp8, err := this.DoubleArray()
	if err != nil {
		return 0, err
	}
	tmp6 := tmp8[0]
	tmp9, err := this.DoubleArray()
	if err != nil {
		return 0, err
	}
	for _, tmp7 := range tmp9 {
		if tmp6 < tmp7 {
			tmp6 = tmp7
		}
	}
	this.doubleArrayMax = float64(tmp6)
	this._f_doubleArrayMax = true
	return this.doubleArrayMax, nil
}
func (this *ExprCalcArrayOps) StrArrayMax() (v string, err error) {
	if (this._f_strArrayMax) {
		return this.strArrayMax, nil
	}
	tmp12, err := this.StrArray()
	if err != nil {
		return "", err
	}
	tmp10 := tmp12[0]
	tmp13, err := this.StrArray()
	if err != nil {
		return "", err
	}
	for _, tmp11 := range tmp13 {
		if tmp10 < tmp11 {
			tmp10 = tmp11
		}
	}
	this.strArrayMax = string(tmp10)
	this._f_strArrayMax = true
	return this.strArrayMax, nil
}
func (this *ExprCalcArrayOps) StrArrayMin() (v string, err error) {
	if (this._f_strArrayMin) {
		return this.strArrayMin, nil
	}
	tmp16, err := this.StrArray()
	if err != nil {
		return "", err
	}
	tmp14 := tmp16[0]
	tmp17, err := this.StrArray()
	if err != nil {
		return "", err
	}
	for _, tmp15 := range tmp17 {
		if tmp14 > tmp15 {
			tmp14 = tmp15
		}
	}
	this.strArrayMin = string(tmp14)
	this._f_strArrayMin = true
	return this.strArrayMin, nil
}
func (this *ExprCalcArrayOps) DoubleArrayMid() (v float64, err error) {
	if (this._f_doubleArrayMid) {
		return this.doubleArrayMid, nil
	}
	tmp18, err := this.DoubleArray()
	if err != nil {
		return 0, err
	}
	this.doubleArrayMid = float64(tmp18[1])
	this._f_doubleArrayMid = true
	return this.doubleArrayMid, nil
}
func (this *ExprCalcArrayOps) StrArray() (v []string, err error) {
	if (this._f_strArray) {
		return this.strArray, nil
	}
	this.strArray = []string([]string{"un", "deux", "trois", "quatre"})
	this._f_strArray = true
	return this.strArray, nil
}
func (this *ExprCalcArrayOps) DoubleArraySize() (v int, err error) {
	if (this._f_doubleArraySize) {
		return this.doubleArraySize, nil
	}
	tmp19, err := this.DoubleArray()
	if err != nil {
		return 0, err
	}
	this.doubleArraySize = int(len(tmp19))
	this._f_doubleArraySize = true
	return this.doubleArraySize, nil
}
func (this *ExprCalcArrayOps) StrArrayFirst() (v string, err error) {
	if (this._f_strArrayFirst) {
		return this.strArrayFirst, nil
	}
	tmp20, err := this.StrArray()
	if err != nil {
		return "", err
	}
	this.strArrayFirst = string(tmp20[0])
	this._f_strArrayFirst = true
	return this.strArrayFirst, nil
}
func (this *ExprCalcArrayOps) StrArrayLast() (v string, err error) {
	if (this._f_strArrayLast) {
		return this.strArrayLast, nil
	}
	tmp22, err := this.StrArray()
	if err != nil {
		return "", err
	}
	tmp21 := tmp22
	this.strArrayLast = string(tmp21[len(tmp21) - 1])
	this._f_strArrayLast = true
	return this.strArrayLast, nil
}
func (this *ExprCalcArrayOps) StrArrayMid() (v string, err error) {
	if (this._f_strArrayMid) {
		return this.strArrayMid, nil
	}
	tmp23, err := this.StrArray()
	if err != nil {
		return "", err
	}
	this.strArrayMid = string(tmp23[1])
	this._f_strArrayMid = true
	return this.strArrayMid, nil
}
func (this *ExprCalcArrayOps) DoubleArrayLast() (v float64, err error) {
	if (this._f_doubleArrayLast) {
		return this.doubleArrayLast, nil
	}
	tmp25, err := this.DoubleArray()
	if err != nil {
		return 0, err
	}
	tmp24 := tmp25
	this.doubleArrayLast = float64(tmp24[len(tmp24) - 1])
	this._f_doubleArrayLast = true
	return this.doubleArrayLast, nil
}
func (this *ExprCalcArrayOps) IntArrayMin() (v int, err error) {
	if (this._f_intArrayMin) {
		return this.intArrayMin, nil
	}
	tmp28, err := this.IntArray()
	if err != nil {
		return 0, err
	}
	tmp26 := tmp28[0]
	tmp29, err := this.IntArray()
	if err != nil {
		return 0, err
	}
	for _, tmp27 := range tmp29 {
		if tmp26 > tmp27 {
			tmp26 = tmp27
		}
	}
	this.intArrayMin = int(tmp26)
	this._f_intArrayMin = true
	return this.intArrayMin, nil
}
func (this *ExprCalcArrayOps) StrArraySize() (v int, err error) {
	if (this._f_strArraySize) {
		return this.strArraySize, nil
	}
	tmp30, err := this.StrArray()
	if err != nil {
		return 0, err
	}
	this.strArraySize = int(len(tmp30))
	this._f_strArraySize = true
	return this.strArraySize, nil
}
func (this *ExprCalcArrayOps) IntArrayFirst() (v int, err error) {
	if (this._f_intArrayFirst) {
		return this.intArrayFirst, nil
	}
	tmp31, err := this.IntArray()
	if err != nil {
		return 0, err
	}
	this.intArrayFirst = int(tmp31[0])
	this._f_intArrayFirst = true
	return this.intArrayFirst, nil
}
func (this *ExprCalcArrayOps) DoubleArrayFirst() (v float64, err error) {
	if (this._f_doubleArrayFirst) {
		return this.doubleArrayFirst, nil
	}
	tmp32, err := this.DoubleArray()
	if err != nil {
		return 0, err
	}
	this.doubleArrayFirst = float64(tmp32[0])
	this._f_doubleArrayFirst = true
	return this.doubleArrayFirst, nil
}
func (this *ExprCalcArrayOps) IntArrayMid() (v int, err error) {
	if (this._f_intArrayMid) {
		return this.intArrayMid, nil
	}
	tmp33, err := this.IntArray()
	if err != nil {
		return 0, err
	}
	this.intArrayMid = int(tmp33[1])
	this._f_intArrayMid = true
	return this.intArrayMid, nil
}
func (this *ExprCalcArrayOps) DoubleArrayMin() (v float64, err error) {
	if (this._f_doubleArrayMin) {
		return this.doubleArrayMin, nil
	}
	tmp36, err := this.DoubleArray()
	if err != nil {
		return 0, err
	}
	tmp34 := tmp36[0]
	tmp37, err := this.DoubleArray()
	if err != nil {
		return 0, err
	}
	for _, tmp35 := range tmp37 {
		if tmp34 > tmp35 {
			tmp34 = tmp35
		}
	}
	this.doubleArrayMin = float64(tmp34)
	this._f_doubleArrayMin = true
	return this.doubleArrayMin, nil
}
func (this *ExprCalcArrayOps) IntArray() (v []int, err error) {
	if (this._f_intArray) {
		return this.intArray, nil
	}
	this.intArray = []int([]int{10, 25, 50, 100, 200, 500, 1000})
	this._f_intArray = true
	return this.intArray, nil
}
func (this *ExprCalcArrayOps) IntArrayLast() (v int, err error) {
	if (this._f_intArrayLast) {
		return this.intArrayLast, nil
	}
	tmp39, err := this.IntArray()
	if err != nil {
		return 0, err
	}
	tmp38 := tmp39
	this.intArrayLast = int(tmp38[len(tmp38) - 1])
	this._f_intArrayLast = true
	return this.intArrayLast, nil
}